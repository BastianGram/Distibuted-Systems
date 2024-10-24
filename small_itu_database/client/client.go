package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	pb "github.com/BastianGram/Distibuted-Systems/tree/handin3v2/small_itu_database/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var name string

	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}

	client := pb.NewITUDatabaseClient(conn)

	// Prepare request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Join the server
	response, err := client.Join(ctx, &pb.JoinRequest{Name: "Let me join"})
	if err != nil {
		log.Fatalf("could not join: %v", err)
	}
	log.Println(response.Name)

	// Define a regular expression to capture the integer between "Welcome" and "Lamport"
    re := regexp.MustCompile(`client nr. (\d+) Lamport`)

	// Find the first match
    match := re.FindStringSubmatch(response.Name)

	name = match[1]

	// Listen for user input
	fmt.Println("Type 'send <message>' to send a message, or 'disconnect' to disconnect from the server:")
	var input string

	// Infinite loop to listen for user input
	for {
		// Read user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		input = scanner.Text()

		// If user types "disconnect", call the disconnect method
		if input == "disconnect" {
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			disconnectResponse, err := client.ClientLeaving(ctx, &pb.ClientLeaves{ClientName: name})
			if err != nil {
				log.Fatalf("could not disconnect: %v", err)
			}
			log.Println(disconnectResponse.ClientName + "has left the session")
			break
		} else if len(input) > 4 && input[:4] == "send" {
			// Extract the message from input
			message := input[5:]
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			sendResponse, err := client.Broadcast(ctx, &pb.BroadcastRequest{ClientName: name, Message: message})
			if err != nil {
				log.Fatalf("could not send message: %v", err)
			}
			log.Println("Server: " + sendResponse.Message)
		} else {
			fmt.Println("Unknown command. Type 'send <message>' to send a message or 'disconnect' to disconnect from the server.")
		}
	}
}
