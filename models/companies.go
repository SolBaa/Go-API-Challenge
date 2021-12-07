package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Code     string `gorm:"primarykey"`
	PublicID string `json:"id"`
	Name     string
	Users    []User `gorm:"many2many:user_companies"`
}
