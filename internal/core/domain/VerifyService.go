package domain

type VerifyService interface {
	Execute(req VerifySvcReq) (EmptyResponse, error)
}

type VerifySvcReq struct {
	UserId   string `json:"userId" validate:"required"`
	Passcode string `json:"passcode" validate:"required"`
}

type EmptyResponse struct{}
