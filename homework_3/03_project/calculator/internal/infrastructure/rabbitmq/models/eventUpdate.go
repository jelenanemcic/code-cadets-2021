package models

// EventUpdate represents a DTO for received event updates.
type EventUpdate struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}
