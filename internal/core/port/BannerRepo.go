package port

type BannerRepo interface {
	GetByUserId(userId string) (Banner, error)
}

const (
	BannersTbl = "banners"
)

type Banner struct {
	BannerId    string `gorm:"column:banner_id"`
	UserId      string `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Image       string `gorm:"column:image"`
}

func (Banner) TableName() string {
	return BannersTbl
}
