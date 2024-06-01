package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/sekthor/go/grpc/interceptor/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	token := md["authorization"]
	// TODO: this is where we would check the token

	log.Printf("token: %s", token)
	return handler(ctx, req)
}

func main() {
	port := os.Getenv("PORT")
	certPath := os.Getenv("CERT_PATH")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cert, err := tls.LoadX509KeyPair(certPath+"/ca-cert.pem", certPath+"/ca-key.pem")
	if err != nil {
		log.Fatalf("could not load certificate: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(Interceptor),
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(opts...)

	pb.RegisterAccountServer(s, &Server{})
	log.Println("starting server on port " + port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
