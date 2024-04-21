package usecase

import (
	"time"

	"github.com/google/uuid"
)

type ProductUpdateDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        *string   `json:"name"`
	CategoryId  *int      `json:"category_id"`
	Status      *string   `json:"status"`
	Kind        *string   `json:"kind"`
	Description *string   `json:"description"`
}

type ProductCreateDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CategoryId  int       `json:"category_id"`
	Kind        string    `json:"kind"`
	Description string    `json:"description"`
}

type ProductResponseDTO struct {
	Id          uuid.UUID `json:"id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CatId       int       `json:"category_id"`
	Name        string    `json:"name"`
	Kind        string    `json:"kind"`
	Description string    `json:"description"`
}
