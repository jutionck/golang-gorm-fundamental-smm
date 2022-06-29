package model

type Customer struct {
	Id               string `gorm:"primaryKey"`
	Name             string `gorm:"size:50;not null"`
	Address          string
	Phone            string
	Email            string `gorm:"unique"`
	Balance          int
	IsStatus         int `gorm:"default:1"`
	UserCredentialID uint
	UserCredential   UserCredential
	BaseModel        BaseModel `gorm:"embedded"`
}

func (Customer) TableName() string {
	// ini akan membuat sebuah nama tabel (customisasi nama tabel)
	return "mst_customer"
}
