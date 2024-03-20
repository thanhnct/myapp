package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"myapp/builder"
	"myapp/component"
	productcontroller "myapp/module/product/controller"
	productusecase "myapp/module/product/domain/usecase"
	productmysql "myapp/module/product/repository/mysql"
	"myapp/module/user/infras/httpservice"
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

	repo := productmysql.NewMysqlRepository(db)
	useCase := productusecase.NewCreateProductUseCase(repo)
	api := productcontroller.NewAPIController(useCase)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		products := v1.Group("/products")
		{
			products.POST("", api.CreateProductAPI(db))
		}
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenProvider := component.NewJWTProvider(jwtSecret, 60*60*24*7, 60*60*24*14)
	//userUC := userusecase.NewUserUseCase(repository.NewUserRepo(db), &common.Hasher{}, tokenProvider, repository.NewSessionMySQLRepo(db))
	userUseCase := userusecase.UseCaseWithBuilder(builder.NewComplexBuilder(builder.NewSimpleBuilder(db, tokenProvider)))
	httpservice.NewUserService(userUseCase).Routes(v1)
	r.Run(":3000")
}
