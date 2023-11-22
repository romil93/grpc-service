package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/romil93/grpc-service/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewLoggerClient(conn)
	loggerMessage := pb.LoggerMessage{UseCase: "test", Log: "test"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	status, err := client.LogMessage(ctx, &loggerMessage)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}

	log.Printf("status code: %d errors: %s", status.GetStatusCode(), status.GetErrors())
}
