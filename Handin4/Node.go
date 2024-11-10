package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
	"errors"
	"bufio"
	"os"

	pb "github.com/BastianGram/Distibuted-Systems/Handin4/Handin4/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Server struct to implement the MyServiceServer interface
type TypeNode struct {
	pb.UnimplementedITUDatabaseServer
	Node int32 // equal to it's port
	Max int32
	conn *grpc.ClientConn // gRPC connection for client functionality
	CriticalSection bool
	clients map[int32]pb.ITUDatabaseClient

	//clients map[string]*Client
	mu sync.Mutex // to protect access to clients map
}

func NewTypeNode(port int) *TypeNode {
	node := &TypeNode{
		Node: int32(port),
		Max:  -1,
		clients: make(map[int32]pb.ITUDatabaseClient),
	}
	// Initialize connection to all expected nodes
	for otherPort := 5051; otherPort <= 5060; otherPort++ {
		
		if int32(otherPort) == node.Node {
			continue
		}
		// Check if the port is available before connecting
		if !isPortAvailable(otherPort) {
			continue
		}
		address := fmt.Sprintf("localhost:%d", otherPort)
		conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to connect to node %d: %v", otherPort, err)
			continue
		}

		node.clients[int32(otherPort)] = pb.NewITUDatabaseClient(conn)
	}
	return node
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	port := findAvailablePort(5051)
	node := NewTypeNode(port)
	node.NotifyExistingNodes()

	// Start the gRPC server in a goroutine
	go startGRPCServer(port, node)

	// Dial the server
	conn, err := grpc.DialContext(ctx, "localhost:"+strconv.Itoa(int(node.Node)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Printf("Failed to connect: %v\n", err)
		return
	}

	// Assign the connection to the TypeNode instance
	node.conn = conn

	client := pb.NewITUDatabaseClient(conn)

	req := &pb.RequestElection{
		ClientName:  int32(node.Node),
	}


	log.Println("Successfully joined. Electing... ")

	// Call the Election method and handle errors
	MaxClient, err := client.Election(ctx, req)
	
	if err != nil {
		log.Fatalf("Election call failed: %v", err)
		return
	}

	if MaxClient != nil {
		node.Max = MaxClient.MAX
	} else {
		log.Fatalf("MaxClient response is nil")
		return
	}

	var input string

	// Infinite loop to listen for user input
	for {
		// Read user input
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use for scanner.Scan() to keep reading
		input = scanner.Text()
		// If user types "disconnect", call the disconnect method
		if len(input) > 4 && input[:4] == "send" {
			checkMap(node.clients)

			if node.Node == node.Max {
				fmt.Println("Leader is in critical section")
				time.Sleep(5 * time.Second)
				continue
			}

			// Extract the message from input
			message := input[5:]

			ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", node.Max))
			if err != nil {
				fmt.Printf("Leader %d is dead. calling a new one...", node.Max)
				node.Election(ctx, &pb.RequestElection{ClientName: node.Node})
			} else{
				ln.Close() // Close the listener since we only wanted to check availability
			}

			client1 := node.clients[node.Max]

			client1.Broadcast(ctx, &pb.RequestCS{ClientName: node.Node, Message: message})
		} else {
			fmt.Println("Unknown command. Type 'send <message>' to send a message")
		}
	}
}

func startGRPCServer(port int, node *TypeNode) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Printf("Failed to listen on port %d: %v\n", port, err)
		return
	}
	grpcServer := grpc.NewServer()

	pb.RegisterITUDatabaseServer(grpcServer, node)

	log.Printf("Starting gRPC server on port %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("Failed to serve gRPC server: %v\n", err)
	}
}

// findAvailablePort tries to find an available localhost port starting from the provided port
func findAvailablePort(startPort int) int {
	for port := startPort; port <= 65535; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			ln.Close() // Close the listener since we only wanted to check availability
			return port
		} // If port is taken, it tries the next port
	}
	//Should not get here
	return 0
}

// Helper function to check if a service is listening on a given port
func isPortAvailable(port int) bool {
    address := fmt.Sprintf("localhost:%d", port)
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    if err != nil {
        return false // Connection failed, meaning no server is listening on this port
    }
    defer conn.Close() // Close the connection if it was successful
    return true // Connection successful, meaning something is listening
}

func checkMap(clients map[int32]pb.ITUDatabaseClient) {
	for port := range clients {
		address := fmt.Sprintf("localhost:%d", port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			log.Printf("Client %d has crashed, removing it from list...", port)
			delete(clients, port)
		}
		conn.Close() // Close the connection if it was successful

	}
}

