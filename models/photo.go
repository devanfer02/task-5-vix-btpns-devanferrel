package models

type Photo struct {
	ID 			int64 		`gorm:"primaryKey;autoIncrement;not null;unique" json:"id"`
	Title 		*string  	`gorm:"not null" json:"title" binding:"required"`
	Caption 	*string 	`gorm:"not null" json:"caption" binding:"required"`
	PhotoUrl 	*string 	`gorm:"not null" json:"photo_url" binding:"required"`
	UserID 		int64 		`gorm:"not null" json:"user_id"`
}