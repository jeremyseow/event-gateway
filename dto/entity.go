package dto

type Entity struct {
	Schema     string                 `json:"schema"`
	Version    string                 `json:"version"`
	Parameters map[string]interface{} `json:"parameters"`
}
