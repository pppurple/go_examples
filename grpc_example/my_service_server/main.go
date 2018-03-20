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
	alice := pb.Person{Name: "alice", Age: 20, Country: pb.Person_AMERICA, Hobby: "tennis"}

	fmt.Println(alice)
	return &pb.SearchResponse{People: nil}, nil
	//people := []pb.Person{alice}
	//return &pb.SearchResponse{People: people}, nil
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

/*
type SearchResponse struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResponse) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

type Person struct {
	Name    string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age     int32          `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Country Person_Country `protobuf:"varint,3,opt,name=country,enum=exampleGrpc.Person_Country" json:"country,omitempty"`
	Hobby   string         `protobuf:"bytes,4,opt,name=hobby" json:"hobby,omitempty"`
}
*/
