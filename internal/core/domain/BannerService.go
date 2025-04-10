package domain

import "auth-service/internal/core/port"

type BannerService interface {
	Execute(req BannerSvcReq) (BannerSvcRes, error)
}

type BannerSvcReq struct {
	UserId string `json:"userId" validate:"required"`
}

type BannerSvcRes struct {
	BannerId    string `json:"bannerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (res *BannerSvcRes) BuildBannerRes(repoRes port.Banner) {
	res.BannerId = repoRes.BannerId
	res.Title = repoRes.Title
	res.Description = repoRes.Description
	res.Image = repoRes.Image
}
