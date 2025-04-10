package service

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/core/port"
	"auth-service/utils"

	liberror "github.com/basputtipong/library/error"
	libmiddleware "github.com/basputtipong/library/middleware"
)

type loginSvc struct {
	usersRepo    port.UsersRepo
	jwtGenerator libmiddleware.JWTGenerator
}

func NewLoginSvc(repo port.UsersRepo, jwtGenerator libmiddleware.JWTGenerator) domain.LoginService {
	return &loginSvc{
		usersRepo:    repo,
		jwtGenerator: jwtGenerator,
	}
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
	err = s.usersRepo.UpdatePasscodeByUserId(repoReq)
	if err != nil {
		return nil, err
	}

	var res domain.LoginSvcRes
	token, err := s.jwtGenerator.Generate(req.UserId)
	if err != nil {
		return nil, liberror.ErrorInternalServerError("failed to generate token", err.Error())
	}

	res.AuthToken = token
	return &res, nil
}
