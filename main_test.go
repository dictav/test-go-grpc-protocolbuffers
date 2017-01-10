package main

import (
	"context"
	"io"
	"log"
	"testing"

	pb "github.com/dictav/test-go-grpc-protocolbuffers/grpcexample"
	"google.golang.org/grpc"
)

var client pb.GRPCExampleClient

func TestMain(m *testing.M) {
	go serve()
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = pb.NewGRPCExampleClient(conn)

	m.Run()
}

func BenchmarkGetPerson(b *testing.B) {
	ctx := context.Background()
	req := &pb.Request{}

	for i := 0; i < b.N; i++ {
		_, err := client.GetPerson(ctx, req)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkListPerson(b *testing.B) {
	ctx := context.Background()
	req := &pb.Request{}
	names := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		strm, err := client.ListPeople(ctx, req)
		if err != nil {
			b.Error(err)
		}

		for {
			res, err := strm.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				b.Error(err)
				continue
			}

			names[i] = res.Name
		}
	}
}
