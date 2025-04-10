package handler

import (
	"auth-service/internal/core/domain"
	"net/http"

	liberror "github.com/basputtipong/library/error"
	"github.com/gin-gonic/gin"
)

type bannerHandler struct {
	svc domain.BannerService
}

func NewBannerHandler(svc domain.BannerService) *bannerHandler {
	return &bannerHandler{svc}
}

func (h *bannerHandler) Handle(c *gin.Context) {
	var req domain.BannerSvcReq

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
	res, err := h.svc.Execute(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, &res)
}
