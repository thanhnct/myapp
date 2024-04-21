package usecase

import (
	"context"
	"myapp/module/product/domain"

	"github.com/viettranx/service-context/core"
)

type createProductUC struct {
	productCommandRepo ProductCommandRepository
}

func NewProductUC(productCommandRepo ProductCommandRepository) *createProductUC {
	return &createProductUC{
		productCommandRepo: productCommandRepo,
	}
}

func (uc *createProductUC) CreateProduct(ctx context.Context, dto ProductCreateDTO) error {
	product, err := domain.NewProduct(
		dto.Name,
		dto.Kind,
		dto.Description,
		0,
	)

	if err != nil {
		return core.ErrInternalServerError.WithError("unable to initialize product").WithDebug(err.Error())
	}

	if err := uc.productCommandRepo.CreateProduct(ctx, product); err != nil {
		return core.ErrInternalServerError.WithError("can not create product").WithDebug(err.Error())
	}

	return nil
}
