package dto

import (
	"strconv"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InformationTradingPointGetPageReq struct {
	dto.Pagination `search:"-"`
	CoinId         int64 `form:"coinId"  search:"type:exact;column:coin_id;table:information_trading_point" comment:"币种ID"`
	Type           int64 `form:"type"  search:"type:exact;column:type;table:information_trading_point" comment:"类型：1-做多，2-做空"`
	Status         int64 `form:"status"  search:"type:exact;column:status;table:information_trading_point" comment:"状态：1-进行中，2-已失效，3-已撤单"`
	IsNotified     int64 `form:"isNotified"  search:"type:exact;column:is_notified;table:information_trading_point" comment:"是否已通知：1-未通知，2-已通知"`
	InformationTradingPointOrder
}

type InformationTradingPointOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:information_trading_point"`
	CoinId        string `form:"coinIdOrder"  search:"type:order;column:coin_id;table:information_trading_point"`
	Type          string `form:"typeOrder"  search:"type:order;column:type;table:information_trading_point"`
	EntryPrice    string `form:"entryPriceOrder"  search:"type:order;column:entry_price;table:information_trading_point"`
	TakeProfits   string `form:"takeProfitsOrder"  search:"type:order;column:take_profits;table:information_trading_point"`
	StopLoss      string `form:"stopLossOrder"  search:"type:order;column:stop_loss;table:information_trading_point"`
	ValidDuration string `form:"validDurationOrder"  search:"type:order;column:valid_duration;table:information_trading_point"`
	ValidUntil    string `form:"validUntilOrder"  search:"type:order;column:valid_until;table:information_trading_point"`
	Status        string `form:"statusOrder"  search:"type:order;column:status;table:information_trading_point"`
	IsNotified    string `form:"isNotifiedOrder"  search:"type:order;column:is_notified;table:information_trading_point"`
	Notes         string `form:"notesOrder"  search:"type:order;column:notes;table:information_trading_point"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_trading_point"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_trading_point"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_trading_point"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:information_trading_point"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_trading_point"`
}

func (m *InformationTradingPointGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationTradingPointInsertReq struct {
	Id            int       `json:"-" comment:"主键编码"` // 主键编码
	CoinId        string    `json:"coinId" comment:"币种ID"`
	Type          string    `json:"type" comment:"类型：1-做多，2-做空"`
	EntryPrice    string    `json:"entryPrice" comment:"入场价格"`
	TakeProfits   string    `json:"takeProfits" comment:"止盈价格数组"`
	StopLoss      string    `json:"stopLoss" comment:"止损价格"`
	ValidDuration int64     `json:"validDuration" comment:"有效时长：6小时，12小时，24小时，48小时"`
	ValidUntil    time.Time `json:"validUntil" comment:"到期时间"`
	Status        string    `json:"status" comment:"状态：1-进行中，2-已失效，3-已撤单"`
	IsNotified    string    `json:"isNotified" comment:"是否已通知：1-未通知，2-已通知"`
	Notes         string    `json:"notes" comment:"备注"`
	common.ControlBy
}

func (s *InformationTradingPointInsertReq) Generate(model *models.InformationTradingPoint) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.Type != "" {
		t, _ := strconv.Atoi(s.Type)
		model.Type = int64(t)
	}
	if s.IsNotified != "" {
		isNotified, _ := strconv.Atoi(s.IsNotified)
		model.IsNotified = int64(isNotified)
	}
	if s.CoinId != "" {
		relatedCoinId, _ := strconv.Atoi(s.CoinId)
		model.CoinId = int64(relatedCoinId)
	}
	model.EntryPrice = s.EntryPrice
	model.TakeProfits = s.TakeProfits
	model.StopLoss = s.StopLoss
	model.ValidDuration = s.ValidDuration
	model.ValidUntil = s.ValidUntil
	model.Notes = s.Notes
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationTradingPointInsertReq) GetId() interface{} {
	return s.Id
}

type InformationTradingPointUpdateReq struct {
	Id            int       `uri:"id" comment:"主键编码"` // 主键编码
	CoinId        string    `json:"coinId" comment:"币种ID"`
	Type          string    `json:"type" comment:"类型：1-做多，2-做空"`
	EntryPrice    string    `json:"entryPrice" comment:"入场价格"`
	TakeProfits   string    `json:"takeProfits" comment:"止盈价格数组"`
	StopLoss      string    `json:"stopLoss" comment:"止损价格"`
	ValidDuration int64     `json:"validDuration" comment:"有效时长：6小时，12小时，24小时，48小时"`
	ValidUntil    time.Time `json:"validUntil" comment:"到期时间"`
	Status        string    `json:"status" comment:"状态：1-进行中，2-已失效，3-已撤单"`
	IsNotified    string    `json:"isNotified" comment:"是否已通知：1-未通知，2-已通知"`
	Notes         string    `json:"notes" comment:"备注"`
	common.ControlBy
}

func (s *InformationTradingPointUpdateReq) Generate(model *models.InformationTradingPoint) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.Type != "" {
		t, _ := strconv.Atoi(s.Type)
		model.Type = int64(t)
	}
	if s.IsNotified != "" {
		isNotified, _ := strconv.Atoi(s.IsNotified)
		model.IsNotified = int64(isNotified)
	}
	if s.CoinId != "" {
		relatedCoinId, _ := strconv.Atoi(s.CoinId)
		model.CoinId = int64(relatedCoinId)
	}
	model.EntryPrice = s.EntryPrice
	model.TakeProfits = s.TakeProfits
	model.StopLoss = s.StopLoss
	model.ValidDuration = s.ValidDuration
	model.ValidUntil = s.ValidUntil
	model.Notes = s.Notes
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationTradingPointUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationTradingPointGetReq 功能获取请求参数
type InformationTradingPointGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationTradingPointGetReq) GetId() interface{} {
	return s.Id
}

// InformationTradingPointDeleteReq 功能删除请求参数
type InformationTradingPointDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationTradingPointDeleteReq) GetId() interface{} {
	return s.Ids
}
