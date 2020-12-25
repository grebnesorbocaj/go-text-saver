// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/grebnesorbocaj/go-text-saver/grpc/model"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultText = ""
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAnalyzerClient(conn)

	// Contact the server and print out its response.
	text := defaultText
	if len(os.Args) > 1 {
		text = os.Args[1]
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Analyze(ctx, &pb.AnalyzeRequest{Text: text})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Number of words: %s", r.GetMessage())
	} else {
		log.Print("Please add some string to the request")
	}

}
