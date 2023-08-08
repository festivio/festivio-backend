package handler

import (
	"net/http"

	"github.com/festivio/festivio-backend/domain"
	"github.com/festivio/festivio-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h handler) SignUpUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signUpInput *domain.SignUpInput
		var errorStruct domain.ErrorStruct

		if err := ctx.ShouldBindJSON(&signUpInput); err != nil {
			errorStruct.Error.Code = http.StatusBadRequest
			errorStruct.Error.Message = "Invalid request body."
			ctx.JSON(http.StatusBadRequest, &errorStruct)
			return
		}

		if signUpInput.Password != signUpInput.PasswordConfirm {
			errorStruct.Error.Code = http.StatusBadRequest
			errorStruct.Error.Message = "Passwords do not match."
			ctx.JSON(http.StatusBadRequest, &errorStruct)
			return
		}

		err := h.srv.SignUpUser(signUpInput)
		if err != nil {
			switch err.Error() {
			case "user with this email address already exists":
				errorStruct.Error.Code = http.StatusConflict
				errorStruct.Error.Message = "User with this email address already exists."
				ctx.JSON(http.StatusConflict, &errorStruct)
				return
			case "something bad happened":
				errorStruct.Error.Code = http.StatusBadGateway
				errorStruct.Error.Message = "Something bad happened."
				ctx.JSON(http.StatusBadGateway, &errorStruct)
				return
			}
		}

		ctx.JSON(http.StatusCreated, gin.H{})
	}
}

func (h handler) SignInUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signInInput *domain.SignInInput
		var errorStruct domain.ErrorStruct

		if err := ctx.ShouldBindJSON(&signInInput); err != nil {
			errorStruct.Error.Code = http.StatusBadRequest
			errorStruct.Error.Message = "Invalid request body."
			ctx.JSON(http.StatusBadRequest, &errorStruct)
			return
		}

		user, err := h.srv.SignInUser(signInInput)
		if err != nil {
			switch err.Error() {
			case "invalid email":
				errorStruct.Error.Code = http.StatusBadRequest
				errorStruct.Error.Message = "Invalid email."
				ctx.JSON(http.StatusBadRequest, &errorStruct)
				return
			case "invalid password":
				errorStruct.Error.Code = http.StatusBadRequest
				errorStruct.Error.Message = "Invalid password."
				ctx.JSON(http.StatusBadRequest, &errorStruct)
				return
			}
		}

		token, err := utils.GenerateToken(h.cfg.Token.ExpiredIn, user.ID, h.cfg.Token.JwtTokenSecret)
		if err != nil {
			errorStruct.Error.Code = http.StatusBadRequest
			errorStruct.Error.Message = "Failed to generate token."
			ctx.JSON(http.StatusBadRequest, &errorStruct)
			return
		}

		data := struct {
			Token string `json:"token"`
		}{token}

		ctx.SetCookie("token", token, h.cfg.Token.MaxAge*60, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, gin.H{"Data": data})
	}
}

func (h handler) LogOutUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := struct {
			Message string `json:"message"`
		}{"You have successfully logged out."}
		ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, gin.H{"Data": data})
	}
}
