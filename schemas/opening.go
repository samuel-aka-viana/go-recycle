package schemas

import (
	"gorm.io/gorm"
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
