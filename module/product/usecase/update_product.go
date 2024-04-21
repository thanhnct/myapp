package usecase

import (
	"context"

	"github.com/viettranx/service-context/core"
)

type updateProductUC struct {
	productQueryRepo ProductQueryRepository
	productCmdRepo   ProductCommandRepository
}

func NewUpdateProductUC(productQueryRepo ProductQueryRepository, productCmdRepo ProductCommandRepository) *updateProductUC {
	return &updateProductUC{
		productQueryRepo: productQueryRepo,
		productCmdRepo:   productCmdRepo,
	}
}

func (uc *updateProductUC) UpdateProduct(ctx context.Context, dto ProductUpdateDTO) error {

	product, err := uc.productQueryRepo.Find(ctx, dto.Id)
	if err != nil {
		return core.ErrInternalServerError.WithError("can not get product").WithDebug(err.Error())
	}

	product.SetName(*dto.Name)
	product.SetKind(*dto.Kind)
	product.SetCategoryId(*dto.CategoryId)
	product.SetDescription(*dto.Description)
	product.SetStatus(*dto.Status)

	if err := uc.productCmdRepo.UpdateProduct(ctx, product); err != nil {
		return core.ErrInternalServerError.WithError("can not update product").WithDebug(err.Error())
	}

	return nil
}
