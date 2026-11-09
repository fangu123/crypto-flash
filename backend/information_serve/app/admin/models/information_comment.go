package models

import (
	"go-admin/common/models"
)

type InformationComment struct {
	models.Model

	NewsId        int64  `json:"newsId" gorm:"type:bigint;comment:资讯ID"`
	UserId        int64  `json:"userId" gorm:"type:bigint;comment:评论用户ID"`
	ParentId      int64  `json:"parentId" gorm:"type:bigint;comment:父评论ID，1表示顶级评论"`
	ReplyToUserId int64  `json:"replyToUserId" gorm:"type:bigint;comment:被回复用户ID"`
	Content       string `json:"content" gorm:"type:text;comment:评论内容"`
	LikeCount     int64  `json:"likeCount" gorm:"type:int;comment:点赞数"`
	IsTop         int64  `json:"isTop" gorm:"type:tinyint(1);comment:是否置顶 2否 1是"`
	Status        int64  `json:"status" gorm:"type:tinyint(1);comment:状态 4待审核 1正常 2违规 3已删除"`
	IpLocation    string `json:"ipLocation" gorm:"type:varchar(50);comment:IP归属地"`
	models.ModelTime
	models.ControlBy
	UserName        string `json:"userName" gorm:"-"`
	ParentName      string `json:"parentName" gorm:"-"`
	ReplyToUserName string `json:"replyToUserName" gorm:"-"`
}

func (InformationComment) TableName() string {
	return "information_comment"
}

func (e *InformationComment) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationComment) GetId() interface{} {
	return e.Id
}
