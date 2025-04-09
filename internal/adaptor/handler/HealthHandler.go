package handler

import (
	"auth-service/infrastructure"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandle(c *gin.Context) {
	err := infrastructure.PingAllDb()
	if err != nil {
		res := map[string]string{
			"status": "unhealthy",
			"msg":    fmt.Sprintf("unhealthy get error: %s", err.Error()),
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
}
