package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Variable global para el secreto JWT
var JWTSecret string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error cargando el archivo .env, usando valores por defecto")
	}

	JWTSecret = os.Getenv("JWT_SECRET")

	if JWTSecret == "" {
		log.Fatal("JWT_SECRET no está configurado en el archivo .env")
	}
}
