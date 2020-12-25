// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/grebnesorbocaj/go-text-saver/grpc/model"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultText = ""
	port        = 8008
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	texts := []string{
		"somehow got this text",
		"this is another text",
		"jinx is an adc",
		"shen is a tank",
		"something something",
		"this is a text of text",
		"i like playing shen",
	}

	c := make(chan string)

	for _, text := range texts {
		go callServer(text, c)
	}

	for l := range c {
		go func(text string) {
			time.Sleep(3 * time.Second)
			callServer(text, c)
		}(l)
	}

	fmt.Fprintf(w, "Thanks for the query:)")
}

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
