package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"context"
	"time"
	

	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin5/Handin5/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}

	// Creating a new client
	client := pb.NewITUDatabaseClient(conn)

	var input string

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)


	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use for scanner.Scan() to keep reading
		input = scanner.Text()
		if len(input) > 3 && input[:3] == "bid" {

			//bid := input[4:]

			
			log.Println("Requesting bid...")
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err := client.Bid(ctx, &pb.BidAmount{
				Id: 0,
				Amount: 0,
			})
			if err != nil {
				log.Fatalf("could not send message: %v", err)
			}

		} else {
			fmt.Println("Unknown command. Type 'bid <amount>' to send a bid to the server.")
		}
		
	}
}
