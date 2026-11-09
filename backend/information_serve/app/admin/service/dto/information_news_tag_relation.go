package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InformationNewsTagRelationGetPageReq struct {
	dto.Pagination `search:"-"`
	NewsId         int64 `form:"newsId"  search:"type:exact;column:news_id;table:information_news_tag_relation" comment:"新闻ID"`
	TagId          int64 `form:"tagId"  search:"type:exact;column:tag_id;table:information_news_tag_relation" comment:"标签ID"`
	InformationNewsTagRelationOrder
}

type InformationNewsTagRelationOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:information_news_tag_relation"`
	NewsId    string `form:"newsIdOrder"  search:"type:order;column:news_id;table:information_news_tag_relation"`
	TagId     string `form:"tagIdOrder"  search:"type:order;column:tag_id;table:information_news_tag_relation"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_news_tag_relation"`
}

func (m *InformationNewsTagRelationGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationNewsTagRelationInsertReq struct {
	Id     int   `json:"-" comment:"主键编码"` // 主键编码
	NewsId int64 `json:"newsId" comment:"新闻ID"`
	TagId  int64 `json:"tagId" comment:"标签ID"`
	common.ControlBy
}

func (s *InformationNewsTagRelationInsertReq) Generate(model *models.InformationNewsTagRelation) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.NewsId = s.NewsId
	model.TagId = s.TagId
}

func (s *InformationNewsTagRelationInsertReq) GetId() interface{} {
	return s.Id
}

type InformationNewsTagRelationUpdateReq struct {
	Id        int   `uri:"id" comment:"主键编码"` // 主键编码
	NewsId    int64 `json:"newsId" comment:"新闻ID"`
	TagId     int64 `json:"tagId" comment:"标签ID"`
	CreatedBy int64 `json:"createdBy" comment:"创建者"`
	common.ControlBy
}

func (s *InformationNewsTagRelationUpdateReq) Generate(model *models.InformationNewsTagRelation) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.NewsId = s.NewsId
	model.TagId = s.TagId
}

func (s *InformationNewsTagRelationUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationNewsTagRelationGetReq 功能获取请求参数
type InformationNewsTagRelationGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationNewsTagRelationGetReq) GetId() interface{} {
	return s.Id
}

// InformationNewsTagRelationDeleteReq 功能删除请求参数
type InformationNewsTagRelationDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationNewsTagRelationDeleteReq) GetId() interface{} {
	return s.Ids
}
