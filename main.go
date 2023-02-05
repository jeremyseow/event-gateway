package main

import (
	"log"
	"net"

	"github.com/jeremyseow/event-gateway/config"
	"github.com/jeremyseow/event-gateway/internal/handler"
	"github.com/jeremyseow/event-gateway/internal/storage/file"
	"github.com/jeremyseow/event-gateway/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	logger.Info("logger created")

	listener, err := net.Listen("tcp", config.Conf.GrpcPort)
	if err != nil {
		logger.Sugar().Fatalf("error starting listener: %s", err.Error())
		panic(err)
	}
	logger.Info("listener started")

	grpcServer := grpc.NewServer()
	fw := file.NewWriter(logger)

	pb.RegisterEventServiceServer(grpcServer, &handler.EventAPI{Storage: fw, Logger: logger})
	reflection.Register(grpcServer)
	logger.Info("server registered and starting")

	if err := grpcServer.Serve(listener); err != nil {
		logger.Sugar().Fatalf("error serving: %s", err.Error())
		panic(err)
	}
}
