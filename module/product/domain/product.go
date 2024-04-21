package domain

import (
	"myapp/common"
)

type Product struct {
	common.EntityBase
	categoryId  int
	name        string
	kind        string
	description string
}

func NewProduct(name string, kind string, description string, categoryId int) (*Product, error) {
	return &Product{
		EntityBase:  common.GenNewEntityBase(),
		name:        name,
		kind:        kind,
		description: description,
		categoryId:  categoryId,
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
