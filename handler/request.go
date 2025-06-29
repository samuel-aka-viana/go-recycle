package handler

import (
	"fmt"
)

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s )is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r == nil {
		return fmt.Errorf("mal formed boby")
	}

	if r.Role == "" {
		return errParamsIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamsIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamsIsRequired("location", "string")
	}
	if r.Remote == nil {
		return errParamsIsRequired("remote", "bool")
	}
	if r.Link == "" {
		return errParamsIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return errParamsIsRequired("salary", "float64")
	}
	return nil
}
