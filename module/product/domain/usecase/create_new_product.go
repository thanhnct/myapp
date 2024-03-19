package productusecase

import (
	"context"
	"myapp/common"
	productdomain "myapp/module/product/domain"
	"strings"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreateDTO) error
}

func NewCreateProductUseCase(repo CreateProductRepository) CreateNewProductUseCase {
	return CreateNewProductUseCase{
		repo: repo,
	}
}

type CreateNewProductUseCase struct {
	repo CreateProductRepository
}

func (uc CreateNewProductUseCase) CreateProduct(ctx context.Context, prod *productdomain.ProductCreateDTO) error {
	prod.Id = common.GenUUID()

	prod.Name = strings.TrimSpace(prod.Name)

	if prod.Name == "" {
		return productdomain.ErrorProductNameCannotBeBlank
	}
	if err := uc.repo.CreateProduct(ctx, prod); err != nil {
		return err
	}
	return nil
}

type CreateProductRepository interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreateDTO) error
}
