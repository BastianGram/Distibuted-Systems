package main

import (
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
	clients map[string]pb.ITUDatabase_BroadcastServer

	//clients map[string]*Client
	mu      sync.Mutex // to protect access to clients map
}

// client id:
var CLINR int = 100

// Join method implementation
func (s *server) Join(req *pb.ClientMessage, stream pb.ITUDatabase_BroadcastServer) (error) {
	//CLINR is the clientID
	CLINR++
	s.mu.Lock()
	var ThisClientID string = strconv.Itoa(CLINR)
	s.clients[ThisClientID] = stream
    s.mu.Unlock()

	if (req.LamportTime > lamport) {
		lamport = req.LamportTime + 1
	} else {
		lamport++
	}

	log.Printf("Lamport: " + strconv.Itoa(int(lamport)) + " Client nr. " + strconv.Itoa(CLINR) + " joined")

	lamport++
	// Create the event notification
	notification := &pb.ServerMessage{
		LamportTime: lamport,
		ClientName: ThisClientID,
		Message: "Welcome " + strconv.Itoa(CLINR),
	}

	// Broadcast the event to all subscribed clients 
	// Adding a goroutine around the code to allow the client to join the notification stream
	go func() {
		s.mu.Lock()
        defer s.mu.Unlock()
		for clientId, observer := range s.clients {
			err := observer.Send(notification)
			if err != nil {
				log.Printf("Error sending event to client %s: %v", clientId, err)
				delete(s.clients, clientId) // Remove disconnected client
			}
		}
	}()

	<-stream.Context().Done()

	// Remove the client when disconnected
	s.mu.Lock()
	delete(s.clients, ThisClientID)
	s.mu.Unlock()

	return nil

}

// Disconnect method implementation
func (s *server) ClientLeaving(req *pb.ClientMessage, stream pb.ITUDatabase_BroadcastServer) (error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if (req.LamportTime > lamport) {
		lamport = req.LamportTime + 1
	} else {
		lamport++
	}
	// Log the received message from the client
	log.Printf("Lamport: " + strconv.Itoa(int(lamport)) + " client nr." + req.GetClientName() + " has left the session" )

	lamport++
	// Create the event notification
	notification := &pb.ServerMessage{
		LamportTime: lamport,
		ClientName: req.ClientName,
		Message: "Client leaving, ID: " + req.ClientName,
	}

	// Broadcast the event to all subscribed clients 
	for clientId, observer := range s.clients {
		err := observer.Send(notification)
		if err != nil {
			log.Printf("Error sending event to client %s: %v", clientId, err)
			delete(s.clients, clientId) // Remove disconnected client
		}
	}

    return nil

}

// SendMessage method implementation
func (s *server) Broadcast(req *pb.ClientMessage, stream pb.ITUDatabase_BroadcastServer) (error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if (req.LamportTime > lamport) {
		lamport = req.LamportTime + 1
	} else {
		lamport++
	}

	// Log the received message from the client
	log.Printf("Lamport: " + strconv.Itoa(int(lamport)) + " Message: "  + req.Message + " recieved from client nr. " + req.GetClientName())

	lamport++
	// Create the event notification
	notification := &pb.ServerMessage{
		LamportTime: lamport,
		ClientName: strconv.Itoa(CLINR),
		Message: req.Message,
	}

	// Broadcast the event to all subscribed clients 
	for clientId, observer := range s.clients {
		err := observer.Send(notification)
		if err != nil {
			log.Printf("Error sending event to client %s: %v", clientId, err)
			delete(s.clients, clientId) // Remove disconnected client
		}
	}
	return nil
}

func main() {
	// Initialize the server
	s := &server{clients: make(map[string]pb.ITUDatabase_JoinServer)}

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
