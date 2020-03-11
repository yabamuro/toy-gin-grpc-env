package main

import (
	pb "ginApl"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	port    = ":50051"
	logfile = "/var/log/apl/ginApl.log"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Send(ctx context.Context, in *pb.GrpcRequest) (*pb.GrpcResponse, error) {
	respMsg := "Rcv(gRPC) " + in.Query + " on " + time.Now().String()
	return &pb.GrpcResponse{Message: respMsg}, nil
}

// LoggingSetting
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings(logfile)
	log.Println("gRPCServer Start")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGrpcServer(s, &server{})
	s.Serve(lis)
}
