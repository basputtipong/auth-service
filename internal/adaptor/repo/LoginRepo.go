package repository

import (
	"auth-service/internal/core/port"

	liberror "github.com/basputtipong/library/error"
	"gorm.io/gorm"
)

type loginRepo struct {
	db *gorm.DB
}

func NewLoginRepo(db *gorm.DB) port.LoginRepo {
	return &loginRepo{db: db}
}

func (r *loginRepo) Insert(req port.LoginRepoReq) error {
	tx := r.db.Begin()

	res := tx.Model(&port.Users{}).
		Where("user_id = ?", req.UserId).
		Update("user_passcode", req.Passcode)

	if res.Error != nil {
		tx.Rollback()
		return liberror.ErrorInternalServerError("failed to update", res.Error.Error())
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return liberror.ErrorNotFound("User not found", "")
	}

	if err := tx.Commit().Error; err != nil {
		return liberror.ErrorInternalServerError("failed to commit transaction", err.Error())
	}

	return nil
}
