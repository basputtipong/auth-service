package httpserv

import (
	liberror "github.com/basputtipong/library/error"
	libmiddleware "github.com/basputtipong/library/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	app := gin.Default()
	app.Use(libmiddleware.CORSMiddleware())
	app.Use(liberror.ErrorHandler())

	bindLoginRoute(app)
	bindVerifyRoute(app)

	port := viper.GetString("app.port")
	if port == "" {
		port = "8080"
	}

	app.Run(":" + port)
}
