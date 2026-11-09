package models

import (
	"go-admin/common/models"
)

type InformationNewsTagRelation struct {
	models.Model

	NewsId int64 `json:"newsId" gorm:"type:bigint;comment:新闻ID"`
	TagId  int64 `json:"tagId" gorm:"type:bigint;comment:标签ID"`
	models.ModelTime
	models.ControlBy
}

func (InformationNewsTagRelation) TableName() string {
	return "information_news_tag_relation"
}

func (e *InformationNewsTagRelation) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationNewsTagRelation) GetId() interface{} {
	return e.Id
}
