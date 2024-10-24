package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"strconv"

	pb "github.com/BastianGram/Distibuted-Systems/tree/handin3v2/small_itu_database/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	var lamport int32 = 0

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
	response, err := client.Join(ctx, &pb.JoinRequest{LamportTime: lamport})
	if err != nil {
		log.Fatalf("could not join: %v", err)
	}
	log.Println("Lamport: " + strconv.Itoa(int(lamport)) + ", " + response.Approved)

	//client gets its name
	name = response.Name

	//time is updated
	if (response.LamportTime > lamport) {
		lamport = response.LamportTime + 1
	} else {
		lamport++
	}

	// Listen for user input
	fmt.Println("Type 'send <message>' to send a message, or 'disconnect' to disconnect from the server:")
	var input string

	// Infinite loop to listen for user input
	for {
		// Read user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		input = scanner.Text()
		lamport++
		// If user types "disconnect", call the disconnect method
		if input == "disconnect" {
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			disconnectResponse, err := client.ClientLeaving(ctx, &pb.ClientLeaves{LamportTime: lamport, ClientName: name})
			if err != nil {
				log.Fatalf("could not disconnect: %v", err)
			}
			log.Println(disconnectResponse.ClientName + " has left the session")
			return
		} else if len(input) > 4 && input[:4] == "send" {
			// Extract the message from input
			message := input[5:]
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			sendResponse, err := client.Broadcast(ctx, &pb.BroadcastRequest{LamportTime: lamport, Message: message, ClientName: name})
			if err != nil {
				log.Fatalf("could not send message: %v", err)
			}
			//time is updated
			if (sendResponse.LamportTime > lamport) {
				lamport = response.LamportTime + 1
			} else {
				lamport++
			}

			//future code for getting messages for other clients
			log.Println("Lamport: " + strconv.Itoa(int(lamport)) + ", " + "Server: " + sendResponse.Message)
		} else {
			fmt.Println("Unknown command. Type 'send <message>' to send a message or 'disconnect' to disconnect from the server.")
		}
	}
}
