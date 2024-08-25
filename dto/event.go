package dto

type Event struct {
	Name      string   `json:"name"`
	EventTime int64    `json:"event_time"`
	Entities  []Entity `json:"entities"`
}

type Events struct {
	Events []Event `json:"events"`
}