func (N1 *TypeNode) Election(ctx context.Context, re1 *pb.RequestElection) (*pb.Answer, error) {
	N1.mu.Lock()
	defer N1.mu.Unlock()

	checkMap(N1.clients)

	if (len(N1.clients) == 0) {
		return &pb.Answer{MAX: N1.Node}, nil
	}

	if client1, exists := N1.clients[re1.ClientName+1]; exists {
		req := &pb.RequestElection{
			ClientName:  int32(N1.Node + 1),
		}
		return client1.Election(ctx, req)
	} else if client2, exists := N1.clients[re1.ClientName+2]; exists {
		req := &pb.RequestElection{
			ClientName:  int32(N1.Node + 1),
		}
		return client2.Election(ctx, req)
	} else if client3, exists := N1.clients[5051]; exists{
		req := &pb.IAmCoordinator{
			ClientName:  5051,
			Sender: N1.Node,
		}
		client3.Coordinator(ctx, req)

		return &pb.Answer{MAX: N1.Node}, nil
	} else if clinet4, exists := N1.clients[5052]; exists{
		req := &pb.IAmCoordinator{
			ClientName:  5052,
			Sender: N1.Node,
		}
		clinet4.Coordinator(ctx, req)
		return &pb.Answer{MAX: N1.Node}, nil
	} else {

		return &pb.Answer{MAX: N1.Node}, errors.New("very bad fail, " + strconv.Itoa(int(re1.ClientName)))
	}
}

func (N1 *TypeNode) Coordinator(ctx context.Context, re1 *pb.IAmCoordinator) (*pb.SendsAllegiance, error) {
	//N1.mu.Lock()

	checkMap(N1.clients)

	if re1.Sender == N1.Node {
		N1.Max = N1.Node
		log.Printf("%d is the new coordinator", N1.Max)
		return &pb.SendsAllegiance{ClientName: N1.Node}, nil
	}

 	// Determine the next client to contact in the chain, if available
 	var nextClient pb.ITUDatabaseClient
 	var nextClientName int32
	

	if client1, exists := N1.clients[re1.ClientName+1]; exists {
		nextClient = client1
        nextClientName = re1.ClientName + 1
	} else if client2, exists := N1.clients[re1.ClientName+2]; exists {
		nextClient = client2
        nextClientName = re1.ClientName + 2
	} else {
		N1.mu.Unlock()
		return &pb.SendsAllegiance{ClientName: N1.Node}, errors.New("next node in coordinator chain not available")
	}

	// Unlock before making the recursive gRPC call
    // N1.mu.Unlock()

    // Create the request for the next client in the chain
    req := &pb.IAmCoordinator{
        ClientName: nextClientName,
        Sender:     re1.Sender,
    }

    // Make the recursive call without holding the lock to avoid deadlock
    response, err := nextClient.Coordinator(ctx, req)
    if err != nil {
        log.Printf("Coordinator chain failed at node %d: %v\n", nextClientName, err)
        return &pb.SendsAllegiance{ClientName: N1.Node}, err
    }

	N1.Max = response.ClientName
	log.Printf("%d is the new coordinator", N1.Max)

    return response, nil
}

func (N1 *TypeNode) Broadcast(ctx context.Context, re1 *pb.RequestCS) (*pb.ResponseCS, error) {
	N1.mu.Lock()
	defer N1.mu.Unlock()

	//client, _ := N1.clients[N1.Max]

	if (N1.Max != N1.Node) {
		return &pb.ResponseCS{ClientName: N1.Node, Message: "fail"}, errors.New("the wrong leader is in charge")
	}

	N1.CriticalSection = false

	fmt.Println("Client is accessing critical section")
	time.Sleep(5 * time.Second)

	N1.CriticalSection = true
	
	return &pb.ResponseCS{ClientName: N1.Node, Message: "Message send"}, nil
}



func (node *TypeNode) NotifyExistingNodes() {
	for port := 5051; port <= 5060; port++ {

		if int32(port) == node.Node {
			continue
		}

		// Try to connect to existing node
		if client, exists := node.clients[int32(port)]; exists {
			req := &pb.JoinRequest{Port: int32(node.Node)}
			_, err := client.NotifyJoin(context.Background(), req)
			if err != nil {
				log.Printf("Failed to notify node %d: %v", port, err)
			}
		}
	}
}

func (N1 *TypeNode) NotifyJoin(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	N1.mu.Lock()
	defer N1.mu.Unlock()

	newNodePort := req.Port
	if newNodePort == N1.Node {
		// Ignore self-notification
		return &pb.JoinResponse{Success: true}, nil
	}

	// Check if the node is already in the clients map
	if _, exists := N1.clients[int32(newNodePort)]; exists {
		return &pb.JoinResponse{Success: true}, nil
	}

	// Attempt to connect to the new node
	address := fmt.Sprintf("localhost:%d", newNodePort)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to new node %d: %v", newNodePort, err)
		return &pb.JoinResponse{Success: false}, err
	}

	// Add the new client to the clients map
	N1.clients[int32(newNodePort)] = pb.NewITUDatabaseClient(conn)
	log.Printf("Connected to new node on port %d", newNodePort)

	return &pb.JoinResponse{Success: true}, nil
}
