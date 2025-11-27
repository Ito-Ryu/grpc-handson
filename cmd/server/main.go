package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/Ito-Ryu/grpc-handson/pkg/time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedTimeServiceServer
}

func (s *server) GetCurrentTime(_ context.Context, in *pb.GetCurrentTimeRequest) (*pb.GetCurrentTimeResponse, error) {
	now := time.Now()
	log.Printf("response to %s", now)
	return &pb.GetCurrentTimeResponse{Date: now.String()}, nil
}

func (s *server) GetCurrentTimeStream(in *pb.GetCurrentTimeRequest, stream pb.TimeService_GetCurrentTimeStreamServer) error {
	count := 5
	for i := 0; i < count; i++ {
		if err := stream.Send(&pb.GetCurrentTimeResponse{
			Date: time.Now().String(),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
