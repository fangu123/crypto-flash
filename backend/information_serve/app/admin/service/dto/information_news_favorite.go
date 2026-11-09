package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InformationNewsFavoriteGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId int64 `form:"userId"  search:"type:exact;column:user_id;table:information_news_favorite" comment:"用户ID"`
    NewsId int64 `form:"newsId"  search:"type:exact;column:news_id;table:information_news_favorite" comment:"快讯ID"`
    InformationNewsFavoriteOrder
}

type InformationNewsFavoriteOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:information_news_favorite"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:information_news_favorite"`
    NewsId string `form:"newsIdOrder"  search:"type:order;column:news_id;table:information_news_favorite"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_news_favorite"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_news_favorite"`
    
}

func (m *InformationNewsFavoriteGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationNewsFavoriteInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    UserId int64 `json:"userId" comment:"用户ID"`
    NewsId int64 `json:"newsId" comment:"快讯ID"`
    common.ControlBy
}

func (s *InformationNewsFavoriteInsertReq) Generate(model *models.InformationNewsFavorite)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.NewsId = s.NewsId
}

func (s *InformationNewsFavoriteInsertReq) GetId() interface{} {
	return s.Id
}

type InformationNewsFavoriteUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    UserId int64 `json:"userId" comment:"用户ID"`
    NewsId int64 `json:"newsId" comment:"快讯ID"`
    common.ControlBy
}

func (s *InformationNewsFavoriteUpdateReq) Generate(model *models.InformationNewsFavorite)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.NewsId = s.NewsId
}

func (s *InformationNewsFavoriteUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationNewsFavoriteGetReq 功能获取请求参数
type InformationNewsFavoriteGetReq struct {
     Id int `uri:"id"`
}
func (s *InformationNewsFavoriteGetReq) GetId() interface{} {
	return s.Id
}

// InformationNewsFavoriteDeleteReq 功能删除请求参数
type InformationNewsFavoriteDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationNewsFavoriteDeleteReq) GetId() interface{} {
	return s.Ids
}
