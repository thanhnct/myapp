package common

import "gorm.io/gorm"

var (
	Activated    = "activated"
	KeyRequester = "requester"
	KeyGorm      = "gorm"
	KeyJWT       = "jwt"
	KeyAWSS3     = "aws_s3"
)

type DbContext interface {
	GetDB() *gorm.DB
}
