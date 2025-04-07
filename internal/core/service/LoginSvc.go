package service

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/core/port"
	"auth-service/utils"

	liberror "github.com/basputtipong/library/error"
	libmiddleware "github.com/basputtipong/library/middleware"
)

type loginSvc struct {
	loginRepo port.LoginRepo
}

func NewLoginSvc(repo port.LoginRepo) domain.LoginService {
	return &loginSvc{loginRepo: repo}
}

func (s *loginSvc) Execute(req domain.LoginSvcReq) (*domain.LoginSvcRes, error) {
	if err := utils.Validate(req); err != nil {
		return nil, liberror.ErrorBadRequest("Invalid request", err.Error())
	}

	hashedPasscode, err := utils.HashPasscode(req.Passcode)
	if err != nil {
		return nil, liberror.ErrorInternalServerError("failed to hash passcode", err.Error())
	}

	repoReq := port.LoginRepoReq{
		UserId:   req.UserId,
		Passcode: hashedPasscode,
	}
	err = s.loginRepo.Insert(repoReq)
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
