package userusecase

import (
	"context"
	userdomain "myapp/module/user/domain"
)

type UseCase interface {
	Register(ctx context.Context, registerDto EmailPasswordRegistrationDTO) error
	LoginEmailPassword(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error)
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
	repo              UserRepository
	sessionRepository SessionRepository
	hasher            Hasher
	tokenProvider     TokenProvider
}

func NewUserUseCase(repo UserRepository, hasher Hasher, tokenProvider TokenProvider, sessionRepository SessionRepository) UseCase {
	return useCase{
		repo:              repo,
		sessionRepository: sessionRepository,
		hasher:            hasher,
		tokenProvider:     tokenProvider,
	}
}

type UserRepository interface {
	//Find(ctx context.Context, id uuid.UUID) (*userdomain.User, error)
	FindByEmail(ctx context.Context, email string) (*userdomain.User, error)
	Create(ctx context.Context, data *userdomain.User) error
	//Update(ctx context.Context, data userdomain.User) error
	//Delete(ctx context.Context, data userdomain.User) error
}

type SessionRepository interface {
	Create(ctx context.Context, data *userdomain.Session) error
}
