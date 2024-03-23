package common

import "gorm.io/gorm"

var (
	Activated    = "activated"
	KeyRequester = "requester"
	KeyGorm      = "gorm"
	KeyJWT       = "jwt"
)

type DbContext interface {
	GetDB() *gorm.DB
}
