package main

import (
	"context"
	"log"
	"time"

	"github.com/rhythmicsoul/nginx-mgmt/proto/controller"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot Connect to Server: %v", err)
	}
	defer conn.Close()

	c := controller.NewAddServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, _ := c.NewAgentToken(ctx, &controller.Empty{})
	log.Printf("the token is: %s", r.GetToken())
}
