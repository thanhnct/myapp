package usecase

import (
	"context"
	"myapp/common"
	"myapp/module/product/domain"
	"time"

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
	now := time.Now().UTC()
	product, err := domain.NewProduct(
		common.GenUUID(),
		dto.Name,
		dto.Kind,
		dto.Description,
		"activated",
		0,
		now,
		now,
	)

	if err != nil {
		return core.ErrInternalServerError.WithError("unable to initialize product").WithDebug(err.Error())
	}

	if err := uc.productCommandRepo.CreateProduct(ctx, product); err != nil {
		return core.ErrInternalServerError.WithError("can not create product").WithDebug(err.Error())
	}

	return nil
}
