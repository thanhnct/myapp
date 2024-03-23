package httpservice

import (
	"myapp/common"
	userusecase "myapp/module/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
)

type service struct {
	uc   userusecase.UseCase
	sctx sctx.ServiceContext
}

func NewUserService(uc userusecase.UseCase, sctx sctx.ServiceContext) service {
	return service{uc: uc, sctx: sctx}
}

func (s service) handleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto userusecase.EmailPasswordRegistrationDTO
		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		if err := s.uc.Register(c.Request.Context(), dto); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func (s service) handleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto userusecase.EmailPasswordLoginDTO
		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		resp, err := s.uc.LoginEmailPassword(c.Request.Context(), dto)
		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": resp,
		})
	}
}

func (s service) handleRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bodyData struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.BindJSON(&bodyData); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		data, err := s.uc.RefreshToken(c.Request.Context(), bodyData.RefreshToken)

		if err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func (s service) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
	g.POST("/authenticate", s.handleLogin())
	g.POST("/refresh-token", s.handleRefreshToken())
}
