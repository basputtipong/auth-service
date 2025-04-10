package service

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/core/port"
	"auth-service/utils"

	liberror "github.com/basputtipong/library/error"
)

type bannerSvc struct {
	bannerRepo port.BannerRepo
}

func NewBannerSvc(bannerRepo port.BannerRepo) domain.BannerService {
	return &bannerSvc{bannerRepo}
}

func (s *bannerSvc) Execute(req domain.BannerSvcReq) (domain.BannerSvcRes, error) {
	var res domain.BannerSvcRes
	if err := utils.Validate(req); err != nil {
		return res, liberror.ErrorBadRequest("Invalid request", err.Error())
	}

	repoRes, err := s.bannerRepo.GetByUserId(req.UserId)
	if err != nil {
		return res, err
	}

	res.BuildBannerRes(repoRes)
	return res, nil
}
