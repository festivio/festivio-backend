package handler

import (
	"net/http"

	"github.com/festivio/festivio-backend/domain"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMe() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var errorStruct domain.ErrorStruct
		currentUser, ok := ctx.MustGet("currentUser").(domain.User)
		if !ok {
			errorStruct.Error.Code = http.StatusUnauthorized
			errorStruct.Error.Message = "Failed to get information."
			ctx.JSON(http.StatusUnauthorized, &errorStruct)
			return
		}

		notifyParams := struct {
			TgUserID int `json:"tg_user_id"`
		}{currentUser.TgUserId}
		params := struct {
			Notifications       bool `json:"notifications"`
			NotificationsParams struct {
				TgUserID int `json:"tg_user_id"`
			} `json:"notifications_params"`
		}{currentUser.Notifications, notifyParams}

		userInfo := &domain.UserInfo{
			Email:  currentUser.Email,
			Name:   currentUser.Name,
			Phone:  currentUser.Phone,
			Role:   currentUser.Role,
			Params: params,
		}

		ctx.JSON(http.StatusOK, userInfo)
	}
}
