package builder

import (
	"gorm.io/gorm"
	"myapp/common"
	"myapp/module/user/infras/repository"
	userusecase "myapp/module/user/usecase"
)

type simpleBuilder struct {
	db *gorm.DB
	tp userusecase.TokenProvider
}

func NewSimpleBuilder(db *gorm.DB, tp userusecase.TokenProvider) simpleBuilder {
	return simpleBuilder{db: db, tp: tp}
}

func (s simpleBuilder) BuildUserQueryRepo() userusecase.UserQueryRepository {
	return repository.NewUserRepo(s.db)
}

func (s simpleBuilder) BuildUserCmdRepo() userusecase.UserCommandRepository {
	return repository.NewUserRepo(s.db)
}

func (simpleBuilder) BuildHasher() userusecase.Hasher {
	return &common.Hasher{}
}

func (s simpleBuilder) BuildTokenProvider() userusecase.TokenProvider {
	return s.tp
}

func (s simpleBuilder) BuildSessionQueryRepo() userusecase.SessionQueryRepository {
	return repository.NewSessionMySQLRepo(s.db)
}

func (s simpleBuilder) BuildSessionCmdRepo() userusecase.SessionCommandRepository {
	return repository.NewSessionMySQLRepo(s.db)
}

func (s simpleBuilder) BuildSessionRepo() userusecase.SessionRepository {
	return repository.NewSessionMySQLRepo(s.db)
}

// Complex builder

func NewComplexBuilder(simpleBuilder simpleBuilder) complexBuilder {
	return complexBuilder{simpleBuilder: simpleBuilder}
}

type complexBuilder struct {
	simpleBuilder
}

// Proxy design pattern
//type userCacheRepo struct {
//	realRepo userusecase.UserQueryRepository
//	cache    map[string]*userdomain.User
//}
//
//func (c userCacheRepo) FindByEmail(ctx context.Context, email string) (*userdomain.User, error) {
//	if user, ok := c.cache[email]; ok {
//		return user, nil
//	}
//
//	user, err := c.realRepo.FindByEmail(ctx, email)
//
//	if err != nil {
//		return nil, err
//	}
//
//	c.cache[email] = user
//
//	return user, nil
//}
//
//func (cb complexBuilder) BuildUserQueryRepo() userusecase.UserQueryRepository {
//	return userCacheRepo{
//		realRepo: cb.simpleBuilder.BuildUserQueryRepo(),
//		cache:    make(map[string]*userdomain.User),
//	}
//}
