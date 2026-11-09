package handler

import (
	"go-admin/common/models"
	"time"
)

type User struct {
	Id               int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"id"`
	Username         string `json:"username" gorm:"size:64;comment:用户名"`
	Password         string `json:"-" gorm:"size:128;comment:密码"`
	Head             string `json:"head" gorm:"size:255;comment:头像"`
	Email            string `json:"email" gorm:"size:128;comment:邮箱"`
	Status           int64  `json:"status" gorm:"size:4;comment:状态"`
	TwoFactorSecret  string `json:"-" gorm:"size:255;comment:2FA密钥"`
	TwoFactorEnabled int    `json:"twoFactorEnabled" gorm:"size:4;comment:是否启用2FA"`
	models.ControlBy
	models.ModelTime
	WalletAddress             string    `json:"walletAddress" gorm:"type:varchar(128);comment:绑定的钱包地址"`
	PointNotificationSettings int64     `json:"pointNotificationSettings" gorm:"type:tinyint;comment:点位权限设置：2-全部，1-部分"`
	Language                  string    `json:"language" gorm:"type:varchar(8);comment:语言：zh-CN（中文），en-US（英文）"`
	BundleId                  string    `json:"bundleId" gorm:"type:varchar(128);comment:App包名 (如 com.wallet.dex)"`
	DeviceId                  string    `json:"deviceId" gorm:"type:varchar(128);comment:设备ID"`
	DeviceType                int64     `json:"deviceType" gorm:"type:tinyint;comment:设备类型：1-iOS，2-Android，3-Web"`
	SystemVersion             string    `json:"systemVersion" gorm:"type:varchar(32);comment:系统版本 (如 16.0, 13.0)"`
	AppVersion                string    `json:"appVersion" gorm:"type:varchar(32);comment:App版本号 (如 1.2.3)"`
	NotificationSettings      string    `json:"notificationSettings" gorm:"type:json;comment:消息通知设置[1.资讯 2新点位 3点位状态 4系统]"`
	LastLoginAt               time.Time `json:"lastLoginAt" gorm:"type:datetime(3);comment:最后登录时间"`
	GoogleId                  string    `json:"googleId" gorm:"size:255;comment:Google用户ID"`
	GoogleEmail               string    `json:"googleEmail" gorm:"size:255;comment:Google邮箱"`
	GoogleName                string    `json:"googleName" gorm:"size:255;comment:Google用户名"`
	GooglePicture             string    `json:"googlePicture" gorm:"size:255;comment:Google头像URL"`
	GoogleAccessToken         string    `json:"googleAccessToken" gorm:"size:255;comment:Google访问令牌"`
	GoogleExpiresAt           time.Time `json:"googleExpiresAt" gorm:"type:datetime(3);comment:Google令牌过期时间"`
	MicrosoftId               string    `json:"microsoftId" gorm:"size:255;comment:Microsoft用户ID"`
	MicrosoftEmail            string    `json:"microsoftEmail" gorm:"size:255;comment:Microsoft邮箱"`
	MicrosoftName             string    `json:"microsoftName" gorm:"size:255;comment:Microsoft用户名"`
	MicrosoftPicture          string    `json:"microsoftPicture" gorm:"size:255;comment:Microsoft头像URL"`
	MicrosoftAccessToken      string    `json:"microsoftAccessToken" gorm:"size:255;comment:Microsoft访问令牌"`
	MicrosoftExpiresAt        time.Time `json:"microsoftExpiresAt" gorm:"type:datetime(3);comment:Microsoft令牌过期时间"`
}

func (*User) TableName() string {
	return "information_user"
}
