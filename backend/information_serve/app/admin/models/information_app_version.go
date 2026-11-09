package models

import (
	"go-admin/common/models"
)

type InformationAppVersion struct {
	models.Model

	Version             string `json:"version" gorm:"type:varchar(32);comment:版本号"`
	BuildNumber         string `json:"buildNumber" gorm:"type:varchar(32);comment:构建号"`
	Platform            int64  `json:"platform" gorm:"type:tinyint;comment:平台：1-iOS，2-Android"`
	DownloadUrl         string `json:"downloadUrl" gorm:"type:varchar(256);comment:下载地址"`
	ForceUpdate         int64  `json:"forceUpdate" gorm:"type:tinyint;comment:是否强制更新：2-否，1-是"`
	Content             string `json:"content" gorm:"type:text;comment:当前更新内容"`
	EnContent           string `json:"enContent" gorm:"type:text;comment:英文的当前更新内容"`
	UpdateType          int64  `json:"updateType" gorm:"type:tinyint;comment:类型 1-功能更新 2-问题修复 3-性能优化"`
	Channel             string `json:"channel" gorm:"type:varchar(32);comment:发布渠道"`
	Status              int64  `json:"status" gorm:"type:tinyint;comment:状态 1-发布 2-草稿 3-撤回"`
	File                string `json:"file" gorm:"type:varchar(512);comment:文件"`
	VersionHistoryLog   string `json:"versionHistoryLog" gorm:"type:text;comment:版本历史日志"`
	EnVersionHistoryLog string `json:"EnVersionHistoryLog" gorm:"type:text;comment:英文的版本历史日志"`
	models.ModelTime
	models.ControlBy
}

func (InformationAppVersion) TableName() string {
	return "information_app_version"
}

func (e *InformationAppVersion) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationAppVersion) GetId() interface{} {
	return e.Id
}
