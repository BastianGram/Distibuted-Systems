package main

import (
	"sync"
	"log"
	"net"
	"strconv"

	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin5/Handin5/grpc"
	"google.golang.org/grpc"
)

type Client struct {
	Name int32
}

// Server struct to implement the MyServiceServer interface
type server struct {
	clients map[int32]pb.ITUDatabaseServer
	mu      sync.Mutex // to protect access to clients map
	currentBid int32
}

func (s *server) bid(req *pb.BidAmount, stream pb.ITUDatabase_BidServer) (error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Println("Client bid received from Client nr. " + strconv.Itoa(int(req.Id)))

	if (req.Amount <= s.currentBid) {

	} else {
		// Create the event notification
		notification := &pb.Ack{
			Id: req.Id,
			Answer: true,
			HighestBid: req.Amount,
		}

		// Broadcast the event to all subscribed clients 
		for clientId, observer := range s.clients {
			err := observer.Send(notification)
			if err != nil {
				log.Printf("Error sending event to client %s: %v", clientId, err)
				delete(s.clients, clientId) // Remove disconnected client
			}
		}
	}


	return nil
}


func main() {
	// Initialize the server
	s := &server{clients: make(map[int32]pb.ITUDatabaseServer)}

	// Create a listener on TCP port 5050
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterITUDatabaseServer(grpcServer, s)

	log.Println("Server started. Listening on port 5050.")

	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}