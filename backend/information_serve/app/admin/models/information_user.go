package models

import (
	"time"

	"go-admin/common/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type InformationUser struct {
	models.Model

	Username                  string    `json:"username" gorm:"type:varchar(64);comment:用户名"`
	Head                      string    `json:"head" gorm:"type:varchar(128);comment:头像"`
	Email                     string    `json:"email" gorm:"type:varchar(128);comment:邮箱"`
	Password                  string    `json:"password" gorm:"type:varchar(128);comment:密码"`
	DeviceId                  string    `json:"deviceId" gorm:"type:varchar(128);comment:设备ID"`
	DeviceType                int64     `json:"deviceType" gorm:"type:tinyint;comment:设备类型：1-iOS，2-Android，3-Web"`
	SystemVersion             string    `json:"systemVersion" gorm:"type:varchar(32);comment:系统版本 (如 16.0, 13.0)"`
	AppVersion                string    `json:"appVersion" gorm:"type:varchar(32);comment:App版本号 (如 1.2.3)"`
	BundleId                  string    `json:"bundleId" gorm:"type:varchar(128);comment:App包名 (如 com.wallet.dex)"`
	WalletAddress             string    `json:"walletAddress" gorm:"type:varchar(128);comment:绑定的钱包地址"`
	Language                  string    `json:"language" gorm:"type:varchar(8);comment:语言：zh-CN（中文），en-US（英文）"`
	Timezone                  string    `json:"timezone" gorm:"type:varchar(32);comment:时区"`
	TwoFactorEnabled          int       `json:"twoFactorEnabled" gorm:"type:tinyint;comment:是否启用两步验证：2-否，1-是"`
	TwoFactorSecret           string    `json:"twoFactorSecret" gorm:"type:varchar(128);comment:两步验证密钥"`
	PointNotificationSettings int64     `json:"pointNotificationSettings" gorm:"type:tinyint;comment:点位权限设置：2-全部，1-部分"`
	NotificationSettings      string    `json:"notificationSettings" gorm:"type:json;comment:消息通知设置[1.资讯 2新点位 3点位状态 4系统]"`
	LastLoginAt               time.Time `json:"lastLoginAt" gorm:"type:datetime(3);comment:最后登录时间"`
	LastLoginIp               string    `json:"lastLoginIp" gorm:"type:varchar(64);comment:最后登录IP"`
	Status                    int64     `json:"status" gorm:"type:tinyint;comment:状态：3-禁用，1-正常，2-锁定"`
	// Google登录相关字段
	GoogleId          string    `json:"googleId" gorm:"type:varchar(128);comment:Google用户ID"`
	GoogleAccessToken  string    `json:"-" gorm:"type:varchar(512);comment:Google访问令牌"`
	GoogleRefreshToken string    `json:"-" gorm:"type:varchar(512);comment:Google刷新令牌"`
	GoogleExpiresAt    time.Time `json:"googleExpiresAt" gorm:"type:datetime(3);comment:Google令牌过期时间"`
	GoogleEmail        string    `json:"googleEmail" gorm:"type:varchar(128);comment:Google邮箱"`
	GoogleName         string    `json:"googleName" gorm:"type:varchar(128);comment:Google用户名"`
	GooglePicture      string    `json:"googlePicture" gorm:"type:varchar(512);comment:Google头像URL"`
	// Microsoft登录相关字段
	MicrosoftId          string    `json:"microsoftId" gorm:"type:varchar(128);comment:Microsoft用户ID"`
	MicrosoftAccessToken  string    `json:"-" gorm:"type:varchar(512);comment:Microsoft访问令牌"`
	MicrosoftRefreshToken string    `json:"-" gorm:"type:varchar(512);comment:Microsoft刷新令牌"`
	MicrosoftExpiresAt    time.Time `json:"microsoftExpiresAt" gorm:"type:datetime(3);comment:Microsoft令牌过期时间"`
	MicrosoftEmail       string    `json:"microsoftEmail" gorm:"type:varchar(128);comment:Microsoft邮箱"`
	MicrosoftName        string    `json:"microsoftName" gorm:"type:varchar(128);comment:Microsoft用户名"`
	MicrosoftPicture     string    `json:"microsoftPicture" gorm:"type:varchar(512);comment:Microsoft头像URL"`
	models.ModelTime
	models.ControlBy
}

func (InformationUser) TableName() string {
	return "information_user"
}

func (e *InformationUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

// Encrypt 加密
func (e *InformationUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *InformationUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *InformationUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}

func (e *InformationUser) GetId() interface{} {
	return e.Id
}
