package usecase

import (
	"context"
	"myapp/common"

	"github.com/viettranx/service-context/core"
)

type listProductUC struct {
	productQueryRepo ProductQueryRepository
}

func NewListProductUC(productQueryRepo ProductQueryRepository) *listProductUC {
	return &listProductUC{
		productQueryRepo: productQueryRepo,
	}
}

type ListProductFilter struct {
	CategoryId string `form:"category_id" json:"category_id"`
}

type ListProductParam struct {
	common.Paging
	ListProductFilter
}

func (uc *listProductUC) ListProduct(ctx context.Context, param *ListProductParam) (*[]ProductResponseDTO, error) {
	list, err := uc.productQueryRepo.ListProduct(ctx, param)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError("can not list product").WithDebug(err.Error())
	}

	listProductResponse := make([]ProductResponseDTO, len(list))

	for i, v := range list {
		prod := ProductResponseDTO{
			Id:          v.Id(),
			Name:        v.Name(),
			Kind:        v.Kind(),
			Description: v.Description(),
			CatId:       v.CategoryId(),
			CreatedAt:   v.CreatedAt(),
			UpdatedAt:   v.UpdatedAt(),
		}
		listProductResponse[i] = prod
	}

	return &listProductResponse, nil
}
