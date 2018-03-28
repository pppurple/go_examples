package main

import (
	"fmt"
	"net"

	pb "github.com/pppurple/go_examples/grpc_example/my_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SearchByName(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	fmt.Println("query=" + request.Query)

	// DBから検索していると想定
	alice := pb.Person{
		Name:    "Alice Wall",
		Age:     20,
		Country: pb.Person_JAPAN,
		Hobby:   "tennis",
	}
	bobby := pb.Person{
		Name:    "Bobby Wall",
		Age:     33,
		Country: pb.Person_CANADA,
		Hobby:   "music",
	}

	people := []*pb.Person{&alice, &bobby}
	return &pb.SearchResponse{People: people}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPeopleServiceServer(s, &server{})

	fmt.Println("Server started... (port=" + port + ")")

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
