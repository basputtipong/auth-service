package repository

import (
	"auth-service/internal/core/port"
	"errors"

	liberror "github.com/basputtipong/library/error"
	"gorm.io/gorm"
)

type bannerRepo struct {
	db *gorm.DB
}

func NewBannerRepo(db *gorm.DB) port.BannerRepo {
	return &bannerRepo{db: db}
}

func (r *bannerRepo) GetByUserId(userId string) (port.Banner, error) {
	var repoRes port.Banner
	res := r.db.Where("user_id = ?", userId).First(&repoRes)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return repoRes, liberror.ErrorNotFound("banner not found", gorm.ErrRecordNotFound.Error())
		} else {
			return repoRes, liberror.ErrorInternalServerError("failed to retrieve data", res.Error.Error())
		}
	}

	return repoRes, nil
}
