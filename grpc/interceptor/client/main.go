package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/sekthor/go/grpc/interceptor/pb"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func getToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}

func main() {
	port := os.Getenv("PORT")
	certPath := os.Getenv("CERT_PATH")

	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(getToken())}
	creds, err := credentials.NewClientTLSFromFile(certPath+"/ca-cert.pem", "localhost")
	if err != nil {
		log.Fatalf("could not load certificate: %v", err)
	}
	conn, err := grpc.NewClient(
		"localhost:"+port,
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(creds),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.UserDetails(ctx, &pb.UserIdRequest{Id: "thisistheid"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("User Details: %s", r)
}
