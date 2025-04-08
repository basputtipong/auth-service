package service

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/core/port"
	"auth-service/utils"
	"fmt"

	liberror "github.com/basputtipong/library/error"
)

type verifySvc struct {
	usersRepo port.UsersRepo
}

func NewVerifySvc(repo port.UsersRepo) domain.VerifyService {
	return &verifySvc{usersRepo: repo}
}

func (s *verifySvc) Execute(req domain.VerifySvcReq) (domain.EmptyResponse, error) {
	var emptyRes domain.EmptyResponse
	if err := utils.Validate(req); err != nil {
		return emptyRes, liberror.ErrorBadRequest("Invalid request", err.Error())
	}

	repoRes, err := s.usersRepo.GetByUserId(req.UserId)
	if err != nil {
		return emptyRes, err
	}

	fmt.Println(repoRes.Passcode, req.Passcode)

	if err := utils.ComparePasscode(repoRes.Passcode, req.Passcode); err != nil {
		return emptyRes, liberror.ErrorInternalServerError("passcode is not matched", err.Error())
	}

	return emptyRes, nil
}
