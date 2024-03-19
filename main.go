package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	productcontroller "myapp/module/product/controller"
	productusecase "myapp/module/product/domain/usecase"
	productmysql "myapp/module/product/repository/mysql"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello world!")

	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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

	r.Run(":3000")
}
