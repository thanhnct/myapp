package repository

import (
	"context"
	"myapp/module/product/domain"

	"gorm.io/gorm"
)

const TbProduct = "products"

type productMySQLRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) productMySQLRepo {
	return productMySQLRepo{db: db}
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
