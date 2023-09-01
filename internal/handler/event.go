package handler

import (
	"context"

	"github.com/jeremyseow/event-gateway/internal/storage"
	"github.com/jeremyseow/event-gateway/internal/utils"
	"github.com/jeremyseow/event-gateway/pb"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const (
	logTag = "handler.EventAPI"
)

type EventAPI struct {
	pb.UnimplementedEventServiceServer
	Storage storage.Storage
	Logger  *zap.Logger
}

func NewEventAPI(writer storage.Storage, allUtilityClients *utils.AllUtilityClients) *EventAPI {
	return &EventAPI{Storage: writer, Logger: allUtilityClients.Logger}
}

func (eventAPI *EventAPI) SendEvent(_ context.Context, request *pb.EventRequest) (*pb.EventResponse, error) {
	eventAPI.Logger.Info("received events", zap.String("logTag", logTag), zap.Int("size of request", len(request.Events)), zap.String("function", "SendEvent"))

	bytes, err := proto.Marshal(request)
	if err != nil {
		eventAPI.Logger.Error("error when unmarshalling", zap.String("logTag", logTag), zap.String("function", "SendEvent"))
		return &pb.EventResponse{Result: "Failed"}, err
	}

	_, err = eventAPI.Storage.Write(bytes)
	if err != nil {
		eventAPI.Logger.Error("error when writing events", zap.String("logTag", logTag), zap.String("function", "SendEvent"))
		return &pb.EventResponse{Result: "Failed"}, err
	}

	return &pb.EventResponse{Result: "Success"}, nil
}
