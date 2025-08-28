package model

import "gorm.io/gorm"

type Register struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
	Phone   string `json:"phone" gorm:"not null"`
	Age     int    `json:"age" gorm:"not null"`
}
