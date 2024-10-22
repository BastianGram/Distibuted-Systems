package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin3/Hand_in_3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func client() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChittyChatClient(conn)

	participantName := os.Args[1]

	// Join the chat
	joinResp, err := client.Join(context.Background(), &pb.JoinRequest{
		ParticipantName: participantName,
	})
	if err != nil {
		log.Fatalf("could not join: %v", err)
	}
	log.Printf("Join response: %s", joinResp.Message)

	// Create a stream for receiving broadcasted messages
	stream, err := client.BroadcastMessage(context.Background())
	if err != nil {
		log.Fatalf("could not establish broadcast stream: %v", err)
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("error receiving broadcast: %v", err)
			}
			log.Printf("Received broadcast: %s at Lamport time %d", resp.Message, resp.LamportTime)
		}
	}()

	// Publish a message
	for {
		var message string
		log.Println("Enter a message:")
		fmt.Scanln(&message)

		publishResp, err := client.Publish(context.Background(), &pb.PublishRequest{
			ParticipantName: participantName,
			Message:         message,
		})
		if err != nil {
			log.Fatalf("could not publish: %v", err)
		}
		log.Printf("Publish response: %s", publishResp.Message)
	}
}
