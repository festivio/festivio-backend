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
			if strings.Contains(err.Error(), "Token is expired") {
				errorStruct.Error.Code = http.StatusForbidden
				errorStruct.Error.Message = "Please log-in to your account."
				ctx.AbortWithStatusJSON(http.StatusForbidden, &errorStruct)
				return
			} else {
				errorStruct.Error.Code = http.StatusUnauthorized
				errorStruct.Error.Message = err.Error()
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, &errorStruct)
				return
			}
		}

		var user domain.User
		result := db.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			errorStruct.Error.Code = http.StatusBadGateway
			errorStruct.Error.Message = result.Error.Error()
			ctx.AbortWithStatusJSON(http.StatusBadGateway, &errorStruct)
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
