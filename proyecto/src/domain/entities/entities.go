package entities

import "time"

// Event represents a sensor event
type Event struct {
	EventID   int       `json:"event_id"`   // ID autoincremental
	Type      string    `json:"type"`       // Type of event
	CreatedAt time.Time `json:"created_at"` 
	Unit      string    `json:"unit"`       // Measurement unit (°C, %, etc.)
}
