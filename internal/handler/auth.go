package handler

import (
	"net/http"

	"github.com/festivio/festivio-backend/domain"
	"github.com/festivio/festivio-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary User SignUp
// @Tags Auth
// @Description User signUp
// @ModuleID userSignUp
// @Accept  json
// @Produce json
// @Param input body domain.SignUpInput true "sign-up info"
// @Success 201 {string} string "ok"
// @Failure 400,409 {object} domain.ErrorStruct
// @Failure 502 {object} domain.ErrorStruct
// @Failure default {object} domain.ErrorStruct
// @Router /sign-up [post]
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

		ctx.Status(http.StatusCreated)
	}
}

// @Summary User SignIn
// @Tags Auth
// @Description User signIn
// @ModuleID userSignIn
// @Accept  json
// @Produce json
// @Param input body domain.SignInInput true "sign-in info"
// @Success 200 {object} domain.SignInResponse
// @Failure 400 {object} domain.ErrorStruct
// @Failure default {object} domain.ErrorStruct
// @Router /sign-in [post]
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

		response := &domain.SignInResponse{
			Data: struct {
				Token string `json:"token"`
			}{
				Token: token,
			},
		}

		ctx.SetCookie("token", token, h.cfg.Token.MaxAge*60, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, response)
	}
}

// @Summary User LogOut
// @Tags Auth
// @Description User logOut
// @Security ApiKeyAuth
// @ModuleID userLogOut
// @Accept  json
// @Produce json
// @Success 200 {object} domain.MessageResponse
// @Router /log-out [post]
func (h handler) LogOutUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := &domain.MessageResponse{
			Data: struct {
				Message string `json:"message"`
			}{
				Message: "You have successfully logged out.",
			},
		}
		ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
		ctx.JSON(http.StatusOK, response)
	}
}
