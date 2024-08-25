package event

import (
	"fmt"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/jeremyseow/event-gateway/dto"
	"github.com/jeremyseow/event-gateway/processor/schema"
	"go.uber.org/zap"
)

const (
	logTag = "handler.EventAPI"
)

type EventAPIIface interface {
	Track(ctx *fiber.Ctx) error
}

type EventAPI struct {
	logger    *zap.Logger
	validator *schema.Validator
}

func NewEventAPI(logger *zap.Logger, validator *schema.Validator) *EventAPI {
	return &EventAPI{
		logger:    logger,
		validator: validator,
	}
}

func (e *EventAPI) Track(ctx *fiber.Ctx) error {
	eventsDTO := &dto.Events{}
	err := ctx.BodyParser(&eventsDTO)
	if err != nil {
		e.logger.Error("error parsing request body", zap.String("logTag", logTag), zap.String("function", "Track"), zap.String("error", err.Error()))
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	for _, event := range eventsDTO.Events {
		for _, entity := range event.Entities {
			// TODO: make this async
			err := e.validator.ValidateEntity(entity)
			if err != nil {
				// should write to DLQ?
				fmt.Println(err)
			}
		}
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "OK"})
}
