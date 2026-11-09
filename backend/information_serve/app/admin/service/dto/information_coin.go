package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strconv"

	"github.com/shopspring/decimal"
)

type InformationCoinGetPageReq struct {
	dto.Pagination `search:"-"`
	BlockchainId   int64  `form:"blockchainId"  search:"type:exact;column:blockchain_id;table:information_coin" comment:"所属链ID"`
	Symbol         string `form:"symbol"  search:"type:exact;column:symbol;table:information_coin" comment:"币种标识"`
	Status         int64  `form:"status"  search:"type:exact;column:status;table:information_coin" comment:"状态：2-禁用，1-启用"`
	InformationCoinOrder
}

type InformationCoinOrder struct {
	Id                string `form:"idOrder"  search:"type:order;column:id;table:information_coin"`
	BlockchainId      string `form:"blockchainIdOrder"  search:"type:order;column:blockchain_id;table:information_coin"`
	Name              string `form:"nameOrder"  search:"type:order;column:name;table:information_coin"`
	Symbol            string `form:"symbolOrder"  search:"type:order;column:symbol;table:information_coin"`
	LogoUrl           string `form:"logoUrlOrder"  search:"type:order;column:logo_url;table:information_coin"`
	Price             string `form:"priceOrder"  search:"type:order;column:price;table:information_coin"`
	PriceChange24h    string `form:"priceChange24hOrder"  search:"type:order;column:price_change_24h;table:information_coin"`
	MarketCap         string `form:"marketCapOrder"  search:"type:order;column:market_cap;table:information_coin"`
	Volume24h         string `form:"volume24hOrder"  search:"type:order;column:volume_24h;table:information_coin"`
	CirculatingSupply string `form:"circulatingSupplyOrder"  search:"type:order;column:circulating_supply;table:information_coin"`
	TotalSupply       string `form:"totalSupplyOrder"  search:"type:order;column:total_supply;table:information_coin"`
	MaxSupply         string `form:"maxSupplyOrder"  search:"type:order;column:max_supply;table:information_coin"`
	Rank              string `form:"rankOrder"  search:"type:order;column:rank;table:information_coin"`
	Status            string `form:"statusOrder"  search:"type:order;column:status;table:information_coin"`
	CreatedAt         string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_coin"`
	UpdatedAt         string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_coin"`
	DeletedAt         string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_coin"`
	CreateBy          string `form:"createByOrder"  search:"type:order;column:create_by;table:information_coin"`
	UpdateBy          string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_coin"`
}

func (m *InformationCoinGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationCoinInsertReq struct {
	Id                int             `json:"-" comment:"主键编码"` // 主键编码
	BlockchainId      int64           `json:"blockchainId" comment:"所属链ID"`
	Name              string          `json:"name" comment:"币种名称"`
	Symbol            string          `json:"symbol" comment:"币种标识"`
	LogoUrl           string          `json:"logoUrl" comment:"币Logo URL"`
	Price             decimal.Decimal `json:"price" comment:"当前价格"`
	PriceChange24h    decimal.Decimal `json:"priceChange24h" comment:"24小时价格变化"`
	MarketCap         decimal.Decimal `json:"marketCap" comment:"市值"`
	Volume24h         decimal.Decimal `json:"volume24h" comment:"24小时交易量"`
	CirculatingSupply decimal.Decimal `json:"circulatingSupply" comment:"流通量"`
	TotalSupply       decimal.Decimal `json:"totalSupply" comment:"总量"`
	MaxSupply         decimal.Decimal `json:"maxSupply" comment:"最大量"`
	Rank              int64           `json:"rank" comment:"排名"`
	Status            string          `json:"status" comment:"状态：2-禁用，1-启用"`
	common.ControlBy
}

func (s *InformationCoinInsertReq) Generate(model *models.InformationCoin) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	model.BlockchainId = s.BlockchainId
	model.Name = s.Name
	model.Symbol = s.Symbol
	model.LogoUrl = s.LogoUrl
	model.Price = s.Price
	model.PriceChange24h = s.PriceChange24h
	model.MarketCap = s.MarketCap
	model.Volume24h = s.Volume24h
	model.CirculatingSupply = s.CirculatingSupply
	model.TotalSupply = s.TotalSupply
	model.MaxSupply = s.MaxSupply
	model.Rank = s.Rank
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationCoinInsertReq) GetId() interface{} {
	return s.Id
}

type InformationCoinUpdateReq struct {
	Id                int             `uri:"id" comment:"主键编码"` // 主键编码
	BlockchainId      int64           `json:"blockchainId" comment:"所属链ID"`
	Name              string          `json:"name" comment:"币种名称"`
	Symbol            string          `json:"symbol" comment:"币种标识"`
	LogoUrl           string          `json:"logoUrl" comment:"币Logo URL"`
	Price             decimal.Decimal `json:"price" comment:"当前价格"`
	PriceChange24h    decimal.Decimal `json:"priceChange24h" comment:"24小时价格变化"`
	MarketCap         decimal.Decimal `json:"marketCap" comment:"市值"`
	Volume24h         decimal.Decimal `json:"volume24h" comment:"24小时交易量"`
	CirculatingSupply decimal.Decimal `json:"circulatingSupply" comment:"流通量"`
	TotalSupply       decimal.Decimal `json:"totalSupply" comment:"总量"`
	MaxSupply         decimal.Decimal `json:"maxSupply" comment:"最大量"`
	Rank              int64           `json:"rank" comment:"排名"`
	Status            string          `json:"status" comment:"状态：2-禁用，1-启用"`
	common.ControlBy
}

func (s *InformationCoinUpdateReq) Generate(model *models.InformationCoin) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	model.BlockchainId = s.BlockchainId
	model.Name = s.Name
	model.Symbol = s.Symbol
	model.LogoUrl = s.LogoUrl
	model.Price = s.Price
	model.PriceChange24h = s.PriceChange24h
	model.MarketCap = s.MarketCap
	model.Volume24h = s.Volume24h
	model.CirculatingSupply = s.CirculatingSupply
	model.TotalSupply = s.TotalSupply
	model.MaxSupply = s.MaxSupply
	model.Rank = s.Rank
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationCoinUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationCoinGetReq 功能获取请求参数
type InformationCoinGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationCoinGetReq) GetId() interface{} {
	return s.Id
}

// InformationCoinDeleteReq 功能删除请求参数
type InformationCoinDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationCoinDeleteReq) GetId() interface{} {
	return s.Ids
}
