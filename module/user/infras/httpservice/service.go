package httpservice

import (
	"github.com/gin-gonic/gin"
	userusecase "myapp/module/user/usecase"
	"net/http"
)

type service struct {
	uc userusecase.UseCase
}

func NewUserService(uc userusecase.UseCase) service {
	return service{uc: uc}
}

func (s service) handleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto userusecase.EmailPasswordRegistrationDTO
		if err := c.BindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := s.uc.Register(c.Request.Context(), dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
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
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		resp, err := s.uc.LoginEmailPassword(c.Request.Context(), dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": resp,
		})
	}
}

func (s service) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
	g.POST("/authenticate", s.handleLogin())
}
