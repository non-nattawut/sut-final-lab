package entity

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:""`
	Email      string
	EmployeeID string `valid:""`
}
