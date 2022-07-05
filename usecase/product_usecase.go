package usecase

import "enigmacamp.com/golang-gorm/repository"

type ProductUseCase interface {
	CalculateProductEachCustomer(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error
}

type productUseCase struct {
	customerRepo repository.CustomerRepository
	productRepo  repository.ProductRepository
}

func (p *productUseCase) CalculateProductEachCustomer(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error {
	return p.customerRepo.
}

func NewProductUseCase(
	customerRepo repository.CustomerRepository,
	productRepo repository.ProductRepository) ProductUseCase {

	uc := new(productUseCase)
	uc.customerRepo = customerRepo
	uc.productRepo = productRepo
	return uc
}
