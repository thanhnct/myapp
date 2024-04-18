package usecase

import (
	"context"
	"myapp/common"
	userDomain "myapp/module/user/domain"
	"time"
)

type loginEmailPasswordUC struct {
	userRepo      UserQueryRepository
	sessionRepo   SessionCommandRepository
	tokenProvider TokenProvider
	hasher        Hasher
}

func NewLoginEmailPasswordUC(userRepo UserQueryRepository, sessionRepo SessionCommandRepository,
	tokenProvider TokenProvider, hasher Hasher) *loginEmailPasswordUC {
	return &loginEmailPasswordUC{userRepo: userRepo, sessionRepo: sessionRepo, tokenProvider: tokenProvider, hasher: hasher}
}

func (uc *loginEmailPasswordUC) LoginEmailPassword(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error) {
	user, err := uc.userRepo.FindByEmail(ctx, dto.Email)
	if err != nil {
		return nil, err
	}

	if ok := uc.hasher.CompareHashPassword(user.Password(), user.Salt(), dto.Password); !ok {
		return nil, userDomain.ErrInvalidEmailPassword
	}

	userId := user.Id()
	sessionId := common.GenUUID()

	//Gen JWT
	accessToken, err := uc.tokenProvider.IssueToken(ctx, sessionId.String(), userId.String())
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.hasher.RandomStr(16)
	if err != nil {
		return nil, err
	}
	tokenExpAt := time.Now().UTC().Add(time.Duration(uc.tokenProvider.TokenExpireInSeconds()) * time.Second)
	refreshExpAt := time.Now().UTC().Add(time.Duration(uc.tokenProvider.RefreshExpireInSeconds()) * time.Second)
	session := userDomain.NewSession(sessionId, userId, refreshToken, tokenExpAt, refreshExpAt)

	//Save session to db
	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	//Return token response dto
	return &TokenResponseDTO{
		AccessToken:       accessToken,
		AccessTokenExpIn:  uc.tokenProvider.TokenExpireInSeconds(),
		RefreshToken:      refreshToken,
		RefreshTokenExpIn: uc.tokenProvider.RefreshExpireInSeconds(),
	}, nil
}
