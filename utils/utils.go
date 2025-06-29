package utils

import (
	"reflect"
)

func StructToMap(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(input).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		jsonTag := fieldType.Tag.Get("json")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				result[jsonTag] = field.String()
			}
		case reflect.Float64, reflect.Float32:
			if field.Float() != 0 {
				result[jsonTag] = field.Float()
			}
		case reflect.Int, reflect.Int64, reflect.Int32:
			if field.Int() != 0 {
				result[jsonTag] = field.Int()
			}
		case reflect.Ptr:
			if !field.IsNil() {
				result[jsonTag] = field.Elem().Interface()
			}
		case reflect.Bool:
			result[jsonTag] = field.Bool()
		default:
			result[jsonTag] = field.Interface()
		}
	}

	return result
}
