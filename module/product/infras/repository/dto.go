package repository

import (
	"myapp/module/product/domain"
	"time"

	"github.com/google/uuid"
)

type ProductDTO struct {
	Id          uuid.UUID `gorm:"column:id;"`
	Name        string    `gorm:"column:name;"`
	Kind        string    `gorm:"column:kind;"`
	Description string    `gorm:"column:description;"`
	CategoryId  int       `gorm:"column:category_id;"`
	Status      string    `gorm:"column:status;"`
	CreatedAt   time.Time `gorm:"column:created_at;"`
	UpdatedAt   time.Time `gorm:"column:updated_at;"`
}

func (dto *ProductDTO) ToEntity() (*domain.Product, error) {

	return domain.NewProduct(dto.Id, dto.Name, dto.Kind, dto.Description, dto.Status, dto.CategoryId, dto.CreatedAt, dto.UpdatedAt)
}
