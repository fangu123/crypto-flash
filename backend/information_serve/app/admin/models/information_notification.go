package models

import (
	"go-admin/common/models"
)

type InformationNotification struct {
	models.Model

	Type      int64  `json:"type" gorm:"type:tinyint;comment:类型：1-系统通知，2-点位提醒，3-快讯推送"`
	Title     string `json:"title" gorm:"type:varchar(128);comment:标题"`
	EnTitle   string `json:"enTitle" gorm:"type:varchar(128);comment:英文的标题"`
	Content   string `json:"content" gorm:"type:text;comment:内容"`
	EnContent string `json:"enContent" gorm:"type:text;comment:英文的内容"`
	Link      string `json:"link" gorm:"type:varchar(512);comment:跳转链接"`
	models.ModelTime
	models.ControlBy
}

func (InformationNotification) TableName() string {
	return "information_notification"
}

func (e *InformationNotification) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationNotification) GetId() interface{} {
	return e.Id
}
