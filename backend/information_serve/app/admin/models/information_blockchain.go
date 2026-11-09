package models

import (

	"go-admin/common/models"

)

type InformationBlockchain struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(64);comment:链名称"` 
    Symbol string `json:"symbol" gorm:"type:varchar(16);comment:链标识"` 
    LogoUrl string `json:"logoUrl" gorm:"type:varchar(255);comment:链Logo URL"` 
    Description string `json:"description" gorm:"type:text;comment:链描述"` 
    Website string `json:"website" gorm:"type:varchar(256);comment:官方网站"` 
    Whitepaper string `json:"whitepaper" gorm:"type:varchar(256);comment:白皮书链接"` 
    Status int64 `json:"status" gorm:"type:tinyint;comment:状态：2-禁用，1-启用"` 
    SortOrder int64 `json:"sortOrder" gorm:"type:int;comment:排序"` 
    models.ModelTime
    models.ControlBy
}

func (InformationBlockchain) TableName() string {
    return "information_blockchain"
}

func (e *InformationBlockchain) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationBlockchain) GetId() interface{} {
	return e.Id
}