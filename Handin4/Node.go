package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"net"
	"strconv"
	"sync"

	pb "github.com/BastianGram/Distibuted-Systems/Handin4/Handin4/grpc"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials/insecure"
)

// Server struct to implement the MyServiceServer interface
type TypeNode struct {
	pb.UnimplementedITUDatabaseServer

	//clients map[string]*Client
	mu      sync.Mutex // to protect access to clients map
}

var Node int32; // equal to it's port
var Max int = -1;

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

	Node := findAvailablePort(5050)
	fmt.Printf("Found available port: %d\n", Node)
	//conn, err := grpc.NewClient(ctx, "localhost:%d" + strv.conv.Itoa(port))
	conn , err := grpc.DialContext(ctx, "localhost:" + strconv.Itoa(Node) , grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}

    //client := pb.NewITUDatabaseClient(conn)

	Max = Election(Node)



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


func (N1 *TypeNode) Election(ctx context.Context, re1 *pb.RequestElection) (*pb.Answer, error) {
	N1.mu.Lock()
    defer N1.mu.Unlock()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", re1.ClientName))
		if err == nil {
			ln.Close() // Close the listener since we only wanted to check availability

			conn, err := grpc.Dial(strconv.Itoa(int(Node + 1)), grpc.WithInsecure())
    		if err != nil {
        		log.Fatalf("Failed to connect")
    		}

			receivingClient := pb.NewITUDatabaseClient(conn)
			req := *pb.RequestElection {
       			LamportTime: 0,
        		ClientName: int32(Node + 1),
    		}
			
			return receivingClient.Election(ctx, req)
    		
			/*
			//resp, err := receivingClient.Election(ctx, req)

			conn.Close()
    		if err != nil {
        		
			log.Printf("Error requesting access from ")
			*/
    		}
		} else {// If port is taken, it tries the next port

		}
	return pb.IAmCoordinator{LamportTime: 0, ClientName: 0}, nil
}


func Coordinator() {
	
}