package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	pb "github.com/BastianGram/Distibuted-Systems/Handin4/Handin4/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Server struct to implement the MyServiceServer interface
type TypeNode struct {
	pb.UnimplementedITUDatabaseServer

	//clients map[string]*Client
	mu sync.Mutex // to protect access to clients map
}

var Node int32 // equal to it's port
var Max int32 = -1
var MaxClient *pb.Answer

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	Node := findAvailablePort(5050)
	fmt.Printf("Found available port: %d\n", Node)
	go startGRPCServer(Node)

	// Dial the server
	conn, err := grpc.DialContext(ctx, "localhost:"+strconv.Itoa(int(Node)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}

	req := &pb.RequestElection{
		LamportTime: 0,
		ClientName:  int32(Node + 1),
	}

	client := pb.NewITUDatabaseClient(conn)

	fmt.Println("Electing... ")

	// Call the Election method and handle errors
	MaxClient, err = client.Election(ctx, req)

	fmt.Println("Electing...1 ")
	
	if err != nil {
		log.Fatalf("Election call failed: %v", err)
		return
	}

	if MaxClient != nil {
		Max = MaxClient.MAX
	} else {
		log.Fatalf("MaxClient response is nil")
		return
	}

	select{}
}

func startGRPCServer(port int) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Printf("Failed to listen on port %d: %v\n", port, err)
		return
	}
	grpcServer := grpc.NewServer()

	pb.RegisterITUDatabaseServer(grpcServer, &TypeNode{})

	fmt.Printf("Starting gRPC server on port %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC server: %v\n", err)
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

func (N1 *TypeNode) Election(ctx context.Context, re1 *pb.RequestElection) (*pb.Answer, error) {
	N1.mu.Lock()
	defer N1.mu.Unlock()

	if !isPortAvailable(int(re1.ClientName)) {
		if !isPortAvailable(int(re1.ClientName + 1)) {
			fmt.Println("1")
			if (Max != Node) {
				N1.Coordinator(ctx, &pb.IAmCoordinator{LamportTime: 0, ClientName: 5051})
			}

			return &pb.Answer{LamportTime: 0, MAX: re1.ClientName}, nil
		} else {
			fmt.Println("Skipped a faulty node")
			conn, _ := grpc.Dial(strconv.Itoa(int(re1.ClientName+1)), grpc.WithInsecure())
			
			receivingClient := pb.NewITUDatabaseClient(conn)
			req := &pb.RequestElection{
				LamportTime: 0,
				ClientName:  int32(Node + 1),
			}

			return receivingClient.Election(ctx, req)

		}

	} else if re1.ClientName == Node {
		if (Max != Node) {
			N1.Coordinator(ctx, &pb.IAmCoordinator{LamportTime: 0, ClientName: 5051})
		}
		return &pb.Answer{LamportTime: 0, MAX: re1.ClientName}, nil
	} else {
		conn, _ := grpc.Dial(strconv.Itoa(int(Node+1)), grpc.WithInsecure())
		
		receivingClient := pb.NewITUDatabaseClient(conn)
		req := &pb.RequestElection {
			LamportTime: 0,
			ClientName:  int32(Node + 1),
		}

		return receivingClient.Election(ctx, req)
	}
}

func (N1 *TypeNode) Coordinator(ctx context.Context, re1 *pb.IAmCoordinator) (*pb.SendsAllegiance, error) {
	N1.mu.Lock()
	defer N1.mu.Unlock()

	if !isPortAvailable(int(re1.ClientName)) {
		if !isPortAvailable(int(re1.ClientName + 1)) {
			fmt.Println("fail1")
			
		} else if((re1.ClientName + 1) == Node) {
			return &pb.SendsAllegiance{LamportTime: 0, ClientName: re1.ClientName}, nil

		} else {
			fmt.Println("Skipped a faulty node")

			if re1.ClientName == Node+1 {
				return &pb.SendsAllegiance{LamportTime: 0, ClientName: Node}, nil
			}

			conn, err := grpc.Dial(strconv.Itoa(int(Node+2)), grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Failed to connect")
			}

			receivingClient := pb.NewITUDatabaseClient(conn)
			req := &pb.IAmCoordinator{
				LamportTime: 0,
				ClientName:  re1.ClientName,
			}

			//Here is the new coordinator:
			Max = re1.ClientName

			return receivingClient.Coordinator(ctx, req)
		}
	} else if re1.ClientName == Node {
		return &pb.SendsAllegiance{LamportTime: 0, ClientName: Node - 1}, nil
	} else {

		conn, err := grpc.Dial(strconv.Itoa(int(Node+1)), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect")
		}

		receivingClient := pb.NewITUDatabaseClient(conn)
		req := &pb.IAmCoordinator{
			LamportTime: 0,
			ClientName:  int32(Node + 1),
		}

		return receivingClient.Coordinator(ctx, req)
	}

	//Here is the new coordinator:
	Max = re1.ClientName

	return &pb.SendsAllegiance{LamportTime: 0, ClientName: 0}, nil
}