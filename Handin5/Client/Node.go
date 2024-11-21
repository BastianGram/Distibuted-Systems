package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin5/Handin5/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish connection to the server
	conn, err := grpc.Dial("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	var ID int32 = -1

	// Create a new gRPC client
	client := pb.NewITUDatabaseClient(conn)

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
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			ack, err := client.Bid(ctx, &pb.BidAmount{
				Id:     ID,        // Replace with appropriate Id if needed
				Amount: bidAmount, // Use the extracted bid amount
			})
			if !ack.Answer {
				log.Printf("No more bids allowed, the auction is over. Highest bid was: " + strconv.Itoa(int(ack.HighestBid)))
				break
			}
			if ID == -1 {
				ID = ack.Id
				log.Print("This client has ID: " , ID)
			}
			if err != nil {
				log.Printf("Failed to send bid: %d", err)
			} else {
				if ack.HighestBid > bidAmount {
					log.Printf("Bid is not large enough. Current largest bid is: %d", ack.HighestBid)
				} else {
					log.Printf("Bid sent successfully")
				}
			}
		} else if input == "result" {
			log.Printf("Requesting result")
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			result, err := client.Result(ctx, &pb.Sync{
				Answer: true,
			})

			if err != nil {
				log.Printf("Failed to get result")
				continue
			} else {
				if result.Success {
					log.Printf("Action is still going. Highest bid is: %d. From Client: %d", result.Amount, result.Id)
				} else {
					log.Printf("Auction is over highest bid was: %d. From Client: %d", result.Amount, result.Id)
				}
			}

		} else {
			fmt.Println("Unknown command. Type 'bid <amount>' to send a bid to the server.")
		}
	}
}
