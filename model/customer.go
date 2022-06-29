package model

import "encoding/json"

type Customer struct {
	Id               string `gorm:"primaryKey"`
	Name             string `gorm:"size:50;not null"`
	Address          []Address
	Phone            string
	Email            string `gorm:"unique"`
	Balance          int
	IsStatus         int `gorm:"default:1"`
	UserCredentialID uint
	UserCredential   UserCredential
	Products         []Product `gorm:"many2many:customer_products;"`
	BaseModel        BaseModel `gorm:"embedded"`
}

func (Customer) TableName() string {
	// ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "mst_customer"
}

func (c *Customer) ToString() string { // cetak seperti Json
	customer, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}
	return string(customer)
}
