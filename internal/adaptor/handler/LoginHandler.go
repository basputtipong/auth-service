package handler

import (
	"auth-service/internal/core/domain"
	"net/http"

	liberror "github.com/basputtipong/library/error"
	"github.com/gin-gonic/gin"
)

type loginHandler struct {
	svc domain.LoginService
}

func NewLoginHandler(svc domain.LoginService) *loginHandler {
	return &loginHandler{svc: svc}
}

func (h *loginHandler) Handle(c *gin.Context) {
	var req domain.LoginSvcReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(liberror.ErrorBadRequest("Invalid request", err.Error()))
		return
	}

	res, err := h.svc.Execute(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, &res)
}
