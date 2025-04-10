package repository

import (
	"auth-service/internal/core/port"
	"errors"
	"fmt"

	liberror "github.com/basputtipong/library/error"
	"gorm.io/gorm"
)

type usersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) port.UsersRepo {
	return &usersRepo{db: db}
}

func (r *usersRepo) UpdatePasscodeByUserId(req port.UsersRepoReq) error {
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

func (r *usersRepo) GetByUserId(userId string) (port.Users, error) {
	var repoRes port.Users
	res := r.db.Where("user_id = ?", userId).First(&repoRes)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return repoRes, liberror.ErrorNotFound("User not found", "")
		} else {
			return repoRes, liberror.ErrorInternalServerError("failed to retrieve data", res.Error.Error())
		}
	}

	return repoRes, nil
}

func (r *usersRepo) GetGreetingByUserId(userId string) (port.UsersWithGreeting, error) {
	var repoRes port.UsersWithGreeting
	res := r.db.Table(port.UsersTbl+" AS u").
		Select("u.user_id, u.name, u.user_passcode, ug.greeting").
		Joins(fmt.Sprintf(`LEFT JOIN %s AS ug ON u.user_id = ug.user_id`, port.UserGreetingsTbl)).
		Where("u.user_id = ?", userId).
		Scan(&repoRes)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return repoRes, liberror.ErrorNotFound("User not found", "")
		} else {
			return repoRes, liberror.ErrorInternalServerError("failed to retrieve data", res.Error.Error())
		}
	}

	return repoRes, nil
}
