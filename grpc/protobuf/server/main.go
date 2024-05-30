package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/sekthor/go/grpc/protobuf/pb"
	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) GreetUser(ctx context.Context, u *pb.User) (*pb.Greeting, error) {
	return &pb.Greeting{
		Msg: fmt.Sprintf("Hello %s", u.Name),
	}, nil
}

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	log.Println("starting server on port " + port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
