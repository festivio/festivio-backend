package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	signUp = "/sign-up"
	signIn = "/sign-in"
	logOut = "/log-out"
)

func MapRoutes(group *gin.RouterGroup, h HandlerInterface, db *gorm.DB) {
	// Authorization routes
	group.POST(signUp, h.SignUpUser())
	group.POST(signIn, h.SignInUser())
	group.POST(logOut, h.LogOutUser())
}
