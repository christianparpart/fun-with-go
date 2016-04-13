package main

import (
	"log"
	"os"
	"strconv"

	"github.com/christianparpart/fun-with-go/grpc-hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := hello.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := "default world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	var gen int = 1
	if len(os.Args) > 2 {
		gen, _ = strconv.Atoi(os.Args[2])
	}
	r, err := c.SayHello(context.Background(), &hello.HelloRequest{
		Greeting:   name,
		Generation: uint64(gen),
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Reply)
}
