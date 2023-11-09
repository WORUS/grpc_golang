package client

import (
	"context"
	"log"
	"time"

	desc "github.com/WORUS/grpc_video-service/pkg/video_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080"
	userID  = 12
)

func StartClient() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("faield to connect to server: %v", err)
	}

	defer conn.Close()

	c := desc.NewVideoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: 10})
	if err != nil {
		log.Fatalf("faield to get video info: %v", err)
	}

	log.Printf("Video info: %v", r.GetInfo())

}
