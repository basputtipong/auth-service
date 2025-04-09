package port

type UsersRepo interface {
	Insert(req UsersRepoReq) error
	GetByUserId(userId string) (Users, error)
	GetGreetingByUserId(userId string) (UsersWithGreeting, error)
}

const (
	UsersTbl         = "users"
	UserGreetingsTbl = "user_greetings"
)

type UsersRepoReq struct {
	UserId   string `json:"userId"`
	Passcode string `json:"passcode"`
}

type UsersWithGreeting struct {
	UserId   string `gorm:"column:user_id"`
	Name     string `gorm:"column:name"`
	Passcode string `gorm:"column:user_passcode"`
	Greeting string `gorm:"column:greeting"`
}

type Users struct {
	UserId   string `gorm:"column:user_id"`
	Name     string `gorm:"column:name"`
	Passcode string `gorm:"column:user_passcode"`
}

func (Users) TableName() string {
	return "users"
}
