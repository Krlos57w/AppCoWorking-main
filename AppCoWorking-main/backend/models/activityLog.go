package models

import "time"

type ActivityLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Accion    string    `json:"accion"`
	Detalles  string    `json:"detalles"`  // Antes: Detalle
	CreatedAt time.Time `json:"created_at"` // Antes: FechaHora (GORM maneja esto automáticamente)
}
