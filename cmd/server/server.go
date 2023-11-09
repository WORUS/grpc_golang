package server

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/WORUS/grpc_video-service/pkg/video_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = 8080

type server struct {
	desc.UnimplementedVideoServer
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Video id: %d", req.GetId())

	return &desc.GetResponse{
		Info: &desc.VideoInfo{
			Id:   req.GetId(),
			Name: "Zoo",
		},
	}, nil
}

func StartServer() {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterVideoServer(s, &server{})

	log.Printf("server listening at %d", list.Addr())

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
