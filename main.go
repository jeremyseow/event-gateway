package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jeremyseow/event-gateway/config"
	"github.com/jeremyseow/event-gateway/pb"
	"github.com/jeremyseow/event-gateway/processor/schema"
	grpcHandler "github.com/jeremyseow/event-gateway/server/grpc/handler"
	httpHandler "github.com/jeremyseow/event-gateway/server/http/handler"
	"github.com/jeremyseow/event-gateway/server/http/router"
	"github.com/jeremyseow/event-gateway/storage/file"
	"github.com/jeremyseow/event-gateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	logTag = "main"
)

func main() {
	// get service configuration
	conf := config.InitConfig()

	// configure clients such as logger, statsd, etc
	allUtilityClients := utils.NewAllUtilityClients(conf)

	// setting up grpc server
	listener, err := net.Listen("tcp", conf.GetString("hostConfig.grpcPort"))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	fw := file.NewWriter(allUtilityClients, file.GetDirName, file.GetFileName)

	pb.RegisterEventServiceServer(grpcServer, grpcHandler.NewEventAPI(fw, allUtilityClients))
	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			panic(err)
		}
	}()

	// setting up http server
	app := fiber.New()

	// configure all APIs
	validator := schema.NewValidator(allUtilityClients.Logger)
	allAPIs := httpHandler.NewAllAPIs(allUtilityClients, validator)
	router.SetupRoutes(app, allAPIs)

	go func() {
		if err := app.Listen(conf.GetString("hostConfig.httpPort")); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // This blocks the main thread until an interrupt is received
	allUtilityClients.Logger.Info("initiating graceful shutdown", zap.String("logTag", logTag), zap.String("function", "main"))

	allUtilityClients.Logger.Info("running cleanup tasks", zap.String("logTag", logTag), zap.String("function", "main"))

	allUtilityClients.Logger.Info("succesfully shutdown", zap.String("logTag", logTag), zap.String("function", "main"))
}
