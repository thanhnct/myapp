package repository

import (
	"context"
	"myapp/common"
	userdomain "myapp/module/user/domain"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const TbName = "users"

type userMySQLRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userMySQLRepo {
	return userMySQLRepo{db: db}
}

func (repo userMySQLRepo) FindByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	var dto UserDTO

	if err := repo.db.Table(TbName).Where("email = ?", email).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return dto.ToEntity()
}

func (repo userMySQLRepo) Find(ctx context.Context, id uuid.UUID) (*userdomain.User, error) {
	var dto UserDTO

	if err := repo.db.Table(TbName).Where("id = ?", id).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return dto.ToEntity()
}

func (repo userMySQLRepo) Create(ctx context.Context, data *userdomain.User) error {
	dto := UserDTO{
		Id:        data.Id(),
		FirstName: data.FirstName(),
		LastName:  data.LastName(),
		Email:     data.Email(),
		Password:  data.Password(),
		Salt:      data.Salt(),
		Role:      data.Role().String(),
		Status:    data.Status(),
	}

	if err := repo.db.Table(TbName).Create(&dto).Error; err != nil {
		return err
	}

	return nil
}

func (repo userMySQLRepo) Update(ctx context.Context, data *userdomain.User) error {
	dto := UserDTO{
		Id:        data.Id(),
		FirstName: data.FirstName(),
		LastName:  data.LastName(),
		Email:     data.Email(),
		Password:  data.Password(),
		Salt:      data.Salt(),
		Role:      data.Role().String(),
		Status:    data.Status(),
		Avatar:    GetStrPt(data.Avatar()),
	}

	if err := repo.db.Table(TbName).Where("id = ?", data.Id()).Updates(&dto).Error; err != nil {
		return err
	}

	return nil
}
