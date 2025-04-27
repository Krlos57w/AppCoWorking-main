package models

import "time"

// Reservation representa una reserva de espacio.
type reservation struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	SpaceID   uint      `json:"space_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
