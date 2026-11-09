package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InformationTradingPointFavoriteGetPageReq struct {
	dto.Pagination     `search:"-"`
    UserId int64 `form:"userId"  search:"type:exact;column:user_id;table:information_trading_point_favorite" comment:"用户ID"`
    TradingPointId int64 `form:"tradingPointId"  search:"type:exact;column:trading_point_id;table:information_trading_point_favorite" comment:"点位ID"`
    InformationTradingPointFavoriteOrder
}

type InformationTradingPointFavoriteOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:information_trading_point_favorite"`
    UserId string `form:"userIdOrder"  search:"type:order;column:user_id;table:information_trading_point_favorite"`
    TradingPointId string `form:"tradingPointIdOrder"  search:"type:order;column:trading_point_id;table:information_trading_point_favorite"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_trading_point_favorite"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_trading_point_favorite"`
    
}

func (m *InformationTradingPointFavoriteGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationTradingPointFavoriteInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    UserId int64 `json:"userId" comment:"用户ID"`
    TradingPointId int64 `json:"tradingPointId" comment:"点位ID"`
    common.ControlBy
}

func (s *InformationTradingPointFavoriteInsertReq) Generate(model *models.InformationTradingPointFavorite)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.TradingPointId = s.TradingPointId
}

func (s *InformationTradingPointFavoriteInsertReq) GetId() interface{} {
	return s.Id
}

type InformationTradingPointFavoriteUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    UserId int64 `json:"userId" comment:"用户ID"`
    TradingPointId int64 `json:"tradingPointId" comment:"点位ID"`
    common.ControlBy
}

func (s *InformationTradingPointFavoriteUpdateReq) Generate(model *models.InformationTradingPointFavorite)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.UserId = s.UserId
    model.TradingPointId = s.TradingPointId
}

func (s *InformationTradingPointFavoriteUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationTradingPointFavoriteGetReq 功能获取请求参数
type InformationTradingPointFavoriteGetReq struct {
     Id int `uri:"id"`
}
func (s *InformationTradingPointFavoriteGetReq) GetId() interface{} {
	return s.Id
}

// InformationTradingPointFavoriteDeleteReq 功能删除请求参数
type InformationTradingPointFavoriteDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationTradingPointFavoriteDeleteReq) GetId() interface{} {
	return s.Ids
}
