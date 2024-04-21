package builder

import (
	"myapp/module/product/infras/repository"
	"myapp/module/product/usecase"
)

func (s simpleBuilder) BuildProductQueryRepo() usecase.ProductQueryRepository {
	return repository.NewProductRepo(s.db)
}

func (s simpleBuilder) BuildProductCmdRepo() usecase.ProductCommandRepository {
	return repository.NewProductRepo(s.db)
}
