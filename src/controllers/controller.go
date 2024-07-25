package controllers

import (
	"pgd-server.com/middleware"
	"pgd-server.com/src/services"
)

var userService = services.UserService{}

// Add middleware here
var authMiddleware = middleware.AuthMiddleware{}
var accessTokenService = services.AccessTokenService{}
var customerService = services.CustomerService{}
