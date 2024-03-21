package userusecase

import (
	"context"
	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"myapp/common"
	userdomain "myapp/module/user/domain"
)

type registerUC struct {
	userQueryRepo UserQueryRepository
	userCmdRepo   UserCommandRepository
	hasher        Hasher
}

func NewRegisterUC(userQueryRepo UserQueryRepository, userCmdRepo UserCommandRepository, hasher Hasher) *registerUC {
	return &registerUC{userQueryRepo: userQueryRepo, userCmdRepo: userCmdRepo, hasher: hasher}
}

func (uc *registerUC) Register(ctx context.Context, registerDto EmailPasswordRegistrationDTO) error {
	// 1. Find user by email:
	// 1.1 Found: return error (email has existed)
	// 2. Generate salt
	// 3. Hash password+salt
	// 4. Create user entity

	user, err := uc.userQueryRepo.FindByEmail(ctx, registerDto.Email)

	if user != nil {
		return core.ErrBadRequest.WithError(userdomain.ErrEmailHasExisted.Error())
	}

	if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
		return core.ErrInternalServerError.WithError("can not register right now").WithDebug(err.Error())
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
		common.Activated,
	)

	if err != nil {
		return err
	}

	if err := uc.userCmdRepo.Create(ctx, userEntity); err != nil {
		return err
	}

	return nil
}
