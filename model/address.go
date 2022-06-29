package model

import "gorm.io/gorm"

type Address struct {
	StreetName string `gorm:"not null"`
	City       string `gorm:"not null"`
	PostalCode string `gorm:"not null"`
	CustomerID string `gorm:"not null"`
	gorm.Model
}

func (a Address) TableName() string {
	return "mst_customer_address"
}
