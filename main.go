package main

import (
	"log"
	"myapp/builder"
	"myapp/common"
	"myapp/component"
	"myapp/middleware"
	imageHTTP "myapp/module/image/infras/httpservice"
	productHTTP "myapp/module/product/infras/httpservice"
	productUsecase "myapp/module/product/usecase"
	userHTTP "myapp/module/user/infras/httpservice"
	"myapp/module/user/infras/repository"
	userUsecase "myapp/module/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/component/gormc"
)

func newService() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName("G11"),
		sctx.WithComponent(gormc.NewGormDB(common.KeyGorm, "")),
		sctx.WithComponent(component.NewJWT(common.KeyJWT)),
		sctx.WithComponent(component.NewAWSS3Provider(common.KeyAWSS3)),
	)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serviceCtx := newService()

	if err := serviceCtx.Load(); err != nil {
		log.Fatalln(err)
	}

	db := serviceCtx.MustGet(common.KeyGorm).(common.DbContext).GetDB()

	r := gin.Default()

	r.Use(middleware.Recovery())

	tokenProvider := serviceCtx.MustGet(common.KeyJWT).(component.TokenProvider)

	authClient := userUsecase.NewIntrospectUC(repository.NewUserRepo(db), repository.NewSessionMySQLRepo(db), tokenProvider)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.DELETE("/v1/revoke-token", middleware.RequireAuth(authClient), func(c *gin.Context) {
		requester := c.MustGet(common.KeyRequester).(common.Requester)
		repo := repository.NewSessionMySQLRepo(db)
		if err := repo.Delete(c.Request.Context(), requester.TokenId()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	})

	v1 := r.Group("/v1")

	//userUC := userUsecase.NewuserUsecase(repository.NewUserRepo(db), &common.Hasher{}, tokenProvider, repository.NewSessionMySQLRepo(db))
	userUC := userUsecase.UseCaseWithBuilder(builder.NewComplexBuilder(builder.NewSimpleBuilder(db, tokenProvider)))
	userHTTP.NewUserService(userUC, serviceCtx).SetAuthClient(authClient).Routes(v1)

	productUC := productUsecase.UseCaseWithBuilder(builder.NewSimpleBuilder(db, nil))
	productHTTP.NewHttpService(productUC, serviceCtx).SetAuthClient(authClient).Routes(v1)

	imageHTTP.NewHTTPService(serviceCtx).Routes(v1)

	r.Run(":3000")
}
