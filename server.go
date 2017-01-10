// +build server

package main

import (
	"log"
	"net"

	pb "github.com/dictav/test-go-grpc-protocolbuffers/grpcexample"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type handler struct{}

func buildPerson(n int32) *pb.Person {
	plen := int(n%3 + 1)
	phones := make([]*pb.Person_PhoneNumber, plen)
	for i := 0; i < plen; i++ {
		phones[i] = &pb.Person_PhoneNumber{
			Number: "phone",
			Type:   pb.Person_PhoneType(i),
		}
	}

	return &pb.Person{
		Id:    n,
		Name:  "name",
		Email: "email",
		Phone: phones,
	}
}

func (h *handler) GetPerson(ctx context.Context, req *pb.Request) (*pb.Person, error) {
	return buildPerson(0), nil
}

func (h *handler) ListPeople(req *pb.Request, stream pb.GRPCExample_ListPeopleServer) error {
	for i := int32(0); i < 100; i++ {
		if err := stream.Send(buildPerson(i)); err != nil {
			return err
		}
	}
	return nil
}

func serve() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	h := &handler{}
	pb.RegisterGRPCExampleServer(s, h)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serve()
}
