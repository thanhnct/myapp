package usecase

import (
	"context"
	"fmt"
	"myapp/common"
	"myapp/module/image/domain"
	"time"

	"github.com/viettranx/service-context/core"
)

type UseCase interface {
	UploadImage(ctx context.Context, dto UploadDTO) (*domain.Image, error)
}

type useCase struct {
	uploader Uploader
	repo     CmdRepository
}

func NewUseCase(uploader Uploader, repo CmdRepository) useCase {
	return useCase{uploader: uploader, repo: repo}
}

func (uc useCase) UploadImage(ctx context.Context, dto UploadDTO) (*domain.Image, error) {
	dstFileName := fmt.Sprintf("%d_%s", time.Now().UTC().UnixNano(), dto.FileName)

	if err := uc.uploader.SaveFileUploaded(ctx, dto.FileData, dstFileName); err != nil {
		return nil, core.ErrInternalServerError.WithError(ErrCannotUploadImage.Error()).WithDebug(err.Error())
	}

	image := domain.Image{
		Id:              common.GenUUID(),
		Title:           dto.Name,
		FileName:        dstFileName,
		FileSize:        dto.FileSize,
		FileType:        dto.FileType,
		StorageProvider: uc.uploader.GetName(),
		Status:          domain.StatusUploaded,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}

	if err := uc.repo.Create(ctx, &image); err != nil {
		return nil, core.ErrInternalServerError.WithError(ErrCannotUploadImage.Error()).WithDebug(err.Error())
	}

	return &image, nil
}

type Uploader interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) error
	GetName() string
	GetDomain() string
}

type CmdRepository interface {
	Create(ctx context.Context, entity *domain.Image) error
}
