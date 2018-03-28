package main

import (
	"fmt"
	"time"

	pb "github.com/pppurple/go_examples/grpc_example/my_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	host  = "localhost"
	port  = ":50051"
	query = "Wall"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(host+port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewPeopleServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SearchByName(ctx, &pb.SearchRequest{Query: query})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Println("[Response]")
	for _, v := range res.People {
		fmt.Println(v)
	}
}
