// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"

	pb "github.com/grebnesorbocaj/go-text-saver/grpc/model"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedAnalyzerServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Analyze(ctx context.Context, in *pb.AnalyzeRequest) (*pb.AnalyzeReply, error) {
	log.Printf("Received: %v", in.GetText())
	text := in.GetText()
	textSlice := strings.Split(text, " ")
	sliceLength := strconv.Itoa(len(textSlice))
	return &pb.AnalyzeReply{Message: sliceLength + " words in text."}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAnalyzerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
