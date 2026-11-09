package models

import (
	"time"

	"go-admin/common/models"
)

type InformationNews struct {
	models.Model

	Title                 string    `json:"title" gorm:"type:varchar(512);comment:标题"`
	EnTitle               string    `json:"enTitle" gorm:"type:varchar(512);comment:英文的标题"`
	Content               string    `json:"content" gorm:"type:text;comment:内容"`
	EnContent             string    `json:"enContent" gorm:"type:text;comment:英文的内容"`
	Summary               string    `json:"summary" gorm:"type:varchar(512);comment:摘要"`
	CoverImage            string    `json:"coverImage" gorm:"type:varchar(256);comment:封面图片"`
	Source                string    `json:"source" gorm:"type:varchar(64);comment:来源"`
	SourceId              int64     `json:"sourceId" gorm:"type:bigint;comment:来源id"`
	LogoUrl               string    `json:"logoUrl" gorm:"type:varchar(64);comment:logo图标"`
	CategoryId            int64     `json:"categoryId" gorm:"type:bigint;comment:分类ID"`
	RelatedCoinId         int64     `json:"relatedCoinId" gorm:"type:bigint;comment:相关币种id"`
	RelatedTradingPointId int64     `json:"relatedTradingPointId" gorm:"type:bigint;comment:相关点位id"`
	Level                 int64     `json:"level" gorm:"type:tinyint;comment:级别：1-普通，2-重要，3-紧急"`
	Sentiment             int64     `json:"sentiment" gorm:"type:tinyint;comment:情绪：1-中性，2-利好，3-利空"`
	IsFeatured            int64     `json:"isFeatured" gorm:"type:tinyint;comment:是否头条：2-否，1-是"`
	ViewCount             int64     `json:"viewCount" gorm:"type:bigint;comment:浏览次数"`
	LikeCount             int64     `json:"likeCount" gorm:"type:bigint;comment:点赞次数"`
	ShareCount            int64     `json:"shareCount" gorm:"type:bigint;comment:分享次数"`
	Status                int64     `json:"status" gorm:"type:tinyint;comment:状态：1-草稿，2-已发布，3-已下架"`
	PublishedAt           time.Time `json:"publishedAt" gorm:"type:datetime(3);comment:发布时间"`
	models.ModelTime
	models.ControlBy
}

func (InformationNews) TableName() string {
	return "information_news"
}

func (e *InformationNews) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationNews) GetId() interface{} {
	return e.Id
}
