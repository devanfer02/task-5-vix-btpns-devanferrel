package models 

import (
	"time"
)

type User struct {
	ID 			int64 		`gorm:"primaryKey;autoIncrement;not null;unique" json:"id"`
	Username 	string  	`gorm:"type:varchar(255);not null" json:"username"`
	Email 		string 		`gorm:"type:varchar(255);not null; unique" json:"email"`
	Password 	string 		`gorm:"type:varchar(255);not null; size:255;min:6" json:"password"`
	Photos 		[]Photo 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos,omitempty"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
}