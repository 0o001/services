package main

import (
	"github.com/micro/services/forex/handler"
	pb "github.com/micro/services/forex/proto"
	"micro.dev/v4/service"
	"micro.dev/v4/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("forex"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterForexHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
