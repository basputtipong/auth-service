package infrastructure

import (
	libmiddleware "github.com/basputtipong/library/middleware"
)

func InitMiddleware() {
	libmiddleware.Init()
	libmiddleware.InitCorsConfig()
}
