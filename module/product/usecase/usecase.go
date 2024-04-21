package usecase

import (
	"context"
	"myapp/module/product/domain"
)

type UseCase interface {
	CreateProduct(ctx context.Context, dto ProductCreateDTO) error
}

type useCase struct {
	*createProductUC
}

type Builder interface {
	BuildProductQueryRepo() ProductQueryRepository
	BuildProductCmdRepo() ProductCommandRepository
}

func UseCaseWithBuilder(b Builder) UseCase {
	return &useCase{
		createProductUC: NewProductUC(b.BuildProductCmdRepo()),
	}
}

type ProductRepository interface {
	ProductQueryRepository
	ProductCommandRepository
}

type ProductQueryRepository interface {
	// Find(ctx context.Context, id uuid.UUID) (*userDomain.User, error)
	// FindByEmail(ctx context.Context, email string) (*userDomain.User, error)
}

type ProductCommandRepository interface {
	CreateProduct(ctx context.Context, data *domain.Product) error
	//Update(ctx context.Context, data *domain.Product) error
}
