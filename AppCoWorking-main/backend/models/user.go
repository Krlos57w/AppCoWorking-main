package models

import "time"

type User struct {
	ID            uint      `json:"id"`
	Nombre        string    `json:"nombre"`
	Email         string    `json:"email"`
	Password      string    `json:"-"`
	RolID         uint      `json:"rol_id"`
	FechaRegistro time.Time `json:"fecha_registro"`
	Role          Role      `json:"role" gorm:"foreignKey:RolID"`
}
