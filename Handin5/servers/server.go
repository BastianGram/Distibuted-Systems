package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"bufio"
	"os"

	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin5/Handin5/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	Name int32
}

// Server struct to implement the MyServiceServer interface
type server struct {
	pb.ITUDatabaseServer
	clients    map[int32]pb.ITUDatabase_BidServer
	mu         sync.Mutex // to protect access to clients map
	currentBid int32
	auctionState bool
}

// Is called whenever a client bids
func (s *server) bid(req *pb.BidAmount, stream pb.ITUDatabase_BidServer) error {
	//Locks and is only unlocked once the function ends
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checkMap()

	var foundclient bool = false
	for ClientID := range s.clients {
		if ClientID == req.Id {
			foundclient = true
		}
	}
	if !foundclient {
		s.clients[req.Id] = stream
	}

	log.Println("Client bid received from Client nr. " + strconv.Itoa(int(req.Id)))

	//Is amount too small
	if req.Amount <= s.currentBid {
		return errors.New("bid is not large enough")
		//Amount is larger than previous bid
	} else {
		// Create the event notification
		notification2 := &pb.Ack{
			Id:         req.Id,
			Answer:     true,
			HighestBid: req.Amount,
		}

		// Broadcast the event to all subscribed clients
		for clientId, observer := range s.clients {
			err := observer.Send(notification2)
			if err != nil {
				log.Printf("Error sending event to client %s: %v", strconv.Itoa(int(clientId)), err)
				delete(s.clients, clientId) // Remove disconnected client
			}
		}
	}
	return nil
}

func (s *server) result() (*pb.Ack, error) {
	notification := &pb.Ack{
		Id:         -1,
		Answer:     s.auctionState,
		HighestBid: s.currentBid,
	}
	return notification, nil
}

func (s *server) checkMap() {
	s.mu.Lock()
	for clientId, observer := range s.clients {
		// Send a lightweight ping or check message
		err := observer.Send(&pb.Ack {
			Id:         clientId,
			Answer:     false, // Not an actual bid update
			HighestBid: s.currentBid, // Current highest bid
		})
		if err != nil {
			log.Printf("Client %d disconnected, removing from clients: %v", clientId, err)
			delete(s.clients, clientId) // Remove the disconnected client
		}
	}
	s.mu.Unlock()
}

func main() {
	
	// Initialize the server
	s := &server{clients: make(map[int32]pb.ITUDatabase_BidServer), currentBid: -1, auctionState: true}

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

	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Listen for user input
	fmt.Println("Type 'end auction' to prevent further bids")
	
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
			break
			
		} else {
			fmt.Println("Unknown command. Type 'end auction' to prevent further bids")
		}
	}
}
