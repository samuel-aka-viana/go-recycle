package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-project/filters"
	"go-project/schemas"
)

// @BasePath /api/v1
// @Summary List Openings with Filters
// @Description Get job openings with automatic filtering (Django-style)
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Items per page"
// @Param role query string false "Exact role match"
// @Param role_contains query string false "Role contains text"
// @Param company query string false "Exact company"
// @Param company_contains query string false "Company contains"
// @Param location query string false "Exact location"
// @Param location_contains query string false "Location contains"
// @Param remote query bool false "Remote work"
// @Param salary_min query int false "Minimum salary"
// @Param salary_max query int false "Maximum salary"
// @Param created_after query string false "Created after (YYYY-MM-DD)"
// @Param created_before query string false "Created before (YYYY-MM-DD)"
// @Param roles query []string false "Roles list"
// @Param companies query []string false "Companies list"
// @Param exclude_roles query []string false "Exclude roles"
// @Success 200 {object} handler.ListOpeningsResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	page := 1
	pageSize := 10

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

	openingFilter := filters.NewOpeningFilter()

	query := db.Model(&schemas.Opening{})
	query = openingFilter.ApplyToQuery(ctx, query)

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("cannot count openings: %s", err))
		return
	}

	var openings []schemas.Opening
	offset := (page - 1) * pageSize

	if err := query.Limit(pageSize).Offset(offset).Find(&openings).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("cannot get openings: %s", err))
		return
	}

	totalPages := int((totalCount + int64(pageSize) - 1) / int64(pageSize))

	response := gin.H{
		"items": openings,
		"meta": gin.H{
			"page":       page,
			"pageSize":   pageSize,
			"totalItems": totalCount,
			"totalPages": totalPages,
		},
	}

	SendSuccess(ctx, "filtered openings", response)
}
