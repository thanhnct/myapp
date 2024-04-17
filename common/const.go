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

type Paging struct {
	Page  int `json:"page"`
	Total int `json:"total"`
	Limit int `json:"limit"`
}

func (p *Paging) Process() {
	if p.Limit < 1 {
		p.Limit = 10
	}

	if p.Limit > 100 {
		p.Limit = 100
	}

	if p.Page < 1 {
		p.Page = 1
	}
}
