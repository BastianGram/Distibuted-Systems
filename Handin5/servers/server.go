package main

import (
	"bufio"
	"context"
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
	s := &server{currentBid: 0, HighestClientID: -1, auctionState: true, clientNumber: 0}

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
			break
		} else {
			log.Println("Unknown command. Type 'end auction' to end the auction manually.")
		}
	}

	// Ensure the gRPC server shuts down cleanly
	grpcServer.GracefulStop()
	log.Println("Server shut down.")
}
