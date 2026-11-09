package models

import (
	"go-admin/common/models"
)

type InformationThirdpartyApikey struct {
	models.Model

	Type       int64  `json:"type" gorm:"type:tinyint;comment:类型：1-coingecko，2-bybit，3-odaily，4-blockBeats"`
	BaseUrl    string `json:"baseUrl" gorm:"type:varchar(128);comment:url地址"`
	Apikey     string `json:"apikey" gorm:"type:varchar(128);comment:apikey"`
	Secretkey  string `json:"secretkey" gorm:"type:varchar(128);comment:secretkey"`
	Passphrase string `json:"passphrase" gorm:"type:varchar(128);comment:passphrase"`
	Status     int64  `json:"status" gorm:"type:tinyint;comment:状态 1-启用 2-停用"`
	models.ModelTime
	models.ControlBy
}

func (InformationThirdpartyApikey) TableName() string {
	return "information_thirdparty_apikey"
}

func (e *InformationThirdpartyApikey) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationThirdpartyApikey) GetId() interface{} {
	return e.Id
}
