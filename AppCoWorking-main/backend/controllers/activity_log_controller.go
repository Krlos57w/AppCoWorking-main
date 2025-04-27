package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Estructura para recibir datos del log
type CreateLogInput struct {
	UserID   uint   `json:"user_id" binding:"required"`
	Accion   string `json:"accion" binding:"required"`
	Detalles string `json:"detalles"`
}

// Función para crear un log desde una solicitud HTTP
func CreateLog(c *gin.Context) {
	var input CreateLogInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log := models.ActivityLog{
		UserID:   input.UserID,
		Accion:   input.Accion,
		Detalles: input.Detalles,
	}

	config.DB.Create(&log)

	c.JSON(http.StatusOK, gin.H{"message": "Log registrado correctamente"})
}

// Función para obtener todos los logs
func GetLogs(c *gin.Context) {
	var logs []models.ActivityLog
	config.DB.Find(&logs)
	c.JSON(http.StatusOK, logs)
}

// Función para registrar logs de actividad (se usa internamente en el sistema)
func CreateActivityLog(userID uint, accion, detalles string) {
	log := models.ActivityLog{
		UserID:    userID,
		Accion:    accion,
		Detalles:  detalles,
		CreatedAt: time.Now(), // GORM lo maneja automáticamente
	}
	config.DB.Create(&log)
}
