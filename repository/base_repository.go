package repository

type BaseRepositoryAggregation interface {
	Count(result interface{}, groupBy string) error
	GroupBy(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error
}

type BaseRepositoryPaging interface {
	Paging(page int, itemPerPage int) (interface{}, error)
}

type BaseRepositoryFindBy interface {
	FindById(id string) (interface{}, error)
	FindFirstBy(by map[string]interface{}) (interface{}, error)   // where column = ? limit 1
	FindAllBy(by map[string]interface{}) ([]interface{}, error)   // where column = ?
	FindBy(by string, vals ...interface{}) ([]interface{}, error) // where column like ?
	FindFirstWithPreload(by map[string]interface{}, preload string) (interface{}, error)
}
