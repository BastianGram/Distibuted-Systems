package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"sync"

	pb "github.com/BastianGram/Distibuted-Systems/tree/handin3v2/small_itu_database/grpc"
	"google.golang.org/grpc"
)

var lamport int32 = 0

type Client struct {
	Name string
}

// Server struct to implement the MyServiceServer interface
type server struct {
	pb.UnimplementedITUDatabaseServer
	clients map[string]*Client
	mu      sync.Mutex // to protect access to clients map
}

// client id:
var CLINR int = 100

// Join method implementation
func (s *server) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	//CLINR is the clientID
	CLINR++

	lamport++
	s.mu.Lock()
	defer s.mu.Unlock()

	// Add the client to the map
	client := &Client{Name: strconv.Itoa(CLINR)}
	s.clients[client.Name] = client

	log.Printf("Client logged as: %s", client.Name)

	return &pb.JoinResponse{
		LamportTime: lamport,
		Name:        strconv.Itoa(CLINR),
		Approved:    "Welcome " + strconv.Itoa(CLINR),
	}, nil
}

// Disconnect method implementation
func (s *server) ClientLeaving(ctx context.Context, req *pb.ClientLeaves) (*pb.ServerClientLeaves, error) {
	lamport++
	s.mu.Lock()
	defer s.mu.Unlock()

	// Remove the client from the map
	if _, exists := s.clients[req.GetClientName()]; exists {
		delete(s.clients, req.GetClientName())
		log.Printf("Lamport: "+strconv.Itoa(int(lamport))+", "+"Client disconnected: %s", req.GetClientName())

		return &pb.ServerClientLeaves{
			LamportTime: lamport,
			ClientName:  req.ClientName,
		}, nil
	}

	return &pb.ServerClientLeaves{ClientName: "Client not found."}, nil
}

// SendMessage method implementation
func (s *server) Broadcast(ctx context.Context, req *pb.BroadcastRequest) (*pb.ServerBroadcast, error) {
	lamport++
	s.mu.Lock()
	defer s.mu.Unlock()

	// Log the received message from the client
	log.Printf("Message from %s: %s", req.GetClientName(), req.GetMessage())

	return &pb.ServerBroadcast{
		LamportTime: lamport,
		Message:     req.GetMessage(),
		ClientName:  req.GetClientName(),
	}, nil

	//code to broadcast to every client
}

func main() {
	// Initialize the server
	s := &server{clients: make(map[string]*Client)}

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
