package models

// Message struct handle response
type Message struct {
	Channel string `json:"channel"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Message string `json:"message"`
}
