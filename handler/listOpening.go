package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project/schemas"
	"net/http"
	"strconv"
)

// @BasePath /api/v1
// @Summary List Openings
// @Description Get a list of all job openings with pagination
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} handler.ListOpeningsResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	// Defaults
	page := 1
	pageSize := 10

	// Read query params
	if p := ctx.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	if ps := ctx.Query("pageSize"); ps != "" {
		if v, err := strconv.Atoi(ps); err == nil && v > 0 {
			pageSize = v
		}
	}

	var openings []schemas.Opening
	offset := (page - 1) * pageSize

	// Query with LIMIT + OFFSET
	if err := db.Limit(pageSize).Offset(offset).Find(&openings).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("cannot get opening with %s", err))
		return
	}

	// Optionally, send pagination metadata
	response := gin.H{
		"items": openings,
		"meta": gin.H{
			"page":     page,
			"pageSize": pageSize,
		},
	}

	SendSuccess(ctx, "list all openings", response)
}
