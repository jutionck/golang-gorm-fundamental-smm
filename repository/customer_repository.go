package repository

import (
	"enigmacamp.com/golang-gorm/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	Update(customer *model.Customer, by map[string]interface{}) error
	UpdateBy(existingCustomer *model.Customer) error
	Delete(customer *model.Customer) error
	FindById(id string) (model.Customer, error)
	FindFirstBy(by map[string]interface{}) (model.Customer, error)   // where column = ? limit 1
	FindAllBy(by map[string]interface{}) ([]model.Customer, error)   // where column = ?
	FindBy(by string, vals ...interface{}) ([]model.Customer, error) // where column like ?
	FindFirstWithPreload(by map[string]interface{}, preload string) (interface{}, error)
	BaseRepositoryAggregation
	BaseRepositoryPaging
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) FindFirstWithPreload(by map[string]interface{}, preload string) (interface{}, error) {
	var customer model.Customer
	result := c.db.Preload(preload).Where(by).First(&customer)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) UpdateBy(existingCustomer *model.Customer) error {
	result := c.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(existingCustomer).Error
	return result
}

func (c *customerRepository) Count(result interface{}, groupBy string) error {
	sqlStmt := c.db.Model(&model.Customer{}).Unscoped()
	var res *gorm.DB
	if groupBy == "" {
		t, ok := result.(*int64) // casting: interface{} to *int64
		if ok {
			res = sqlStmt.Count(t)
		} else {
			return errors.New("must be int64") // custom error
		}
	} else {
		res = sqlStmt.Select(fmt.Sprintf("%s,%s", groupBy, "count(*) as total")).Group(groupBy).Find(result)
	}
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) GroupBy(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error {
	res := c.db.Model(&model.Customer{}).Unscoped().Select(selectedBy).Where(whereBy).Group(groupBy).Find(result)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (c *customerRepository) Paging(page int, itemPerPage int) (interface{}, error) {
	var customers []model.Customer
	offset := itemPerPage * (page - 1)
	res := c.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&customers)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil
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
