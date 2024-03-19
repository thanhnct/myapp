package productdomain

import "github.com/google/uuid"

type ProductUpdateDTO struct {
	Name        *string `gorm:"column:name" json:"name"`
	CategoryId  *int    `gorm:"column:category_id" json:"category_id"`
	Status      *string `gorm:"column:status" json:"status"`
	Type        *string `gorm:"column:type" json:"type"`
	Description *string `gorm:"column:description" json:"description"`
}

type ProductCreateDTO struct {
	Id          uuid.UUID `column:"id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	CategoryId  int       `gorm:"column:category_id" json:"category_id"`
	Type        string    `gorm:"column:type" json:"type"`
	Description string    `gorm:"column:description" json:"description"`
}

func (ProductUpdateDTO) TableName() string {
	return "products"
}

func (ProductCreateDTO) TableName() string {
	return ProductUpdateDTO{}.TableName()
}
