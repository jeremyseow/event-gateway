package event

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeremyseow/event-gateway/server/http/handler/event"
)

func SetupEventRoutes(router fiber.Router, eventAPI event.EventAPIIface) {
	eventRouter := router.Group("/events")
	eventRouter.Post("/track", eventAPI.Track)
}
