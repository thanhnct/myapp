package usecase

import "errors"

type UploadDTO struct {
	Name     string
	FileName string
	FileType string
	FileSize int
	FileData []byte
}

var (
	ErrCannotUploadImage = errors.New("cannot upload image")
	ErrCannotFindImage   = errors.New("cannot find image")
)
