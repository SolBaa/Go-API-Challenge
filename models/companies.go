package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	// UserID   uint
	PublicID string
	Name     string
}
