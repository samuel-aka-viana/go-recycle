package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/schemas"
	"net/http"
)

// @BasePath /api/v1
// @Summary Show Opening
// @Description Get one job opening by ID
// @Accept json
// @Produce json
// @Param id query int true "Opening ID"
// @Success 200 {object} handler.ShowOpeningResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}

	SendSuccess(ctx, "show opening", opening)
}
