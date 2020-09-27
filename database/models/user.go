package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"first_name"`
	LastName string `gorm:"not null" json:"last_name"`
	Email string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Status bool `gorm:"default:false" json:"status"`
	Phone string `gorm:"unique" json:"phone"`
}

func (u *User) TableName() string {
	return "users"
}