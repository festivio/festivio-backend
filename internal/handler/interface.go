package handler

import "github.com/gin-gonic/gin"

type HandlerInterface interface {
	// Authorization methods
	SignUpUser() gin.HandlerFunc
	SignInUser() gin.HandlerFunc
	LogOutUser() gin.HandlerFunc
	// User methods
	GetMe() gin.HandlerFunc
}
