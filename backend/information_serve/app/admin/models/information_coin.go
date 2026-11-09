package models

import (
	"go-admin/common/models"

	"github.com/shopspring/decimal"
)

type InformationCoin struct {
	models.Model

	BlockchainId      int64           `json:"blockchainId" gorm:"type:bigint;comment:所属链ID"`
	Name              string          `json:"name" gorm:"type:varchar(64);comment:币种名称"`
	Symbol            string          `json:"symbol" gorm:"type:varchar(16);comment:币种标识"`
	LogoUrl           string          `json:"logoUrl" gorm:"type:varchar(255);comment:币Logo URL"`
	Price             decimal.Decimal `json:"price" gorm:"type:decimal(36,18);comment:当前价格"`
	PriceChange24h    decimal.Decimal `json:"priceChange24h" gorm:"type:decimal(16,4);comment:24小时价格变化"`
	MarketCap         decimal.Decimal `json:"marketCap" gorm:"type:decimal(36,18);comment:市值"`
	Volume24h         decimal.Decimal `json:"volume24h" gorm:"type:decimal(36,18);comment:24小时交易量"`
	CirculatingSupply decimal.Decimal `json:"circulatingSupply" gorm:"type:decimal(36,18);comment:流通量"`
	TotalSupply       decimal.Decimal `json:"totalSupply" gorm:"type:decimal(36,18);comment:总量"`
	MaxSupply         decimal.Decimal `json:"maxSupply" gorm:"type:decimal(36,18);comment:最大量"`
	Rank              int64           `json:"rank" gorm:"type:int;comment:排名"`
	Status            int64           `json:"status" gorm:"type:tinyint;comment:状态：2-禁用，1-启用"`
	models.ModelTime
	models.ControlBy
}

func (InformationCoin) TableName() string {
	return "information_coin"
}

func (e *InformationCoin) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationCoin) GetId() interface{} {
	return e.Id
}
