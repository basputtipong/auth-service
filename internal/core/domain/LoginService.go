package domain

type LoginService interface {
	Execute(req LoginSvcReq) (*LoginSvcRes, error)
}

type LoginSvcReq struct {
	UserId   string `json:"userId" validate:"required"`
	Passcode string `json:"passcode" validate:"required"`
}

type LoginSvcRes struct {
	AuthToken string `json:"authToken"`
}
