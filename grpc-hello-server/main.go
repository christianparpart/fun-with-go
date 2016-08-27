package main

import (
	"errors"
	"log"
	"net"

	"github.com/christianparpart/fun-with-go/grpc-hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type MyServer struct {
}

func (s *MyServer) SayHello(c context.Context, r *hello.HelloRequest) (*hello.HelloResponse, error) {
	if r.Greeting == "naknak" {
		return nil, errors.New("Gna!")
	}

	response := &hello.HelloResponse{
		Reply: "Cheers, " + r.Greeting,
	}

	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &MyServer{})

	s.Serve(listener)
}
