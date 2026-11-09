package dto

import (
	"strconv"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InformationNotificationUserGetPageReq struct {
	dto.Pagination `search:"-"`
	UserId         int64 `form:"userId"  search:"type:exact;column:user_id;table:information_notification_user" comment:"用户ID"`
	NotificationId int64 `form:"notificationId"  search:"type:exact;column:notification_id;table:information_notification_user" comment:"通知ID，关联 information_notification 表"`
	IsRead         int64 `form:"isRead"  search:"type:exact;column:is_read;table:information_notification_user" comment:"是否已读：2-未读，1-已读"`
	InformationNotificationUserOrder
}

type InformationNotificationUserOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:information_notification_user"`
	UserId         string `form:"userIdOrder"  search:"type:order;column:user_id;table:information_notification_user"`
	NotificationId string `form:"notificationIdOrder"  search:"type:order;column:notification_id;table:information_notification_user"`
	IsRead         string `form:"isReadOrder"  search:"type:order;column:is_read;table:information_notification_user"`
	ReadAt         string `form:"readAtOrder"  search:"type:order;column:read_at;table:information_notification_user"`
	CreatedAt      string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_notification_user"`
	DeletedAt      string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_notification_user"`
	CreateBy       string `form:"createByOrder"  search:"type:order;column:create_by;table:information_notification_user"`
	UpdateBy       string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_notification_user"`
}

func (m *InformationNotificationUserGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationNotificationUserInsertReq struct {
	Id             int        `json:"-" comment:"主键编码"` // 主键编码
	UserId         int64      `json:"userId" comment:"用户ID"`
	NotificationId int64      `json:"notificationId" comment:"通知ID，关联 information_notification 表"`
	IsRead         string     `json:"isRead" comment:"是否已读：2-未读，1-已读"`
	ReadAt         *time.Time `json:"readAt" comment:"阅读时间，用户点击消息时更新"`
	common.ControlBy
}

func (s *InformationNotificationUserInsertReq) Generate(model *models.InformationNotificationUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.IsRead != "" {
		isRead, _ := strconv.Atoi(s.IsRead)
		model.IsRead = int64(isRead)
	}
	model.UserId = s.UserId
	model.NotificationId = s.NotificationId
	model.ReadAt = s.ReadAt
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationNotificationUserInsertReq) GetId() interface{} {
	return s.Id
}

type InformationNotificationUserUpdateReq struct {
	Id             int        `uri:"id" comment:"主键编码"` // 主键编码
	UserId         int64      `json:"userId" comment:"用户ID"`
	NotificationId int64      `json:"notificationId" comment:"通知ID，关联 information_notification 表"`
	IsRead         string     `json:"isRead" comment:"是否已读：2-未读，1-已读"`
	ReadAt         *time.Time `json:"readAt" comment:"阅读时间，用户点击消息时更新"`
	common.ControlBy
}

func (s *InformationNotificationUserUpdateReq) Generate(model *models.InformationNotificationUser) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.IsRead != "" {
		isRead, _ := strconv.Atoi(s.IsRead)
		model.IsRead = int64(isRead)
	}
	model.UserId = s.UserId
	model.NotificationId = s.NotificationId
	model.ReadAt = s.ReadAt
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationNotificationUserUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationNotificationUserGetReq 功能获取请求参数
type InformationNotificationUserGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationNotificationUserGetReq) GetId() interface{} {
	return s.Id
}

// InformationNotificationUserDeleteReq 功能删除请求参数
type InformationNotificationUserDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationNotificationUserDeleteReq) GetId() interface{} {
	return s.Ids
}
