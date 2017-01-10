// +build client

package main

import (
	"context"
	"log"

	pb "github.com/dictav/test-go-grpc-protocolbuffers/grpcexample"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewGRPCExampleClient(conn)

	req := &pb.Request{}
	ctx := context.Background()
	person, err := client.GetPerson(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Person.Name", person.Name)

	stream, err := client.ListPeople(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		person, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Person.Id", person.Id)
	}
}
