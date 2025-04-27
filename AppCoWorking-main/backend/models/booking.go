package models

import "time"

type Booking struct {
	ID          uint      `json:"id"`
	UsuarioID   uint      `json:"usuario_id"`
	EspacioID   uint      `json:"espacio_id"`
	FechaInicio time.Time `json:"fecha_inicio"`
	FechaFin    time.Time `json:"fecha_fin"`
	Estado      string    `json:"estado"` // 'pendiente', 'confirmada', 'cancelada', 'completada'
	User        User      `json:"user" gorm:"foreignKey:UsuarioID"`
	Space       Space     `json:"space" gorm:"foreignKey:EspacioID"`
}
