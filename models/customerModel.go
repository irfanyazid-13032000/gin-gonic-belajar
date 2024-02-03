package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	UserID  uint     `gorm:"foreignKey:UserID"`
	User    *User    `json:"user"`
	Balance float64  `json:"balance" binding:"required"`
}
