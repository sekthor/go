package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/sekthor/go/grpc/protobuf/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	port := os.Getenv("PORT")
	conn, err := grpc.NewClient("localhost:"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.GreetUser(ctx, &pb.User{Name: "sekthor"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMsg())
}
