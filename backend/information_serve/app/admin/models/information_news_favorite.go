package models

import (

	"go-admin/common/models"

)

type InformationNewsFavorite struct {
    models.Model
    
    UserId int64 `json:"userId" gorm:"type:bigint;comment:用户ID"` 
    NewsId int64 `json:"newsId" gorm:"type:bigint;comment:快讯ID"` 
    models.ModelTime
    models.ControlBy
}

func (InformationNewsFavorite) TableName() string {
    return "information_news_favorite"
}

func (e *InformationNewsFavorite) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationNewsFavorite) GetId() interface{} {
	return e.Id
}