package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/schemas"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {
	opening := []schemas.Opening{}
	if err := db.Find(&opening).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("cannot get opening with %s", err))
		return
	}
	SendSuccess(ctx, "list all openings", opening)
}
