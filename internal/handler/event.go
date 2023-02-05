package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/jeremyseow/event-gateway/internal/storage"
	"github.com/jeremyseow/event-gateway/pb"

	"github.com/google/uuid"
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

func (eventAPI *EventAPI) SendEvent(_ context.Context, request *pb.EventRequest) (*pb.EventResponse, error) {
	eventAPI.Logger.Info("received events", zap.String("logTag", logTag), zap.Int("size of request", len(request.Events)))

	bytes, err := proto.Marshal(request)
	if err != nil {
		eventAPI.Logger.Error("error when unmarshalling", zap.String("logTag", logTag), zap.String("error", err.Error()))
		return &pb.EventResponse{Result: "Failed"}, err
	}

	currTime := time.Now().UTC()
	uuid := uuid.New().String()
	filepath := fmt.Sprintf("output/year=%d/month=%d/day=%d/hour=%d", currTime.Year(), currTime.Month(), currTime.Day(), currTime.Hour())
	filename := fmt.Sprintf("%s-%s.txt", "events", uuid)
	err = eventAPI.Storage.Write(filepath, filename, bytes)
	if err != nil {
		eventAPI.Logger.Error("error when writing events", zap.String("logTag", logTag), zap.String("error", err.Error()))
		return &pb.EventResponse{Result: "Failed"}, err
	}

	return &pb.EventResponse{Result: "Success"}, nil
}
