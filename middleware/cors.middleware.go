package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsMiddleware struct{}

func (cm *CorsMiddleware) Cors() gin.HandlerFunc {
	// if helpers.Environtment().ApiEnv == "development" {
	// 	return cors.New(cors.Config{
	// 		AllowOrigins: []string{"*"},
	// 		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS"},
	// 		AllowHeaders: []string{"*"},
	// 	})
	// }
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS"},
		AllowHeaders: []string{"*"},
	})
	// for production only
	// corsWhitelist := helpers.Environtment().CorsWhitelist
	// corsWhitelisted := strings.Split(corsWhitelist, ",")

	// return cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS"},
	// 	AllowHeaders: []string{"*"},
	// })
}
