package handler

import (
	"github.com/gin-gonic/gin"
	"go-project/schemas"
	"go-project/utils"
	"net/http"
)

// @BasePath /api/v1
// @Summary Update Opening
// @Description Update fields of a job opening
// @Accept json
// @Produce json
// @Param id query int true "Opening ID"
// @Param request body handler.UpdateOpeningRequest true "Request Body"
// @Success 200 {object} handler.UpdateOpeningResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Router /opening [patch]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	_ = ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamsIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//update
	changes := utils.StructToMap(&request)
	if len(changes) == 0 {
		SendError(ctx, http.StatusBadRequest, "no fields to update")
		return
	}

	if err := db.Model(&opening).Updates(changes).Error; err != nil {
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	SendSuccess(ctx, "Opening Update", opening)
}
