package service

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/core/port"
	"auth-service/utils"

	liberror "github.com/basputtipong/library/error"
)

type verifySvc struct {
	usersRepo port.UsersRepo
}

func NewVerifySvc(repo port.UsersRepo) domain.VerifyService {
	return &verifySvc{usersRepo: repo}
}

func (s *verifySvc) Execute(req domain.VerifySvcReq) (domain.VerifySvcRes, error) {
	var res domain.VerifySvcRes
	if err := utils.Validate(req); err != nil {
		return res, liberror.ErrorBadRequest("Invalid request", err.Error())
	}

	repoRes, err := s.usersRepo.GetGreetingByUserId(req.UserId)
	if err != nil {
		return res, err
	}

	if err := utils.ComparePasscode(repoRes.Passcode, req.Passcode); err != nil {
		return res, liberror.ErrorInternalServerError("passcode is not matched", "")
	}

	res.Name = repoRes.Name
	res.GreetingMsg = repoRes.Greeting
	return res, nil
}
