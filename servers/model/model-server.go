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
	score := analyzeText(text)
	res := strconv.Itoa(score)

	return &pb.AnalyzeReply{Message: "Sentiment Analyzed Score: " + res}, nil
}

func reverseText(s string) string {
	textSlice := strings.Split(s, " ")
	reversedText := ""
	for i := len(textSlice); i > 0; i-- {
		reversedText += textSlice[i-1] + " "
	}
	return reversedText
}

func analyzeText(s string) int {
	return len(s)
}

// func analyzeText(s string) string {
// 	model, err := sentiment.Train()
// 	if err != nil {
// 		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
// 	}

// 	analysis := model.SentimentAnalysis(s, sentiment.English)
// 	fmt.Println(analysis)
// 	return strconv.Itoa(int(analysis.Score))
// }

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
