package handler

import (
	"auth-service/internal/core/domain"
	"net/http"

	liberror "github.com/basputtipong/library/error"
	"github.com/gin-gonic/gin"
)

type verifyHandler struct {
	svc domain.VerifyService
}

func NewVerifyHandler(svc domain.VerifyService) *verifyHandler {
	return &verifyHandler{svc}
}

func (h *verifyHandler) Handle(c *gin.Context) {
	var req domain.VerifySvcReq

	userIdRaw, ok := c.Get("user_id")
	if !ok {
		c.Error(liberror.ErrorBadRequest("Invalid request", "user_id missing from context"))
		return
	}

	userId, ok := userIdRaw.(string)
	if !ok {
		c.Error(liberror.ErrorBadRequest("Invalid request", "user_id must be string"))
		return
	}

	req.UserId = userId
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
