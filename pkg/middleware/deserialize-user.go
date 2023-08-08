package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/festivio/festivio-backend/config"
	"github.com/festivio/festivio-backend/domain"
	"github.com/festivio/festivio-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeserializeUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var errFlag bool
		var errorStruct domain.ErrorStruct
		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			errorStruct.Error.Code = http.StatusUnauthorized
			errorStruct.Error.Message = "You are not logged in."
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &errorStruct)
			return
		}

		cfg := config.MustLoad()
		sub, err := utils.ValidateToken(token, cfg.Token.JwtTokenSecret)
		if err != nil {
			if err == utils.ErrTokenExpired {
				errFlag = true
				return
			}
			errorStruct.Error.Code = http.StatusUnauthorized
			errorStruct.Error.Message = "There was a problem with token validation."
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, &errorStruct)
			return
		}

		var user domain.User
		result := db.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			if errFlag {
				newToken, err := utils.GenerateToken(cfg.Token.ExpiredIn, user.ID, cfg.Token.JwtTokenSecret)
				if err != nil {
					errorStruct.Error.Code = http.StatusInternalServerError
					errorStruct.Error.Message = "There was a problem with token refresh."
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, &errorStruct)
					return
				}
				ctx.Header("Authorization", "Bearer "+newToken)
			}
			println(result.Error.Error())
			errorStruct.Error.Code = http.StatusBadGateway
			errorStruct.Error.Message = "The user could not be found by token."
			ctx.AbortWithStatusJSON(http.StatusForbidden, &errorStruct)
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
