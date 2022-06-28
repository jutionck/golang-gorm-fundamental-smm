package repository

import (
	"enigmacamp.com/golang-gorm/model"
	"errors"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	Update(customer *model.Customer, by map[string]interface{}) error
	Delete(customer *model.Customer) error
	FindById(id string) (model.Customer, error)
	FindFirstBy(by map[string]interface{}) (model.Customer, error)   // where column = ? limit 1
	FindAllBy(by map[string]interface{}) ([]model.Customer, error)   // where column = ?
	FindBy(by string, vals ...interface{}) ([]model.Customer, error) // where column like ?
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) FindFirstBy(by map[string]interface{}) (model.Customer, error) {
	var customer model.Customer
	result := c.db.Unscoped().Where(by).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindAllBy(by map[string]interface{}) ([]model.Customer, error) {
	var customer []model.Customer
	result := c.db.Unscoped().Where(by).Find(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindBy(by string, vals ...interface{}) ([]model.Customer, error) {
	var customer []model.Customer
	result := c.db.Unscoped().Where(by, vals...).Find(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindById(id string) (model.Customer, error) {
	var customer model.Customer
	result := c.db.Unscoped().First(&customer, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) Delete(customer *model.Customer) error {
	result := c.db.Delete(customer).Error
	return result
}

func (c *customerRepository) Update(customer *model.Customer, by map[string]interface{}) error {
	result := c.db.Model(customer).Updates(by).Error
	return result
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
