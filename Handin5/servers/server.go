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
	"time"

	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin5/Handin5/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	Name int32
}

// Server struct to implement the MyServiceServer interface
type server struct {
	pb.ITUDatabaseServer
	mu              sync.Mutex // to protect access to clients map
	currentBid      int32
	HighestClientID int32
	auctionState    bool
	clientNumber    int32
	isPrimary bool
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
		log.Println("Client bid received from NEW Client nr. " + strconv.Itoa(int(s.clientNumber)) + ", bid is: " + strconv.Itoa(int(req.Amount)))
	} else {
		log.Println("Client bid received from Client nr. " + strconv.Itoa(int(req.Id)) + ", bid is: " + strconv.Itoa(int(req.Amount)))
	}

	//Is amount too small
	if req.Amount <= s.currentBid {
		log.Printf("Bid is not large enough")
		return &pb.Ack{
			Id:         s.clientNumber,
			Answer:     false,
			HighestBid: -1,
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

func (s *server) Result(ctx context.Context, sync *pb.Sync) (*pb.Results, error) {
	notification := &pb.Results{
		Id:         s.HighestClientID,
		Success:     s.auctionState,
		Amount: s.currentBid,
	}
	return notification, nil
}

func GoServe(grpcServer *grpc.Server, lis net.Listener) {
	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func isPortAvailable(port int) bool {
    address := fmt.Sprintf("localhost:%d", port)
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    if err != nil {
        return false // Connection failed, meaning no server is listening on this port
    }
    defer conn.Close() // Close the connection if it was successful
    return true // Connection successful, meaning something is listening
}

func main() {
	if 

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterITUDatabaseServer(grpcServer, s)

	log.Println("Server started. Listening on port 5050.")
	log.Println("Auction started, awaiting bids")

	// Start the gRPC server in a separate goroutine
	go GoServe(grpcServer, lis)

	// Define auction duration
	auctionDuration := 200 // Auction duration in seconds

	// Start a timer for the auction
	go func() {
		log.Printf("Auction will automatically end in %d seconds.", auctionDuration)
		<-time.After(time.Duration(auctionDuration) * time.Second)
		s.mu.Lock()
		s.auctionState = false
		s.mu.Unlock()
		log.Println("Auction has ended. No further bids are allowed.")
	}()

	// Listen for user input
	log.Println("Type 'end auction' to end the auction manually.")

	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Read user input
		scanner.Scan()
		input = scanner.Text()

		// If user types "end auction", stop the auction
		if input == "end auction" {
			s.mu.Lock()
			s.auctionState = false
			s.mu.Unlock()
			log.Println("Auction ended manually. No further bids are allowed.")
		} else {
			log.Println("Unknown command. Type 'end auction' to end the auction manually.")
		}
	}

	// Ensure the gRPC server shuts down cleanly
	grpcServer.GracefulStop()
	log.Println("Server shut down.")
}
