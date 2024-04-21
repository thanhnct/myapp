package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/viettranx/service-context/core"
)

type deleteProductUC struct {
	productQueryRepo ProductQueryRepository
	productCmdRepo   ProductCommandRepository
}

func NewDeleteProductUC(productQueryRepo ProductQueryRepository, productCmdRepo ProductCommandRepository) *deleteProductUC {
	return &deleteProductUC{
		productQueryRepo: productQueryRepo,
		productCmdRepo:   productCmdRepo,
	}
}

func (uc *updateProductUC) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	product, err := uc.productQueryRepo.Find(ctx, id)
	if err != nil {
		return core.ErrInternalServerError.WithError("can not get product").WithDebug(err.Error())
	}

	if err := uc.productCmdRepo.DeleteProduct(ctx, product.Id()); err != nil {
		return core.ErrInternalServerError.WithError("can not update product").WithDebug(err.Error())
	}

	return nil
}
