package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string         `gorm:"not null"`
	Author   pq.StringArray `gorm:"type:text[]"`
	ISBN     string
	Weight   int
	Language string `gorm:"not null"`
	Pages    int
	Price    float32        `gorm:"not null"`
	Genre    pq.StringArray `gorm:"type:text[]"`
	Image    pq.StringArray `gorm:"type:text[]"`

	Available bool `gorm:"default:true"`
	Slug      string
	UserID    uint
}
