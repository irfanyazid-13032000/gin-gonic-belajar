package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `gorm:"column:judul"`
	Body  string `gorm:"column:body"`
}
