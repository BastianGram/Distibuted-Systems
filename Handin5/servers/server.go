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
	"google.golang.org/grpc/credentials/insecure"
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
	isPrimary       bool
	secondary pb.ITUDatabaseClient
}

// Is called whenever a client bids
func (s *server) Bid(ctx context.Context, req *pb.BidAmount) (*pb.Ack, error) {
	var returnClientID int32 = -1
	if s.isPrimary {
		s.secondary.Bid(ctx, req)
	}

	if !s.auctionState {
		log.Printf("No more bids allowed, the auction is over")
		return &pb.Ack{
			Id:         returnClientID,
			Answer:     false,
			HighestBid: s.currentBid,
		}, nil
	}
	s.mu.Lock()
	if req.Id == -1 {
		s.clientNumber++
		returnClientID = s.clientNumber
		log.Println("Client bid received from NEW Client nr. " + strconv.Itoa(int(s.clientNumber)) + ", bid is: " + strconv.Itoa(int(req.Amount)))
	} else {
		log.Println("Client bid received from Client nr. " + strconv.Itoa(int(req.Id)) + ", bid is: " + strconv.Itoa(int(req.Amount)))
	}
	s.mu.Unlock()


	//Is amount too small
	if req.Amount <= s.currentBid {
		log.Printf("Bid is not large enough")
		return &pb.Ack{
			Id:         returnClientID,
			Answer:     true,
			HighestBid: -1,
		}, nil
		//Amount is larger than previous bid
	} else {
		s.mu.Lock()
		s.currentBid = req.Amount
		s.HighestClientID = req.Id
		s.mu.Unlock()
		// Create the event notification
		return &pb.Ack{
			Id:         returnClientID,
			Answer:     true,
			HighestBid: req.Amount,
		}, nil
	}
}

func (s *server) Result(ctx context.Context, sync *pb.Sync) (*pb.Results, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	notification := &pb.Results{
		Id:      s.HighestClientID,
		Success: s.auctionState,
		Amount:  s.currentBid,
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
	return true        // Connection successful, meaning something is listening
}

func (s *server) endAuction(ctx context.Context, sync *pb.Sync) (*pb.Ack, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.auctionState = false
	return &pb.Ack{
		Id:         0,
		Answer:     true,
		HighestBid: 0,
	}, nil
}

func main() {

	s := &server{currentBid: 0, HighestClientID: -1, auctionState: true, clientNumber: 0}
	var lis net.Listener
	var err error
	if !isPortAvailable(5050) {
		// Create a listener on TCP port 5050
		lis, err = net.Listen("tcp", ":5050")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("Server started. Listening on port 5050.")

		s.isPrimary = true

		log.Println("Primary server open. start second server to complete setup")

		address := fmt.Sprintf("localhost:%d", 5051)
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to connect to secondary %d: %v", 5051, err)
		} else  {
			log.Println("connected to secondary server")
		}
		s.secondary = pb.NewITUDatabaseClient(conn)

	} else {
		lis, err = net.Listen("tcp", ":5051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Println("Server started. Listening on port 5051.")

		s.isPrimary = false
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterITUDatabaseServer(grpcServer, s)

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
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			s.auctionState = false
			s.secondary.EndAuction(ctx, &pb.Sync{})
			cancel()
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
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			s.auctionState = false
			if (s.isPrimary) {
				ack, err := s.secondary.EndAuction(ctx, &pb.Sync{})
				if (err != nil){
					log.Printf("Failed to send end auction to secondary: %s", err)
				} else{
					log.Printf("recieved acknoledgement %t", ack.Answer)
				}
			}
			s.mu.Unlock()
			log.Println("Auction ended manually. No further bids are allowed.")
		} else {
			log.Println("Unknown command. Type 'end auction' to end the auction manually.")
		}
	}
}
