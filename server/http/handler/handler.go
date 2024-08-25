package handler

import (
	"github.com/jeremyseow/event-gateway/processor/schema"
	"github.com/jeremyseow/event-gateway/server/http/handler/event"
	"github.com/jeremyseow/event-gateway/utils"
)

type AllAPIs struct {
	EventAPI event.EventAPIIface
}

func NewAllAPIs(allUtilityClients *utils.AllUtilityClients, validator *schema.Validator) *AllAPIs {
	return &AllAPIs{
		EventAPI: event.NewEventAPI(allUtilityClients.Logger, validator),
	}
}
