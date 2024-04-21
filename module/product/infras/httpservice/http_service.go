package httpservice

import (
	"myapp/common"
	"myapp/middleware"
	"myapp/module/product/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
)

type httpService struct {
	uc         usecase.UseCase
	sctx       sctx.ServiceContext
	authClient middleware.AuthClient
}

func NewHttpService(uc usecase.UseCase, sctx sctx.ServiceContext) httpService {
	return httpService{uc: uc, sctx: sctx}
}

func (s httpService) handleListProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param usecase.ListProductParam

		if err := c.Bind(&param); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		result, err := s.uc.ListProduct(c.Request.Context(), &param)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.SuccessResponse(&result, param.Paging, param.ListProductFilter))
	}
}

func (s httpService) handleCreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.ProductCreateDTO
		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := s.uc.CreateProduct(c.Request.Context(), dto); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (s httpService) handleUpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.ProductUpdateDTO
		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := s.uc.UpdateProduct(c.Request.Context(), dto); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (s httpService) handleDeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := s.uc.DeleteProduct(c.Request.Context(), id); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (s httpService) Routes(g *gin.RouterGroup) {
	products := g.Group("products")
	{
		products.GET("", middleware.RequireAuth(s.authClient), s.handleListProduct())
		products.POST("", middleware.RequireAuth(s.authClient), s.handleCreateProduct())
		products.PATCH("", middleware.RequireAuth(s.authClient), s.handleUpdateProduct())
		products.DELETE("/:id", middleware.RequireAuth(s.authClient), s.handleDeleteProduct())
	}
}

func (s httpService) SetAuthClient(ac middleware.AuthClient) httpService {
	s.authClient = ac
	return s
}
