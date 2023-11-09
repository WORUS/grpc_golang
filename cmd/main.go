package main

import (
	"sync"

	"github.com/WORUS/grpc_video-service/cmd/client"
	"github.com/WORUS/grpc_video-service/cmd/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go server.StartServer()
	go client.StartClient()
	wg.Wait()
}
