package httpserv

import (
	"auth-service/infrastructure"
	"auth-service/internal/adaptor/handler"
	repository "auth-service/internal/adaptor/repo"
	"auth-service/internal/core/service"

	libmiddleware "github.com/basputtipong/library/middleware"
	"github.com/gin-gonic/gin"
)

func bindLoginRoute(app *gin.Engine) {
	repo := repository.NewUsersRepo(infrastructure.DB)
	svc := service.NewLoginSvc(repo)
	hdl := handler.NewLoginHandler(svc)

	app.POST("/login", hdl.Handle)
}

func bindVerifyRoute(app *gin.Engine) {
	repo := repository.NewUsersRepo(infrastructure.DB)
	svc := service.NewVerifySvc(repo)
	hdl := handler.NewVerifyHandler(svc)

	app.POST(
		"/verify",
		libmiddleware.JWTMiddleware(),
		hdl.Handle,
	)
}

func bindBannerRoute(app *gin.Engine) {
	repo := repository.NewBannerRepo(infrastructure.DB)
	svc := service.NewBannerSvc(repo)
	hdl := handler.NewBannerHandler(svc)

	app.GET("/banner", libmiddleware.JWTMiddleware(), hdl.Handle)
}

func bindHelthRoute(app *gin.Engine) {
	app.GET("/health", handler.HealthHandle)
}
