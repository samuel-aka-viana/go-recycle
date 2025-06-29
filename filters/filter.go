package filters

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FilterConfig struct {
	Field      string
	Operation  string
	QueryParam string
}

type BaseFilter interface {
	ApplyToQuery(ctx *gin.Context, query *gorm.DB) *gorm.DB
}

type AutoFilterProcessor struct{}

func (p *AutoFilterProcessor) ApplyFilters(ctx *gin.Context, query *gorm.DB, filter BaseFilter) *gorm.DB {
	filterValue := reflect.ValueOf(filter).Elem()
	filterType := reflect.TypeOf(filter).Elem()

	for i := 0; i < filterType.NumField(); i++ {
		field := filterType.Field(i)
		fieldValue := filterValue.Field(i)

		filterTag := field.Tag.Get("filter")
		if filterTag == "" {
			continue
		}

		config := p.parseFilterTag(filterTag)
		if config == nil {
			continue
		}

		queryValue := p.getQueryValue(ctx, config, fieldValue.Type())
		if queryValue == nil {
			continue
		}

		query = p.applyFilterToQuery(query, config, queryValue)
	}

	return query
}

func (p *AutoFilterProcessor) parseFilterTag(tag string) *FilterConfig {
	config := &FilterConfig{}

	parts := strings.Split(tag, ";")
	for _, part := range parts {
		kv := strings.Split(part, ":")
		if len(kv) != 2 {
			continue
		}

		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])

		switch key {
		case "field":
			config.Field = value
		case "op":
			config.Operation = value
		case "param":
			config.QueryParam = value
		}
	}

	if config.Field == "" || config.Operation == "" || config.QueryParam == "" {
		return nil
	}

	return config
}

func (p *AutoFilterProcessor) getQueryValue(ctx *gin.Context, config *FilterConfig, fieldType reflect.Type) interface{} {
	switch fieldType.Kind() {
	case reflect.Ptr:
		elemType := fieldType.Elem()
		queryVal := ctx.Query(config.QueryParam)
		if queryVal == "" {
			return nil
		}

		switch elemType.Kind() {
		case reflect.String:
			return &queryVal
		case reflect.Bool:
			if val, err := strconv.ParseBool(queryVal); err == nil {
				return &val
			}
		case reflect.Int64:
			if val, err := strconv.ParseInt(queryVal, 10, 64); err == nil {
				return &val
			}
		case reflect.Int:
			if val, err := strconv.Atoi(queryVal); err == nil {
				return &val
			}
		default:
			if elemType == reflect.TypeOf(time.Time{}) {
				if val, err := time.Parse("2006-01-02", queryVal); err == nil {
					return &val
				}
			}
		}

	case reflect.Slice:
		queryVals := ctx.QueryArray(config.QueryParam)
		if len(queryVals) == 0 {
			return nil
		}
		return queryVals
	}

	return nil
}

func (p *AutoFilterProcessor) applyFilterToQuery(query *gorm.DB, config *FilterConfig, value interface{}) *gorm.DB {
	switch config.Operation {
	case "exact":
		return query.Where(fmt.Sprintf("%s = ?", config.Field), value)
	case "contains":
		if str, ok := value.(*string); ok {
			return query.Where(fmt.Sprintf("LOWER(%s) LIKE ?", config.Field), "%"+strings.ToLower(*str)+"%")
		}
	case "gte":
		return query.Where(fmt.Sprintf("%s >= ?", config.Field), value)
	case "lte":
		return query.Where(fmt.Sprintf("%s <= ?", config.Field), value)
	case "in":
		if arr, ok := value.([]string); ok && len(arr) > 0 {
			return query.Where(fmt.Sprintf("%s IN ?", config.Field), arr)
		}
	case "not_in":
		if arr, ok := value.([]string); ok && len(arr) > 0 {
			return query.Where(fmt.Sprintf("%s NOT IN ?", config.Field), arr)
		}
	case "startswith":
		if str, ok := value.(*string); ok {
			return query.Where(fmt.Sprintf("LOWER(%s) LIKE ?", config.Field), strings.ToLower(*str)+"%")
		}
	case "endswith":
		if str, ok := value.(*string); ok {
			return query.Where(fmt.Sprintf("LOWER(%s) LIKE ?", config.Field), "%"+strings.ToLower(*str))
		}
	case "isnull":
		if isNull, ok := value.(*bool); ok {
			if *isNull {
				return query.Where(fmt.Sprintf("%s IS NULL", config.Field))
			} else {
				return query.Where(fmt.Sprintf("%s IS NOT NULL", config.Field))
			}
		}
	}

	return query
}
