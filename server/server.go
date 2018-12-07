package main

import (
	"log"
	"net"
	pb "go/grpc/demo/protocol"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":5001"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) BuyPro(ctx context.Context, in *pb.BuyRequest) (*pb.BuyReply, error) {
	return &pb.BuyReply{Message: in.Name, Price: 10}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	pb.RegisterBuyServerServer(s, &server{})
	s.Serve(lis)
}