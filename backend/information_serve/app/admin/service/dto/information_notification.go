package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strconv"
)

type InformationNotificationGetPageReq struct {
	dto.Pagination `search:"-"`
	Type           int64  `form:"type"  search:"type:exact;column:type;table:information_notification" comment:"类型：1-系统通知，2-点位提醒，3-快讯推送"`
	Title          string `form:"title"  search:"type:exact;column:title;table:information_notification" comment:"标题"`
	InformationNotificationOrder
}

type InformationNotificationOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:information_notification"`
	Type      string `form:"typeOrder"  search:"type:order;column:type;table:information_notification"`
	Title     string `form:"titleOrder"  search:"type:order;column:title;table:information_notification"`
	Content   string `form:"contentOrder"  search:"type:order;column:content;table:information_notification"`
	Link      string `form:"linkOrder"  search:"type:order;column:link;table:information_notification"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_notification"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_notification"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:information_notification"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_notification"`
}

func (m *InformationNotificationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationNotificationInsertReq struct {
	Id        int    `json:"-" comment:"主键编码"` // 主键编码
	Type      string `json:"type" comment:"类型：1-系统通知，2-点位提醒，3-快讯推送"`
	Title     string `json:"title" comment:"标题"`
	EnTitle   string `json:"enTitle" comment:"英文的标题"`
	Content   string `json:"content" comment:"内容"`
	EnContent string `json:"enContent" comment:"英文的内容"`
	Link      string `json:"link" comment:"跳转链接"`
	common.ControlBy
}

func (s *InformationNotificationInsertReq) Generate(model *models.InformationNotification) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Type != "" {
		t, _ := strconv.Atoi(s.Type)
		model.Type = int64(t)
	}
	model.Title = s.Title
	model.EnTitle = s.EnTitle
	model.Content = s.Content
	model.EnContent = s.EnContent
	model.Link = s.Link
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationNotificationInsertReq) GetId() interface{} {
	return s.Id
}

type InformationNotificationUpdateReq struct {
	Id        int    `uri:"id" comment:"主键编码"` // 主键编码
	Type      string `json:"type" comment:"类型：1-系统通知，2-点位提醒，3-快讯推送"`
	Title     string `json:"title" comment:"标题"`
	EnTitle   string `json:"enTitle" comment:"英文的标题"`
	Content   string `json:"content" comment:"内容"`
	EnContent string `json:"enContent" comment:"英文的内容"`
	Link      string `json:"link" comment:"跳转链接"`
	common.ControlBy
}

func (s *InformationNotificationUpdateReq) Generate(model *models.InformationNotification) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Type != "" {
		t, _ := strconv.Atoi(s.Type)
		model.Type = int64(t)
	}
	model.Title = s.Title
	model.EnTitle = s.EnTitle
	model.Content = s.Content
	model.EnContent = s.EnContent
	model.Link = s.Link
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationNotificationUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationNotificationGetReq 功能获取请求参数
type InformationNotificationGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationNotificationGetReq) GetId() interface{} {
	return s.Id
}

// InformationNotificationDeleteReq 功能删除请求参数
type InformationNotificationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationNotificationDeleteReq) GetId() interface{} {
	return s.Ids
}
