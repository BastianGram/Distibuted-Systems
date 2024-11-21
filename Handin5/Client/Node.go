package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin5/Handin5/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clientType struct {
	dials       map[int32]pb.ITUDatabaseClient // Map over all server
	currentport int32
}

func dial(client *clientType) {

	for otherPort := 5050; otherPort <= 5060; otherPort++ {
		conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			continue
		}
		client.dials[int32(otherPort)] = pb.NewITUDatabaseClient(conn)
	}
}

func checkMap(client *clientType) {
	for port := range client.dials {
		address := fmt.Sprintf("localhost:%d", port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			log.Printf("Client %d has crashed, removing it from list...", port)
			delete(client.dials, port)
			continue
		} else {
			client.currentport = port
		}
		conn.Close() // Close the connection if it was successful
	}
}

func main() {
	// Establish connection to the server
	//conn, err := grpc.Dial("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))

	/*conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	//defer conn.Close()


	// Create a new gRPC client
	client := pb.NewITUDatabaseClient(conn)*/
	var ID int32 = -1

	client := &clientType{
		dials: make(map[int32]pb.ITUDatabaseClient),
	}

	dial(client)

	// Main loop to process user input
	for {
		fmt.Print("Enter command: ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if len(input) > 3 && input[:3] == "bid" {
			// Extract the bid amount from the input string
			bidStr := input[4:] // Get the part after "bid "
			bidFloat, err := strconv.ParseFloat(strings.TrimSpace(bidStr), 64)
			if err != nil {
				fmt.Println("Invalid bid amount. Please use the format 'bid <amount>'.")
				continue
			}
			bidAmount := int32(bidFloat)

			// Send the bid to the server
			log.Printf("Sending bid of %d...", bidAmount)
			checkMap(client)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			ack, err := client.dials[client.currentport].Bid(ctx, &pb.BidAmount{
				Id:     ID,        // Replace with appropriate Id if needed
				Amount: bidAmount, // Use the extracted bid amount
			})

			if err != nil {
				log.Printf("Failed to send bid: %d", err)
				continue
			}

			if !ack.Answer {
				if ack.HighestBid == -1 {
					log.Printf("Bid is not large enough")
					continue
				}
				log.Printf("Auction ended")
				continue
			}
			if ID == -1 {
				ID = ack.Id
				log.Print("This client has ID: ", ID)
			}

			log.Printf("Bid sent successfully")
		} else if input == "result" {
			log.Printf("Requesting result")
			checkMap(client)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			result, err := client.dials[client.currentport].Result(ctx, &pb.Sync{})

			if err != nil {
				log.Printf("Failed to get result")
				continue
			} else {
				if result.Success {
					log.Printf("Auction is still going. Highest bid is: %d. From Client: %d", result.Amount, result.Id)
				} else {
					log.Printf("Auction is over highest bid was: %d. From Client: %d", result.Amount, result.Id)
					break
				}
			}

		} else {
			fmt.Println("Unknown command. Type 'bid <amount>' to send a bid to the server.")
		}
	}
}
