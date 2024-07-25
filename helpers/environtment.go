package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvirontmentVariable struct {
	ApiEnv        string
	DBName        string
	DBUser        string
	DBHost        string
	DBPass        string
	DBPort        string
	CorsWhitelist string
	ClientOrigin  string
	JwtSecretKey  string
}

func Environtment() (ev EnvirontmentVariable) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error .env failed to load")
	}

	ev = EnvirontmentVariable{}
	ev.ApiEnv = os.Getenv("API_ENV")
	ev.DBHost = os.Getenv("DBHOST")
	ev.DBName = os.Getenv("DBNAME")
	ev.DBPass = os.Getenv("DBPASS")
	ev.DBUser = os.Getenv("DBUSER")
	ev.DBPort = os.Getenv("DBPORT")
	ev.CorsWhitelist = os.Getenv("CORS_WHITELIST")
	ev.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	ev.ClientOrigin = os.Getenv("CLIENT_ORIGIN")

	return ev
}
