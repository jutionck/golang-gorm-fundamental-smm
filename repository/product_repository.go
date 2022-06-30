package repository

import (
	"enigmacamp.com/golang-gorm/model"
	"errors"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindById(id uint) (model.Product, error)
	FindAll() ([]model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	result := p.db.Find(&products)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return products, nil
}

func (p *productRepository) FindById(id uint) (model.Product, error) {
	var product model.Product
	result := p.db.Unscoped().First(&product, "id = ?", id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		} else {
			return product, err
		}
	}
	return product, nil
}

func (p *productRepository) Create(product *model.Product) error {
	result := p.db.Create(product).Error
	return result
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
