package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Obtener todas las reservas
func GetBookings(c *gin.Context) {
	var bookings []models.Booking
	config.DB.Preload("User").Preload("Space").Find(&bookings)
	c.JSON(http.StatusOK, bookings)
}

// Crear reserva
func CreateBooking(c *gin.Context) {
	var booking models.Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&booking)
	c.JSON(http.StatusOK, booking)
}

// Obtener reserva por ID
func GetBookingByID(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking

	if err := config.DB.Preload("User").Preload("Space").First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reserva no encontrada"})
		return
	}

	c.JSON(http.StatusOK, booking)
}

// Actualizar reserva
func UpdateBooking(c *gin.Context) {
	id := c.Param("id")
	var booking models.Booking

	if err := config.DB.First(&booking, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reserva no encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&booking)
	c.JSON(http.StatusOK, booking)
}

// Eliminar reserva
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Booking{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Reserva eliminada"})
}
