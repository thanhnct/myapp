package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"myapp/builder"
	"myapp/common"
	"myapp/component"
	"myapp/middleware"
	productcontroller "myapp/module/product/controller"
	productusecase "myapp/module/product/domain/usecase"
	productmysql "myapp/module/product/repository/mysql"
	"myapp/module/user/infras/httpservice"
	"myapp/module/user/infras/repository"
	userusecase "myapp/module/user/usecase"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello world!")

	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	jwtSecret := os.Getenv("JWT_SECRET")
	tokenProvider := component.NewJWTProvider(jwtSecret, 60*60*24*7, 60*60*24*14)

	authClient := userusecase.NewIntrospectUC(repository.NewUserRepo(db), repository.NewSessionMySQLRepo(db), tokenProvider)

	r.GET("/ping", middleware.RequireAuth(authClient), func(c *gin.Context) {

		requester := c.MustGet(common.KeyRequester).(common.Requester)

		c.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"requester": requester.LastName(),
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

	repo := productmysql.NewMysqlRepository(db)
	useCase := productusecase.NewCreateProductUseCase(repo)
	api := productcontroller.NewAPIController(useCase)

	v1 := r.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", api.CreateProductAPI(db))
		}
	}

	//userUC := userusecase.NewUserUseCase(repository.NewUserRepo(db), &common.Hasher{}, tokenProvider, repository.NewSessionMySQLRepo(db))
	userUseCase := userusecase.UseCaseWithBuilder(builder.NewComplexBuilder(builder.NewSimpleBuilder(db, tokenProvider)))
	httpservice.NewUserService(userUseCase).Routes(v1)
	r.Run(":3000")
}
