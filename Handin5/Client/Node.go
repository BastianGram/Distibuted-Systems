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
	primarydial       pb.ITUDatabaseClient
	secondarydial     pb.ITUDatabaseClient
	primarydialport   int32
	secondarydialport int32
	currentport       int32
	currentdial       pb.ITUDatabaseClient
}

func dial(client *clientType) {

	log.Println("Dialing..")
	primaryaddress := fmt.Sprintf("localhost:%d", client.primarydialport)
	secondaryaddress := fmt.Sprintf("localhost:%d", client.secondarydialport)

	conn1, err1 := grpc.NewClient(primaryaddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err1 != nil {
		log.Printf("Error connecting to primary")
	}
	client.primarydial = pb.NewITUDatabaseClient(conn1)
	conn2, err2 := grpc.NewClient(secondaryaddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err2 != nil {
		log.Printf("Error connecting to secondary")
	}
	client.secondarydial = pb.NewITUDatabaseClient(conn2)
	log.Println("Dialing done")
}

func checkPort(client *clientType) {

	address1 := fmt.Sprintf("localhost:%d", client.primarydialport)
	_, err1 := net.DialTimeout("tcp", address1, 10*time.Second)
	if err1 != nil {
		log.Printf("Primary server has crashed")

	} else {
		log.Printf("Using port, %d", client.primarydialport)
		client.currentport = client.primarydialport
		client.currentdial = client.primarydial
		return
	}

	address2 := fmt.Sprintf("localhost:%d", client.secondarydialport)
	_, err2 := net.DialTimeout("tcp", address2, 10*time.Second)
	if err2 != nil {
		log.Printf("Secondary server has crashed")

	} else {
		log.Printf("Using port, %d", client.secondarydialport)
		client.currentport = client.secondarydialport
		client.currentdial = client.secondarydial
		return
	}

}

func main() {
	
	var ID int32 = -1

	client := &clientType{
		primarydialport:   5050,
		secondarydialport: 5051,
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
			checkPort(client)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			ack, err := client.currentdial.Bid(ctx, &pb.BidAmount{
				Id:     ID,        // Replace with appropriate Id if needed
				Amount: bidAmount, // Use the extracted bid amount
			})

			if ID == -1 {
				ID = ack.Id
				log.Print("This client has ID: ", ID)
			}

			if !ack.Answer {
				log.Printf("Auction ended")
				continue
			}

			if err != nil {
				log.Printf("Failed to send bid: %d", err)
				continue
			}

			if ack.HighestBid == -1 {
				log.Printf("Bid is not large enough")
				continue
			}

			log.Printf("Bid sent successfully")
		} else if input == "result" {
			log.Printf("Requesting result")
			checkPort(client)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			result, err := client.currentdial.Result(ctx, &pb.Sync{})

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
