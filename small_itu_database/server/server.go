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

var timestamp = 0

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
func (s *server) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponce, error) {
	//CLINR is the clientID
	CLINR++
	
	timestamp++
	s.mu.Lock()
	defer s.mu.Unlock()

	// Add the client to the map
	client := &Client{Name: strconv.Itoa(CLINR)}
	s.clients[client.Name] = client
	
	log.Printf("Client joined with request: %s", req.GetName())
	log.Printf("Client logged as: %s", client.Name)

	//converts timestamp to string
	timestampStr := strconv.Itoa(timestamp)
	return &pb.JoinResponce{Name: "Welcome client nr. " + strconv.Itoa(CLINR) + " Lamport timestamp: " + timestampStr}, nil
}

// Disconnect method implementation
func (s *server) ClientLeaving(ctx context.Context, req *pb.ClientLeaves) (*pb.ServerClientLeaves, error) {
	timestamp++
	s.mu.Lock()
	defer s.mu.Unlock()

	// Remove the client from the map
	if _, exists := s.clients[req.GetClientName()]; exists {
		delete(s.clients, req.GetClientName())
		log.Printf("Client disconnected: %s", req.GetClientName())

		//converts timestamp to string
		timestampStr := strconv.Itoa(timestamp)

		return &pb.ServerClientLeaves{ClientName: "Goodbye " + req.GetClientName() + "! Lamport timestamp: " + timestampStr}, nil
	}

	return &pb.ServerClientLeaves{ClientName: "Client not found."}, nil
}

// SendMessage method implementation
func (s *server) Broadcast(ctx context.Context, req *pb.BroadcastRequest) (*pb.ServerBroadcast, error) {
	timestamp++
	s.mu.Lock()
	defer s.mu.Unlock()

	// Log the received message from the client
	log.Printf("Message from %s: %s", req.GetClientName(), req.GetMessage())

	//converts timestamp to string
	timestampStr := strconv.Itoa(timestamp)

	return &pb.ServerBroadcast{Message: "Message received from " + req.GetClientName() + "!" + ", Message: " + req.GetMessage() + "! Lamport timestamp: " + timestampStr}, nil
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

/*
func (s *ITU_databaseServer) GetStudents(ctx context.Context, in *pb.Empty) (*pb.Students, error) {
	return &pb.Students{Students: s.students}, nil
}

func main() {
	file, err := os.Open("messages.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Print the CSV data
	for i, row := range data {
		if i == 0 {
			continue
		}
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}

	writer := csv.NewWriter(file)

	messageExample := []string{"Hey John"}
	writer.Write(messageExample)

	server := &ITU_databaseServer{students: []string{}}
	server.students = append(server.students, "John")
	server.students = append(server.students, "Jane")
	server.students = append(server.students, "Alice")
	server.students = append(server.students, "Bob")

	server.start_server()
}

func (s *ITU_databaseServer) start_server() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Did not work")
	}

	proto.RegisterITUDatabaseServer(grpcServer, s)

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Did not work")
	}

}
*/
