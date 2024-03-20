package userusecase

import (
	"context"
	"errors"
	"myapp/common"
	userdomain "myapp/module/user/domain"
)

type UseCase interface {
	Register(ctx context.Context, registerDto EmailPasswordRegistrationDTO) error
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	//CompareHashPassword(hashedPassword, salt, password string) bool
}

type useCase struct {
	repo   UserRepository
	hasher Hasher
}

func NewUserUseCase(repo UserRepository, hasher Hasher) UseCase {
	return useCase{repo: repo, hasher: hasher}
}

func (uc useCase) Register(ctx context.Context, registerDto EmailPasswordRegistrationDTO) error {
	// 1. Find user by email:
	// 1.1 Found: return error (email has existed)
	// 2. Generate salt
	// 3. Hash password+salt
	// 4. Create user entity

	user, err := uc.repo.FindByEmail(ctx, registerDto.Email)

	if user != nil {
		return userdomain.ErrEmailHasExisted
	}

	if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
		return err
	}

	salt, err := uc.hasher.RandomStr(30)

	if err != nil {
		return err
	}

	hashedPassword, err := uc.hasher.HashPassword(salt, registerDto.Password)

	if err != nil {
		return err
	}

	userEntity, err := userdomain.NewUser(
		common.GenUUID(),
		registerDto.FirstName,
		registerDto.LastName,
		registerDto.Email,
		hashedPassword,
		salt,
		userdomain.RoleUser,
	)

	if err != nil {
		return err
	}

	if err := uc.repo.Create(ctx, userEntity); err != nil {
		return err
	}

	return nil
}

type UserRepository interface {
	//Find(ctx context.Context, id uuid.UUID) (*userdomain.User, error)
	FindByEmail(ctx context.Context, email string) (*userdomain.User, error)
	Create(ctx context.Context, data *userdomain.User) error
	//Update(ctx context.Context, data userdomain.User) error
	//Delete(ctx context.Context, data userdomain.User) error
}
