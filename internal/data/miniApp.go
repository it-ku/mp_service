package data

import "github.com/go-kratos/kratos/v2/log"

type MiniAppEntity struct {
	BaseFields
	Name         string `gorm:"type:varchar(64);not null;comment:小程序名称"`
	OriginalID   string `gorm:"type:varchar(32);not null;default:'';comment:小程序原始id"`
	AppID        string `gorm:"type:varchar(32);not null;default:'';comment:小程序app_id"`
	AppSecret    string `gorm:"type:varchar(64);not null;default:'';comment:小程序秘钥"`
	Platform     string `gorm:"type:varchar(32);not null;comment:平台"`
	RefreshToken uint8  `gorm:"type:tinyint(1);not null;default:'0';comment:是否自动获取access_token 1自动获取"`
}

func (MiniAppEntity) TableName() string {
	return "ik_mini_apps"
}

type MiniAppRepo struct {
	data *Data
	log  *log.Helper
}

func NewMiniAppRepo(data *Data, logger log.Logger) {

}
