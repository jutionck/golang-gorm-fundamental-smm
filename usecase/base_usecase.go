package usecase

type BaseRepositoryAggregation interface {
	CalculateProductEachCustomer(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error
}
