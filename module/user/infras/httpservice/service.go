package httpservice

import (
	"myapp/common"
	"myapp/middleware"
	imageRepo "myapp/module/image/infras/repository"
	"myapp/module/user/infras/repository"
	userUsecase "myapp/module/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
)

type service struct {
	uc         userUsecase.UseCase
	sctx       sctx.ServiceContext
	authClient middleware.AuthClient
}

func NewUserService(uc userUsecase.UseCase, sctx sctx.ServiceContext) service {
	return service{uc: uc, sctx: sctx}
}

func (s service) handleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto userUsecase.EmailPasswordRegistrationDTO
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
		var dto userUsecase.EmailPasswordLoginDTO
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

func (s service) handleChangeAvatar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto userUsecase.SingleImageDTO

		if err := c.BindJSON(&dto); err != nil {
			common.WriteErrorResponse(c, core.ErrBadRequest.WithError(err.Error()))
			return
		}

		dto.Requester = c.MustGet(common.KeyRequester).(common.Requester)

		dbCtx := s.sctx.MustGet(common.KeyGorm).(common.DbContext)

		userRepo := repository.NewUserRepo(dbCtx.GetDB())
		imgRepo := imageRepo.NewRepo(dbCtx.GetDB())

		if err := userUsecase.NewChangeAvtUC(userRepo, userRepo, imgRepo).ChangeAvatar(c.Request.Context(), dto); err != nil {
			common.WriteErrorResponse(c, err)
			return
		}

		c.JSON(http.StatusOK, core.ResponseData(true))
	}
}

func (s service) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
	g.POST("/authenticate", s.handleLogin())
	g.POST("/refresh-token", s.handleRefreshToken())
	g.PATCH("/profile/change-avatar", middleware.RequireAuth(s.authClient), s.handleChangeAvatar()) // RPC-restful
}

func (s service) SetAuthClient(ac middleware.AuthClient) service {
	s.authClient = ac
	return s
}
