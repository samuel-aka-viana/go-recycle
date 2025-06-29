package filters

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type OpeningFilter struct {
	Role             *string `filter:"field:role;op:exact;param:role"`
	RoleContains     *string `filter:"field:role;op:contains;param:role_contains"`
	Company          *string `filter:"field:company;op:exact;param:company"`
	CompanyContains  *string `filter:"field:company;op:contains;param:company_contains"`
	Location         *string `filter:"field:location;op:exact;param:location"`
	LocationContains *string `filter:"field:location;op:contains;param:location_contains"`

	Remote *bool `filter:"field:remote;op:exact;param:remote"`

	SalaryMin   *int64 `filter:"field:salary;op:gte;param:salary_min"`
	SalaryMax   *int64 `filter:"field:salary;op:lte;param:salary_max"`
	SalaryExact *int64 `filter:"field:salary;op:exact;param:salary"`

	CreatedAfter  *time.Time `filter:"field:created_at;op:gte;param:created_after"`
	CreatedBefore *time.Time `filter:"field:created_at;op:lte;param:created_before"`

	RolesIn     []string `filter:"field:role;op:in;param:roles"`
	CompaniesIn []string `filter:"field:company;op:in;param:companies"`
	LocationsIn []string `filter:"field:location;op:in;param:locations"`

	RolesNotIn     []string `filter:"field:role;op:not_in;param:exclude_roles"`
	CompaniesNotIn []string `filter:"field:company;op:not_in;param:exclude_companies"`

	processor *AutoFilterProcessor
}

func NewOpeningFilter() *OpeningFilter {
	return &OpeningFilter{
		processor: &AutoFilterProcessor{},
	}
}

func (f *OpeningFilter) ApplyToQuery(ctx *gin.Context, query *gorm.DB) *gorm.DB {
	return f.processor.ApplyFilters(ctx, query, f)
}
