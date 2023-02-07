package structs

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name     string
	Email    string
	Products []Product
}

// Product struct
type Product struct {
	gorm.Model
	Name   string
	Price  float64
	UserID uint
}
