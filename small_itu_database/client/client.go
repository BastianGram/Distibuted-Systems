package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode/utf8"

	pb "github.com/BastianGram/Distibuted-Systems/tree/handin3v2/small_itu_database/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var lamport int32 = 1
var name string = ""

func Join(client pb.ITUDatabaseClient) {
	req := &pb.ClientMessage{LamportTime: lamport + 1}
	stream, err := client.Join(context.Background(), req)
	log.Println("requesting server to join with lamport 2 (incremented from 1)")
	if err != nil {
		log.Fatalf("Error subscribing to events: %v", err)
	}
	
	//Increased from message sent
	lamport++


	// Receive event notifications from the server
	for {
		event, err := stream.Recv()
		if name == "" {
			name = event.ClientName
		}
		if err != nil {
			log.Fatalf("Error receiving event: %v", err)
		}
		//time is updated
		if (event.LamportTime > lamport) {
			lamport = event.LamportTime + 1
		} else {
			lamport++
		}
		log.Println("Server: Lamport: " + strconv.Itoa(int(lamport)) + ". Message: " + event.Message + ". from: " + event.GetClientName())
	}

}

func main() {

	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}

	// Creating a new client
	client := pb.NewITUDatabaseClient(conn)
	
	// Join the server
	go Join(client)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Listen for user input
	fmt.Println("Type 'send <message>' to send a message, or 'disconnect' to disconnect from the server:")
	
	var input string

	// Infinite loop to listen for user input
	for {
		// Read user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use for scanner.Scan() to keep reading
		input = scanner.Text()
		// If user types "disconnect", call the disconnect method
		if input == "disconnect" {
			lamport++
			log.Println("Leaving server. Requesting with lamport: " + strconv.Itoa(int(lamport)))
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err := client.ClientLeaving(ctx, &pb.ClientMessage{LamportTime: lamport, ClientName: name, Message: "Client disconnected"})
			if err != nil {
				log.Fatalf("could not disconnect: %v", err)
			}

			//allow time for leave message broadcast
			time.Sleep(500 * time.Millisecond)
			return
		} else if len(input) > 4 && input[:4] == "send" {
			// Extract the message from input
			message := input[5:]

			if !CheckMessageLength(message) {
				fmt.Println("Error: Message too long. Cannot exceed 128 characters.")
				continue
			}
			
			lamport++
			log.Println("Requesting server broadcast with lamport: " + strconv.Itoa(int(lamport)))
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err := client.Broadcast(ctx, &pb.ClientMessage{
				LamportTime: lamport,
				Message: message,
				ClientName: name,
			})
			if err != nil {
				log.Fatalf("could not send message: %v", err)
			}

		} else {
			fmt.Println("Unknown command. Type 'send <message>' to send a message or 'disconnect' to disconnect from the server.")
		}
	}
}

func CheckMessageLength(message string) bool {
	if utf8.RuneCountInString(message) <= 128 {
		return true
	} else {
		return false
	}
}
