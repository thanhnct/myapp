package userusecase

import (
	"context"
	"github.com/google/uuid"
	userdomain "myapp/module/user/domain"
)

type UseCase interface {
	Register(ctx context.Context, registerDto EmailPasswordRegistrationDTO) error
	LoginEmailPassword(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponseDTO, error)
}

type Hasher interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, err error)
	TokenExpireInSeconds() int
	RefreshExpireInSeconds() int
}

type useCase struct {
	*loginEmailPasswordUC
	*registerUC
	*refreshTokenUC
}

type Builder interface {
	BuildUserQueryRepo() UserQueryRepository
	BuildUserCmdRepo() UserCommandRepository
	BuildHasher() Hasher
	BuildTokenProvider() TokenProvider
	BuildSessionQueryRepo() SessionQueryRepository
	BuildSessionCmdRepo() SessionCommandRepository
	BuildSessionRepo() SessionRepository
}

func UseCaseWithBuilder(b Builder) UseCase {
	return &useCase{
		registerUC:           NewRegisterUC(b.BuildUserQueryRepo(), b.BuildUserCmdRepo(), b.BuildHasher()),
		loginEmailPasswordUC: NewLoginEmailPasswordUC(b.BuildUserQueryRepo(), b.BuildSessionCmdRepo(), b.BuildTokenProvider(), b.BuildHasher()),
		refreshTokenUC:       NewRefreshTokenPasswordUC(b.BuildUserQueryRepo(), b.BuildSessionRepo(), b.BuildTokenProvider(), b.BuildHasher()),
	}
}

func NewUserUseCase(repo UserRepository, hasher Hasher, tokenProvider TokenProvider, sessionRepo SessionRepository) UseCase {
	return &useCase{
		loginEmailPasswordUC: NewLoginEmailPasswordUC(repo, sessionRepo, tokenProvider, hasher),
		registerUC:           NewRegisterUC(repo, repo, hasher),
	}
}

type UserRepository interface {
	UserQueryRepository
	UserCommandRepository
}

type UserQueryRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*userdomain.User, error)
	FindByEmail(ctx context.Context, email string) (*userdomain.User, error)
}

type UserCommandRepository interface {
	Create(ctx context.Context, data *userdomain.User) error
	//Update(ctx context.Context, data userdomain.User) error
	//Delete(ctx context.Context, data userdomain.User) error
}

type SessionRepository interface {
	SessionQueryRepository
	SessionCommandRepository
}

type SessionQueryRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*userdomain.Session, error)
	FindByRefreshToken(ctx context.Context, rt string) (*userdomain.Session, error)
}

type SessionCommandRepository interface {
	Create(ctx context.Context, data *userdomain.Session) error
	Delete(ctx context.Context, id uuid.UUID) error
}
