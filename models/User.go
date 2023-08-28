package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Mobile   int32 `gorm:"unique"`
	FName    string
	LName    string
	Books    []Book
	Image    string
}
