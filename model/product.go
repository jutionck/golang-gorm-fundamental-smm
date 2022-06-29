package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint       `gorm:"column:id;primaryKey"`
	ProductName string     `gorm:"column:name;not null"`
	Customers   []Customer `gorm:"many2many:customer_products"`
	gorm.Model
}

func (p Product) TableName() string {
	return "mst_product"
}

func (p *Product) ToString() string {
	product, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return ""
	}
	return string(product)
}
