package models

import (
	"time"

	"go-admin/common/models"
)

type InformationTradingPoint struct {
	models.Model

	CoinId        int64     `json:"coinId" gorm:"type:bigint;comment:币种ID"`
	Type          int64     `json:"type" gorm:"type:tinyint;comment:类型：1-做多，2-做空"`
	EntryPrice    string    `json:"entryPrice" gorm:"type:varchar(512);comment:入场价格"`
	TakeProfits   string    `json:"takeProfits" gorm:"type:json;comment:止盈价格数组"`
	StopLoss      string    `json:"stopLoss" gorm:"type:varchar(512);comment:止损价格"`
	ValidDuration int64     `json:"validDuration" gorm:"type:tinyint;comment:有效时长：6小时，12小时，24小时，48小时"`
	ValidUntil    time.Time `json:"validUntil" gorm:"type:datetime(3);comment:到期时间"`
	Status        int64     `json:"status" gorm:"type:tinyint;comment:状态：1-进行中，2-已失效，3-已撤单"`
	IsNotified    int64     `json:"isNotified" gorm:"type:tinyint;comment:是否已通知：1-未通知，2-已通知"`
	Notes         string    `json:"notes" gorm:"type:varchar(512);comment:备注"`
	//不存数据库的字段
	Name string `json:"name" gorm:"-"`
	models.ModelTime
	models.ControlBy
}

func (InformationTradingPoint) TableName() string {
	return "information_trading_point"
}

func (e *InformationTradingPoint) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationTradingPoint) GetId() interface{} {
	return e.Id
}
