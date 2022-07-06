package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/Romero027/echo-server/pb"
	"google.golang.org/grpc"
)

type server struct {
	port string
}

func (s *server) echo(ctx context.Context, x *pb.Message) (*pb.Message, error) {
	log.Printf("[%s] got: [%s]", s.port, x.GetMsg())
	return x, nil
}

func runServer(port string) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterEchoServiceServer(s, &server{port})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
