package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	pb "github.com/BastianGram/Distibuted-Systems/tree/Handin3/small_itu_database/grpc"
)

type ITU_databaseServer struct {
	proto.UnimplementedITUDatabaseServer
	students []string
}

func (s *ITU_databaseServer) GetStudents(ctx context.Context, in *proto.Empty) (*proto.Students, error) {
	return &proto.Students{Students: s.students}, nil
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
