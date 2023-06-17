package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Auther   string
	category string `gorm:"not null;default:null"`
	Title    string `gorm:"unique;not null;type:varchar(100);default:null"`
	Body     string `gorm:"not null;default:null"`
	HashTag  string
}
