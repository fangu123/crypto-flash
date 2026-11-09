package models

import (
	"go-admin/common/models"
)

type InformationContentTag struct {
	models.Model

	Name   string `json:"name" gorm:"type:varchar(64);comment:标签名称"`
	EnName string `json:"enName" gorm:"type:varchar(64);comment:英文的标签名称"`
	Code   string `json:"code" gorm:"type:varchar(32);comment:标签编码"`
	Status int64  `json:"status" gorm:"type:tinyint;comment:状态：2-禁用，1-启用"`
	models.ModelTime
	models.ControlBy
}

func (InformationContentTag) TableName() string {
	return "information_content_tag"
}

func (e *InformationContentTag) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationContentTag) GetId() interface{} {
	return e.Id
}
