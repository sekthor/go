package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/sekthor/go/grpc/auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	pb.UnimplementedAccountServer
}

func (s *Server) UserDetails(ctx context.Context, req *pb.UserIdRequest) (*pb.User, error) {
	return &pb.User{
		Id:       req.GetId(),
		Username: "username",
		Email:    "user@domain.com",
	}, nil
}

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	auth := md["authorization"]
	/*
		if !valid() {
			return nil, errors.New("invalid token")
		}
	*/
	log.Println(auth)
	return handler(ctx, req)
}

func main() {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(Interceptor),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterAccountServer(s, &Server{})
	log.Println("starting server on port " + port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
