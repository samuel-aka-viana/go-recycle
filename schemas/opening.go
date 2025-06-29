package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Opening struct {
	gorm.Model
	Role     string `gorm:"not null"`
	Company  string `gorm:"not null"`
	Location string
	Remote   bool
	Link     string
	Salary   float64 `gorm:"not null"`
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   bool    `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}
