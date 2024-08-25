package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeremyseow/event-gateway/server/http/handler"
	eventRouter "github.com/jeremyseow/event-gateway/server/http/router/event"
)

func SetupRoutes(app *fiber.App, appAPIs *handler.AllAPIs) {
	apiRouter := app.Group("/api")
	eventRouter.SetupEventRoutes(apiRouter, appAPIs.EventAPI)
}
