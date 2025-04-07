package httpserv

import (
	liberror "github.com/basputtipong/library/error"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	app := gin.Default()
	app.Use(liberror.ErrorHandler())

	bindLoginRoute(app)

	port := viper.GetString("app.port")
	if port == "" {
		port = "8080"
	}

	app.Run(":" + port)
}
