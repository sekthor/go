package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/sekthor/go/grpc/auth/pb"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
)

func getToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}

func main() {
	port := os.Getenv("PORT")

	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(getToken())}
	conn, err := grpc.NewClient(
		"localhost:"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(perRPC),
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
	log.Printf("Greeting: %s", r.GetUsername())
}
