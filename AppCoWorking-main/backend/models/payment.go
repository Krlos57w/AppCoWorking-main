package models

import "time"

// Payment representa los pagos asociados a las reservas.
type Payment struct {
	ID            uint      `json:"id"`
	ReservationID uint      `json:"reservation_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
