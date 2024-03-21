package userusecase

import (
	"context"
	"errors"
	"myapp/common"
	userdomain "myapp/module/user/domain"
	"time"
)

type refreshTokenUC struct {
	userRepo           UserQueryRepository
	sessionQueryRepo   SessionQueryRepository
	sessionCommandRepo SessionCommandRepository
	tokenProvider      TokenProvider
	hasher             Hasher
}

func NewRefreshTokenPasswordUC(userRepo UserQueryRepository, sessionQueryRepo SessionQueryRepository,
	tokenProvider TokenProvider, hasher Hasher, sessionCommandRepo SessionCommandRepository) *refreshTokenUC {
	return &refreshTokenUC{userRepo: userRepo, sessionQueryRepo: sessionQueryRepo, sessionCommandRepo: sessionCommandRepo, tokenProvider: tokenProvider, hasher: hasher}
}

func (uc *refreshTokenUC) RefreshToken(ctx context.Context, refreshToken string) (*TokenResponseDTO, error) {
	session, err := uc.sessionQueryRepo.FindByRefreshToken(ctx, refreshToken)

	if err != nil {
		return nil, err
	}

	if session.RefreshExpAt().UnixNano() < time.Now().UTC().UnixNano() {
		return nil, errors.New("refresh token has expired")
	}

	user, err := uc.userRepo.Find(ctx, session.UserId())

	if err != nil {
		return nil, err
	}

	if user.Status() == "banned" {
		return nil, errors.New("user has been banned")
	}

	userId := user.Id()
	sessionId := common.GenUUID()

	// 3. Gen JWT
	accessToken, err := uc.tokenProvider.IssueToken(ctx, sessionId.String(), userId.String())

	if err != nil {
		return nil, err
	}

	// 4. Insert session into DB
	newRefreshToken, _ := uc.hasher.RandomStr(16)
	tokenExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.TokenExpireInSeconds()))
	refreshExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.RefreshExpireInSeconds()))

	newSession := userdomain.NewSession(sessionId, userId, newRefreshToken, tokenExpAt, refreshExpAt)

	if err := uc.sessionCommandRepo.Create(ctx, newSession); err != nil {
		return nil, err
	}

	go func() {
		_ = uc.sessionCommandRepo.Delete(ctx, session.Id())
	}()

	// 5. Return token response dto

	return &TokenResponseDTO{
		AccessToken:       accessToken,
		AccessTokenExpIn:  uc.tokenProvider.TokenExpireInSeconds(),
		RefreshToken:      newRefreshToken,
		RefreshTokenExpIn: uc.tokenProvider.RefreshExpireInSeconds(),
	}, nil
}
