package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin5/Handin5/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	Name int32
}

// Server struct to implement the MyServiceServer interface
type server struct {
	pb.ITUDatabaseServer
	mu           sync.Mutex // to protect access to clients map
	currentBid   int32
	HighestClientID int32
	auctionState bool
	clientNumber int32
}


// Is called whenever a client bids
func (s *server) Bid(ctx context.Context, req *pb.BidAmount) (*pb.Ack, error) {
	//Locks and is only unlocked once the function ends
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.auctionState {
		log.Printf("No more bids allowed, the auction is over")
		return &pb.Ack{
			Id:         s.clientNumber,
			Answer:     false,
			HighestBid: s.currentBid,
		}, nil
	}


	if req.Id == -1 {
		s.clientNumber++
		log.Println("Client bid received from Client nr. " + strconv.Itoa(int(s.clientNumber)) + ", bid is: " + strconv.Itoa(int(req.Amount)))
	} else {
		log.Println("Client bid received from Client nr. " + strconv.Itoa(int(req.Id)) + ", bid is: " + strconv.Itoa(int(req.Amount)))
	}

	//Is amount too small
	if req.Amount <= s.currentBid {
		log.Printf("Bid is not large enough")
		return &pb.Ack{
			Id:         s.clientNumber,
			Answer:     true,
			HighestBid: s.currentBid,
		}, nil
		//Amount is larger than previous bid
	} else {
		s.currentBid = req.Amount
		s.HighestClientID = req.Id
		// Create the event notification
		return &pb.Ack{
			Id:         s.clientNumber,
			Answer:     true,
			HighestBid: req.Amount,
		}, nil
	}
}

func (s *server) result() (*pb.Ack, error) {

	notification := &pb.Ack{
		Id:         s.HighestClientID,
		Answer:     s.auctionState,
		HighestBid: s.currentBid,
	}
	return notification, nil
}

/*
	func (s *server) checkMap() {
		s.mu.Lock()
		defer s.mu.Lock()
		for clientId, observer := range s.clients {
			// Send a lightweight ping or check message
			err := observer.Send(&pb.Ack{
				Id:         clientId,
				Answer:     false,        // Not an actual bid update
				HighestBid: s.currentBid, // Current highest bid
			})
			if err != nil {
				log.Printf("Client %d disconnected, removing from clients: %v", clientId, err)
				delete(s.clients, clientId) // Remove the disconnected client
			}
		}
	}
*/
func GoServe(grpcServer *grpc.Server, lis net.Listener) {
	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {

	// Initialize the server
	s := &server{currentBid: 0,HighestClientID: -1 , auctionState: true, clientNumber: 0}

	// Create a listener on TCP port 5050
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	pb.RegisterITUDatabaseServer(grpcServer, s)

	log.Println("Server started. Listening on port 5050.")
	log.Println("Auction started, awaiting bids")

	go GoServe(grpcServer, lis)

	// Listen for user input
	log.Println("Type 'end auction' to prevent further bids")

	var input string

	// Infinite loop to listen for user input
	for {
		// Read user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use for scanner.Scan() to keep reading
		input = scanner.Text()
		// If user types "disconnect", call the disconnect method
		if input == "end auction" {
			s.auctionState = false
		} else {
			fmt.Println("Unknown command. Type 'end auction' to prevent further bids")
		}
	}
}
