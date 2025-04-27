package main

import (
	"backend/config"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar variables de entorno
	config.LoadEnv()
	
	// Conectar a la base de datos
	config.ConnectDB()

	// Inicializar el router de Gin
	r := gin.Default()

	// Configurar las rutas
	routes.SetupRoutes(r)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
