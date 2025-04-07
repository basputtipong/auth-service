package main

import (
	"auth-service/httpserv"
	"auth-service/infrastructure"
)

func init() {
	infrastructure.InitConfig()
}

func main() {
	infrastructure.InitMiddleware()
	infrastructure.InitDB()
	httpserv.Run()
}
