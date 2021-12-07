package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PublicID string
	Name     string
	LastName string
	Email    string
	Company  []Company `gorm:"many2many:user_companies;"`
}
