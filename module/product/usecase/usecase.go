package usecase

import (
	"context"
	"myapp/module/product/domain"

	"github.com/google/uuid"
)

type UseCase interface {
	CreateProduct(ctx context.Context, dto ProductCreateDTO) error
	ListProduct(ctx context.Context, param *ListProductParam) (*[]ProductResponseDTO, error)
	UpdateProduct(ctx context.Context, dto ProductUpdateDTO) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

type useCase struct {
	*createProductUC
	*listProductUC
	*updateProductUC
}

type Builder interface {
	BuildProductQueryRepo() ProductQueryRepository
	BuildProductCmdRepo() ProductCommandRepository
}

func UseCaseWithBuilder(b Builder) UseCase {
	return &useCase{
		createProductUC: NewProductUC(b.BuildProductCmdRepo()),
		listProductUC:   NewListProductUC(b.BuildProductQueryRepo()),
		updateProductUC: NewUpdateProductUC(b.BuildProductQueryRepo(), b.BuildProductCmdRepo()),
	}
}

type ProductRepository interface {
	ProductQueryRepository
	ProductCommandRepository
}

type ProductQueryRepository interface {
	ListProduct(ctx context.Context, param *ListProductParam) ([]*domain.Product, error)
	Find(ctx context.Context, id uuid.UUID) (*domain.Product, error)
}

type ProductCommandRepository interface {
	CreateProduct(ctx context.Context, data *domain.Product) error
	UpdateProduct(ctx context.Context, data *domain.Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}
