package repository

import (
	"context"
	"errors"
	"myapp/common"
	"myapp/module/product/domain"
	"myapp/module/product/usecase"

	"github.com/google/uuid"
	"github.com/viettranx/service-context/core"
	"gorm.io/gorm"
)

const TbProduct = "products"

type productMySQLRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) productMySQLRepo {
	return productMySQLRepo{db: db}
}

func (repo productMySQLRepo) Find(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	var dto ProductDTO

	if err := repo.db.Table(TbProduct).Where("id = ?", id).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return dto.ToEntity()
}

func (repo productMySQLRepo) CreateProduct(ctx context.Context, data *domain.Product) error {

	dto := ProductDTO{
		Id:          data.Id(),
		Name:        data.Name(),
		Kind:        data.Kind(),
		Description: data.Description(),
		CategoryId:  data.CategoryId(),
		CreatedAt:   data.CreatedAt(),
		UpdatedAt:   data.UpdatedAt(),
		Status:      data.Status(),
	}

	if err := repo.db.Table(TbProduct).Create(&dto).Error; err != nil {
		return err
	}
	return nil
}

func (repo productMySQLRepo) UpdateProduct(ctx context.Context, data *domain.Product) error {

	dto := ProductDTO{
		Id:          data.Id(),
		Name:        data.Name(),
		Kind:        data.Kind(),
		Description: data.Description(),
		CategoryId:  data.CategoryId(),
		CreatedAt:   data.CreatedAt(),
		UpdatedAt:   data.UpdatedAt(),
		Status:      data.Status(),
	}

	if err := repo.db.Table(TbProduct).Updates(&dto).Error; err != nil {
		return err
	}
	return nil
}

func (repo productMySQLRepo) DeleteProduct(ctx context.Context, id uuid.UUID) error {

	if err := repo.db.Table(TbProduct).Where("id = ?", id).Delete(&ProductDTO{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo productMySQLRepo) ListProduct(ctx context.Context, param *usecase.ListProductParam) ([]*domain.Product, error) {
	var products []ProductDTO

	db := repo.db.Table(TbProduct)

	if param.CategoryId != "" {
		db = db.Where("category_id = ?", param.CategoryId)
	}

	var count int64
	db.Count(&count)
	param.Total = int(count)

	param.Process()

	offset := param.Limit * (param.Page - 1)

	//db = db.Preload("Category")

	if err := db.Offset(offset).Limit(param.Limit).Order("id desc").Find(&products).Error; err != nil {
		return nil, core.ErrBadRequest.WithError("cannot list product").WithDebug(err.Error())
	}

	response := make([]*domain.Product, len(products))

	for i, v := range products {
		p, _ := v.ToEntity()
		response[i] = p
	}

	return response, nil
}
