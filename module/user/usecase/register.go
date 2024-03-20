package userusecase

import (
	"context"
	"github.com/pkg/errors"
	"myapp/common"
	userdomain "myapp/module/user/domain"
)

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
