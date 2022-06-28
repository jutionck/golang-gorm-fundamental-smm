package repository

import (
	"enigmacamp.com/golang-gorm/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Create(customer *model.Customer) error {
	// sudah otomatis deteksi gorm nya kalo fungsi di bawah adalah insert
	// SQL Builder
	result := c.db.Create(customer).Error
	return result
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
