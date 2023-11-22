package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	pb "github.com/romil93/grpc-service/logger"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type loggerServer struct {
	pb.UnimplementedLoggerServer
	mu sync.Mutex // protects routeNotes
}

// GetFeature returns the feature at the given point.
func (s *loggerServer) LogMessage(ctx context.Context, message *pb.LoggerMessage) (*pb.Status, error) {
	log.Printf("use case: %s log: %s", message.GetUseCase(), message.GetLog())
	status := pb.Status{}
	status.StatusCode = pb.Status_OK
	status.Errors = []string{}
	return &status, nil
}

func newServer() *loggerServer {
	s := &loggerServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterLoggerServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
