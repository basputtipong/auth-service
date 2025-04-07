package httpserv

import (
	"auth-service/infrastructure"
	"auth-service/internal/adaptor/handler"
	repository "auth-service/internal/adaptor/repo"
	"auth-service/internal/core/service"

	"github.com/gin-gonic/gin"
)

func bindLoginRoute(app *gin.Engine) {
	repo := repository.NewLoginRepo(infrastructure.DB)
	svc := service.NewLoginSvc(repo)
	hdl := handler.NewLoginHandler(svc)

	app.POST("/login", hdl.Handle)
}
