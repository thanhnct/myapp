package builder

import (
	userusecase "myapp/module/user/usecase"

	"gorm.io/gorm"
)

type simpleBuilder struct {
	db *gorm.DB
	tp userusecase.TokenProvider
}

func NewSimpleBuilder(db *gorm.DB, tp userusecase.TokenProvider) simpleBuilder {
	return simpleBuilder{db: db, tp: tp}
}
