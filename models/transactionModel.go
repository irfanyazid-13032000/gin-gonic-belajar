package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	GiverID    uint64 `gorm:"column:giver"`  
	ReceiverID uint64 `gorm:"column:receiver"`
	Giver      User   `gorm:"foreignKey:GiverID"`
	Receiver   User   `gorm:"foreignKey:ReceiverID"`
	Amount     float64 `json:"amount" binding:"required"`
}