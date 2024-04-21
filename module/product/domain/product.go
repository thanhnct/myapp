package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	id          uuid.UUID
	status      string
	createdAt   time.Time
	updatedAt   time.Time
	categoryId  int
	name        string
	kind        string
	description string
}

func NewProduct(id uuid.UUID, name, kind, description, status string, categoryId int, createdAt, updatedAt time.Time) (*Product, error) {
	return &Product{
		id:          id,
		name:        name,
		kind:        kind,
		description: description,
		status:      status,
		categoryId:  categoryId,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}

func (u Product) Name() string {
	return u.name
}

func (u Product) Kind() string {
	return u.kind
}

func (u Product) Description() string {
	return u.description
}

func (u Product) CategoryId() int {
	return u.categoryId
}

func (u Product) Id() uuid.UUID {
	return u.id
}

func (u Product) Status() string {
	return u.status
}

func (u Product) CreatedAt() time.Time {
	return u.createdAt
}

func (u Product) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *Product) SetName(name string) {
	u.name = name
}

func (u *Product) SetStatus(status string) {
	u.status = status
}

func (u *Product) SetKind(kind string) {
	u.kind = kind
}

func (u *Product) SetDescription(description string) {
	u.description = description
}

func (u *Product) SetCategoryId(categoryId int) {
	u.categoryId = categoryId
}
