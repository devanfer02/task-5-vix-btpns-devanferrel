package models

type Photo struct {
	ID 			int64 		`gorm:"primaryKey;autoIncrement;not null;unique" json:"id"`
	Title 		string  	`gorm:"not null" json:"title"`
	Caption 	string 		`gorm:"not null" json:"caption"`
	PhotoUrl 	string 		`gorm:"not null" json:"photoUrl"`
	UserID 		int64 		`gorm:"not null" json:"userId"`
	User 		User		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}