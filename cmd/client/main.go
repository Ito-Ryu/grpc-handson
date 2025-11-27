package main

import (
	"context"
	"fmt"

	pb "github.com/Ito-Ryu/grpc-handson/pkg/time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "localhost:8080"
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return
	}

	defer conn.Close()

	client := pb.NewTimeServiceClient(conn)
	GetCurrentTime(client)
}

func GetCurrentTime(client pb.TimeServiceClient) {
	req := &pb.GetCurrentTimeRequest{}
	res, err := client.GetCurrentTime(context.Background(), req)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetDate())
	}
}
