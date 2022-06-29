package model

import "gorm.io/gorm"

type UserCredential struct {
	UserName string `gorm:"unique;size:50;not null"`
	Password string `gorm:"size:10;not null"`
	IsActive bool   `gorm:"default:false"`
	gorm.Model
}
