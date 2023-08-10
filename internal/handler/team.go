package handler

import (
	"net/http"

	"github.com/festivio/festivio-backend/domain"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTeam() gin.HandlerFunc {
	return	func(ctx *gin.Context) {
		var errorStruct domain.ErrorStruct
		users, err := h.srv.GetTeam()
		if err != nil {
			errorStruct.Error.Code = http.StatusBadGateway
			errorStruct.Error.Message = "Failed to get team."
			ctx.JSON(http.StatusUnauthorized, &errorStruct)
			return
		}
		
		ctx.JSON(http.StatusOK, &users)
	}
}