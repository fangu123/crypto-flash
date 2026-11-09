package models

import (
	"time"

	"go-admin/common/models"
)

type InformationNotificationUser struct {
	models.Model

	UserId         int64      `json:"userId" gorm:"type:bigint;comment:用户ID"`
	NotificationId int64      `json:"notificationId" gorm:"type:bigint;comment:通知ID，关联 information_notification 表"`
	IsRead         int64      `json:"isRead" gorm:"type:tinyint;comment:是否已读：2-未读，1-已读"`
	ReadAt         *time.Time `json:"readAt" gorm:"type:datetime(3);comment:阅读时间，用户点击消息时更新"`
	models.ModelTime
	models.ControlBy
}

func (InformationNotificationUser) TableName() string {
	return "information_notification_user"
}

func (e *InformationNotificationUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationNotificationUser) GetId() interface{} {
	return e.Id
}
