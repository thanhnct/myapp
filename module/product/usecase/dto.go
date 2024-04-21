package usecase

import "github.com/google/uuid"

type ProductUpdateDTO struct {
	Name        *string `gorm:"column:name" json:"name"`
	CategoryId  *int    `gorm:"column:category_id" json:"category_id"`
	Status      *string `gorm:"column:status" json:"status"`
	Kind        *string `gorm:"column:kind" json:"kind"`
	Description *string `gorm:"column:description" json:"description"`
}

type ProductCreateDTO struct {
	Id          uuid.UUID `column:"id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	CategoryId  int       `gorm:"column:category_id" json:"category_id"`
	Kind        string    `gorm:"column:kind" json:"kind"`
	Description string    `gorm:"column:description" json:"description"`
}
