package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	UserID 		uint 		`gorm:"not null" json:"user_id"`
	Title 		string  	`gorm:"not null" json:"title" valid:"required"`
	Caption 	string 		`gorm:"not null" json:"caption" valid:"required"`
	PhotoUrl 	string 		`gorm:"not null" json:"photo_url" valid:"required"`
}