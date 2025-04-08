package service

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/core/port"
	"auth-service/utils"

	liberror "github.com/basputtipong/library/error"
	libmiddleware "github.com/basputtipong/library/middleware"
)

type loginSvc struct {
	usersRepo port.UsersRepo
}

func NewLoginSvc(repo port.UsersRepo) domain.LoginService {
	return &loginSvc{usersRepo: repo}
}

func (s *loginSvc) Execute(req domain.LoginSvcReq) (*domain.LoginSvcRes, error) {
	if err := utils.Validate(req); err != nil {
		return nil, liberror.ErrorBadRequest("Invalid request", err.Error())
	}

	hashedPasscode, err := utils.HashPasscode(req.Passcode)
	if err != nil {
		return nil, liberror.ErrorInternalServerError("failed to hash passcode", err.Error())
	}

	repoReq := port.UsersRepoReq{
		UserId:   req.UserId,
		Passcode: hashedPasscode,
	}
	err = s.usersRepo.Insert(repoReq)
	if err != nil {
		return nil, err
	}

	var res domain.LoginSvcRes
	token, err := libmiddleware.GenerateJWT(req.UserId)
	if err != nil {
		return nil, liberror.ErrorInternalServerError("failed to generate token", err.Error())
	}

	res.AuthToken = token
	return &res, nil
}
