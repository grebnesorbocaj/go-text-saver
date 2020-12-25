// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grebnesorbocaj/go-text-saver/grpc/model"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultText = ""
	port        = 8008
)

func callServer(s string, c chan string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAnalyzerClient(conn)

	// Contact the server and print out its response.
	if len(s) > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.Analyze(ctx, &pb.AnalyzeRequest{Text: s})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Received: %s", r.GetMessage())
	} else {
		log.Print("Please add some string to the request")
	}
	c <- s
}
