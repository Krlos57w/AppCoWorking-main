package models

type Space struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Nombre        string  `json:"nombre"`
	Descripcion   string  `json:"descripcion"`
	Capacidad     int     `json:"capacidad"`
	PrecioPorHora float64 `json:"precio_por_hora"`
	Estado        string  `json:"estado"`
	Ubicacion     string  `json:"ubicacion"` // ✅ Agregado
}
