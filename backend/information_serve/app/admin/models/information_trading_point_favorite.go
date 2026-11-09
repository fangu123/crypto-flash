package models

import (

	"go-admin/common/models"

)

type InformationTradingPointFavorite struct {
    models.Model
    
    UserId int64 `json:"userId" gorm:"type:bigint;comment:用户ID"` 
    TradingPointId int64 `json:"tradingPointId" gorm:"type:bigint;comment:点位ID"` 
    models.ModelTime
    models.ControlBy
}

func (InformationTradingPointFavorite) TableName() string {
    return "information_trading_point_favorite"
}

func (e *InformationTradingPointFavorite) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationTradingPointFavorite) GetId() interface{} {
	return e.Id
}