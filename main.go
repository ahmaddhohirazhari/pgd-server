package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pgd-server.com/config"
	"pgd-server.com/helpers"
	"pgd-server.com/middleware"
	"pgd-server.com/migrate"
	"pgd-server.com/src/controllers"
)

func main() {
	ev := helpers.Environtment()
	iConnectDB := config.IConnectDB{
		DBHost: ev.DBHost,
		DBUser: ev.DBUser,
		DBName: ev.DBName,
		DBPass: ev.DBPass,
		DBPort: ev.DBPort,
	}
	config.ConnectDB(iConnectDB)

	migrate.MigrateDB()
	r := gin.Default()

	// Controller
	var userController = controllers.UserController{}
	var authController = controllers.AuthController{}
	var customerController = controllers.CustomerController{}

	// Middleware
	var authMiddleware = middleware.AuthMiddleware{}
	var corsMiddleware = middleware.CorsMiddleware{}

	if ev.ApiEnv == "development" {
		r.Use(corsMiddleware.Cors())
	} else {
		r.Use(corsMiddleware.Cors())
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"version": "v1.13",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.GeneralResponse{
			Status:  false,
			Message: "404 not found",
			Data:    nil,
		})
	})

	// Auth
	r.POST("/api/v1/login", authController.SignIn)

	authorizeRouter := r.Group("/api/v1")
	authorizeRouter.Use(authMiddleware.JwtTokenCheck)

	// Users
	authorizeRouter.POST("/users", userController.Create)
	authorizeRouter.GET("/users", userController.FindAll)
	authorizeRouter.GET("/users/:id", userController.FindOne)
	authorizeRouter.PATCH("/users/:id", userController.Update)
	authorizeRouter.DELETE("/users/:id", userController.Delete)
	r.GET("/api/v1/user", userController.FindAll)

	// Customer
	authorizeRouter.POST("/customers", customerController.Create)
	authorizeRouter.GET("/customers", customerController.FindAll)
	authorizeRouter.GET("/customers/:id", customerController.FindOne)
	authorizeRouter.PATCH("/customers/:id", customerController.Update)
	authorizeRouter.DELETE("/customers/:id", customerController.Delete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
