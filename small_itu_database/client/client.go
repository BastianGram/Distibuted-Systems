package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/BastianGram/Distibuted-Systems/tree/handin3v2/small_itu_database/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}

	client := pb.NewITUDatabaseClient(conn)

	students, err := client.GetStudents(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatalf("Not working")
	}

	for _, student := range students.Students {
		println(" - " + student)
	}
}
