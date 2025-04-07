package port

type LoginRepo interface {
	Insert(LoginRepoReq) error
}

type LoginRepoReq struct {
	UserId   string `json:"userId"`
	Passcode string `json:"passcode"`
}

type Users struct {
	UserId   string `gorm:"column:user_id"`
	Name     string `gorm:"column:name"`
	Passcode string `gorm:"column:user_passcode"`
}

func (Users) TableName() string {
	return "users"
}
