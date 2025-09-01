package model

import "gorm.io/gorm"

type InterviewDate struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
	Date    string `json:"date" gorm:"not null"`
	Images  string `json:"image" gorm:"column:image_url"`
}
