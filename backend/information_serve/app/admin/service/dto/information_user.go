package dto

import (
	"strconv"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InformationUserGetPageReq struct {
	dto.Pagination            `search:"-"`
	Username                  string `form:"username"  search:"type:exact;column:username;table:information_user" comment:"用户名"`
	DeviceType                int64  `form:"deviceType"  search:"type:exact;column:device_type;table:information_user" comment:"设备类型：1-iOS，2-Android，3-Web"`
	WalletAddress             string `form:"walletAddress"  search:"type:exact;column:wallet_address;table:information_user" comment:"绑定的钱包地址"`
	PointNotificationSettings int64  `form:"pointNotificationSettings"  search:"type:exact;column:point_notification_settings;table:information_user" comment:"点位通知设置：2-全部，1-部分"`
	Status                    int64  `form:"status"  search:"type:exact;column:status;table:information_user" comment:"状态：3-禁用，1-正常，2-锁定"`
	InformationUserOrder
}

type InformationUserOrder struct {
	Id                        string `form:"idOrder"  search:"type:order;column:id;table:information_user"`
	Username                  string `form:"usernameOrder"  search:"type:order;column:username;table:information_user"`
	Head                      string `form:"headOrder"  search:"type:order;column:head;table:information_user"`
	Email                     string `form:"emailOrder"  search:"type:order;column:email;table:information_user"`
	Password                  string `form:"passwordOrder"  search:"type:order;column:password;table:information_user"`
	DeviceId                  string `form:"deviceIdOrder"  search:"type:order;column:device_id;table:information_user"`
	DeviceType                string `form:"deviceTypeOrder"  search:"type:order;column:device_type;table:information_user"`
	SystemVersion             string `form:"systemVersionOrder"  search:"type:order;column:system_version;table:information_user"`
	AppVersion                string `form:"appVersionOrder"  search:"type:order;column:app_version;table:information_user"`
	BundleId                  string `form:"bundleIdOrder"  search:"type:order;column:bundle_id;table:information_user"`
	WalletAddress             string `form:"walletAddressOrder"  search:"type:order;column:wallet_address;table:information_user"`
	Language                  string `form:"languageOrder"  search:"type:order;column:language;table:information_user"`
	Timezone                  string `form:"timezoneOrder"  search:"type:order;column:timezone;table:information_user"`
	TwoFactorEnabled          string `form:"twoFactorEnabledOrder"  search:"type:order;column:two_factor_enabled;table:information_user"`
	TwoFactorSecret           string `form:"twoFactorSecretOrder"  search:"type:order;column:two_factor_secret;table:information_user"`
	PointNotificationSettings string `form:"pointNotificationSettingsOrder"  search:"type:order;column:point_notification_settings;table:information_user"`
	LastLoginAt               string `form:"lastLoginAtOrder"  search:"type:order;column:last_login_at;table:information_user"`
	LastLoginIp               string `form:"lastLoginIpOrder"  search:"type:order;column:last_login_ip;table:information_user"`
	Status                    string `form:"statusOrder"  search:"type:order;column:status;table:information_user"`
	CreatedAt                 string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_user"`
	UpdatedAt                 string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_user"`
	DeletedAt                 string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_user"`
}

func (m *InformationUserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationUserInsertReq struct {
	Id                        int       `json:"-" comment:"主键编码"` // 主键编码
	Username                  string    `json:"username" comment:"用户名"`
	Head                      string    `json:"head" comment:"头像"`
	Email                     string    `json:"email" comment:"邮箱"`
	Password                  string    `json:"password" comment:"密码"`
	DeviceId                  string    `json:"deviceId" comment:"设备ID"`
	DeviceType                string    `json:"deviceType" comment:"设备类型：1-iOS，2-Android，3-Web"`
	SystemVersion             string    `json:"systemVersion" comment:"系统版本 (如 16.0, 13.0)"`
	AppVersion                string    `json:"appVersion" comment:"App版本号 (如 1.2.3)"`
	BundleId                  string    `json:"bundleId" comment:"App包名 (如 com.wallet.dex)"`
	WalletAddress             string    `json:"walletAddress" comment:"绑定的钱包地址"`
	Language                  string    `json:"language" comment:"语言：zh-CN（中文），en-US（英文）"`
	Timezone                  string    `json:"timezone" comment:"时区"`
	TwoFactorEnabled          string    `json:"twoFactorEnabled" comment:"是否启用两步验证：2-否，1-是"`
	TwoFactorSecret           string    `json:"twoFactorSecret" comment:"两步验证密钥"`
	PointNotificationSettings string    `json:"pointNotificationSettings" comment:"点位通知设置：2-全部，1-部分"`
	LastLoginAt               time.Time `json:"lastLoginAt" comment:"最后登录时间"`
	LastLoginIp               string    `json:"lastLoginIp" comment:"最后登录IP"`
	Status                    string    `json:"status" comment:"状态：3-禁用，1-正常，2-锁定"`
	common.ControlBy
}

func (s *InformationUserInsertReq) Generate(model *models.InformationUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.DeviceType != "" {
		t, _ := strconv.Atoi(s.DeviceType)
		model.DeviceType = int64(t)
	}
	if s.TwoFactorEnabled != "" {
		model.TwoFactorEnabled, _ = strconv.Atoi(s.TwoFactorEnabled)
	}
	if s.PointNotificationSettings != "" {
		pointNotificationSettings, _ := strconv.Atoi(s.PointNotificationSettings)
		model.PointNotificationSettings = int64(pointNotificationSettings)
	}
	model.Username = s.Username
	model.Head = s.Head
	model.Email = s.Email
	model.Password = s.Password
	model.DeviceId = s.DeviceId
	model.SystemVersion = s.SystemVersion
	model.AppVersion = s.AppVersion
	model.BundleId = s.BundleId
	model.WalletAddress = s.WalletAddress
	model.Language = s.Language
	model.Timezone = s.Timezone
	model.TwoFactorSecret = s.TwoFactorSecret
	model.LastLoginAt = s.LastLoginAt
	model.LastLoginIp = s.LastLoginIp
}

func (s *InformationUserInsertReq) GetId() interface{} {
	return s.Id
}

type InformationUserUpdateReq struct {
	Id                        int       `uri:"id" comment:"主键编码"` // 主键编码
	Username                  string    `json:"username" comment:"用户名"`
	Head                      string    `json:"head" comment:"头像"`
	Email                     string    `json:"email" comment:"邮箱"`
	Password                  string    `json:"password" comment:"密码"`
	DeviceId                  string    `json:"deviceId" comment:"设备ID"`
	DeviceType                string    `json:"deviceType" comment:"设备类型：1-iOS，2-Android，3-Web"`
	SystemVersion             string    `json:"systemVersion" comment:"系统版本 (如 16.0, 13.0)"`
	AppVersion                string    `json:"appVersion" comment:"App版本号 (如 1.2.3)"`
	BundleId                  string    `json:"bundleId" comment:"App包名 (如 com.wallet.dex)"`
	WalletAddress             string    `json:"walletAddress" comment:"绑定的钱包地址"`
	Language                  string    `json:"language" comment:"语言：zh-CN（中文），en-US（英文）"`
	Timezone                  string    `json:"timezone" comment:"时区"`
	TwoFactorEnabled          string    `json:"twoFactorEnabled" comment:"是否启用两步验证：2-否，1-是"`
	TwoFactorSecret           string    `json:"twoFactorSecret" comment:"两步验证密钥"`
	PointNotificationSettings string    `json:"pointNotificationSettings" comment:"点位通知设置：2-全部，1-部分"`
	LastLoginAt               time.Time `json:"lastLoginAt" comment:"最后登录时间"`
	LastLoginIp               string    `json:"lastLoginIp" comment:"最后登录IP"`
	Status                    string    `json:"status" comment:"状态：3-禁用，1-正常，2-锁定"`
	common.ControlBy
}

func (s *InformationUserUpdateReq) Generate(model *models.InformationUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.DeviceType != "" {
		t, _ := strconv.Atoi(s.DeviceType)
		model.DeviceType = int64(t)
	}
	if s.TwoFactorEnabled != "" {
		model.TwoFactorEnabled, _ = strconv.Atoi(s.TwoFactorEnabled)
	}
	if s.PointNotificationSettings != "" {
		pointNotificationSettings, _ := strconv.Atoi(s.PointNotificationSettings)
		model.PointNotificationSettings = int64(pointNotificationSettings)
	}
	model.Username = s.Username
	model.Head = s.Head
	model.Email = s.Email
	model.Password = s.Password
	model.DeviceId = s.DeviceId
	model.SystemVersion = s.SystemVersion
	model.AppVersion = s.AppVersion
	model.BundleId = s.BundleId
	model.WalletAddress = s.WalletAddress
	model.Language = s.Language
	model.Timezone = s.Timezone
	model.TwoFactorSecret = s.TwoFactorSecret
	model.LastLoginAt = s.LastLoginAt
	model.LastLoginIp = s.LastLoginIp
}

func (s *InformationUserUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationUserGetReq 功能获取请求参数
type InformationUserGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationUserGetReq) GetId() interface{} {
	return s.Id
}

// InformationUserDeleteReq 功能删除请求参数
type InformationUserDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationUserDeleteReq) GetId() interface{} {
	return s.Ids
}
