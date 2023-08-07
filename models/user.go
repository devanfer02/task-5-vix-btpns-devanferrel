package models 

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserAuth
	Email 		string 		`gorm:"type:varchar(255);not null; unique" json:"email" valid:"email, required"`
	Photos 		[]Photo 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos,omitempty"`
}

type UserAuth struct {
	Username 	string  `gorm:"type:varchar(255);not null" json:"username" valid:"alphanum, required"`
	Password 	string 	`gorm:"type:varchar(255);not null; size:255;" json:"password" valid:"required, stringlength(6|255)"`
}