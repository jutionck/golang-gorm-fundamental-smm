package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string      `gorm:"column:name;not null"`
	Customers   []*Customer `gorm:"many2many:customer_products"`
}

func (p Product) TableName() string {
	return "mst_product"
}
