package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/schemas"
	"net/http"
)

// @BasePath /api/v1
// @Summary Delete Opening
// @Description Delete a job opening
// @Accept json
// @Produce json
// @Param id query int true "Opening ID"
// @Success 200 {object} handler.DeleteOpeningResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	//find
	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}

	//delete
	if err := db.Delete(&opening).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("cannot delete opening with %s id", id))
		return
	}

	SendSuccess(ctx, "delete opening", opening)
}
