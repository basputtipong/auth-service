package domain

type VerifyService interface {
	Execute(req VerifySvcReq) (VerifySvcRes, error)
}

type VerifySvcReq struct {
	UserId   string `json:"userId" validate:"required"`
	Passcode string `json:"passcode" validate:"required"`
}

type VerifySvcRes struct {
	Name        string `json:"name"`
	GreetingMsg string `json:"greetingMsg"`
}
