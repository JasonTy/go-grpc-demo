package main

import (
	"strconv"
	pb "go/grpc/demo/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)



const (
	address     = "localhost:5001"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	c02 := pb.NewBuyServerClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})

	r02, _ := c02.BuyPro(context.Background(), &pb.BuyRequest{Name: "FeiZao"})


	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	log.Printf("Greeting: %s", r02.Message + ":" + strconv.Itoa(int(r02.Price)))
}