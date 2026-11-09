package dto

import (
	"strconv"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InformationNewsGetPageReq struct {
	dto.Pagination        `search:"-"`
	Title                 string `form:"title"  search:"type:contains;column:title;table:information_news" comment:"标题"`
	RelatedCoinId         int64  `form:"relatedCoinId"  search:"type:exact;column:related_coin_id;table:information_news" comment:"相关币种id"`
	RelatedTradingPointId int64  `form:"relatedTradingPointId"  search:"type:exact;column:related_trading_point_id;table:information_news" comment:"相关点位id"`
	Level                 int64  `form:"level"  search:"type:exact;column:level;table:information_news" comment:"级别：1-普通，2-重要，3-紧急"`
	Sentiment             int64  `form:"sentiment"  search:"type:exact;column:sentiment;table:information_news" comment:"情绪：1-中性，2-利好，3-利空"`
	IsFeatured            int64  `form:"isFeatured"  search:"type:exact;column:is_featured;table:information_news" comment:"是否头条：2-否，1-是"`
	Status                int64  `form:"status"  search:"type:exact;column:status;table:information_news" comment:"状态：1-草稿，2-已发布，3-已下架"`
	InformationNewsOrder
}

type InformationNewsOrder struct {
	Id                    string `form:"idOrder"  search:"type:order;column:id;table:information_news"`
	Title                 string `form:"titleOrder"  search:"type:order;column:title;table:information_news"`
	Content               string `form:"contentOrder"  search:"type:order;column:content;table:information_news"`
	Summary               string `form:"summaryOrder"  search:"type:order;column:summary;table:information_news"`
	CoverImage            string `form:"coverImageOrder"  search:"type:order;column:cover_image;table:information_news"`
	Source                string `form:"sourceOrder"  search:"type:order;column:source;table:information_news"`
	SourceId              string `form:"sourceIdOrder"  search:"type:order;column:source_id;table:information_news"`
	LogUrl                string `form:"logUrlOrder"  search:"type:order;column:log_url;table:information_news"`
	CategoryId            string `form:"categoryIdOrder"  search:"type:order;column:category_id;table:information_news"`
	RelatedCoinId         string `form:"relatedCoinIdOrder"  search:"type:order;column:related_coin_id;table:information_news"`
	RelatedTradingPointId string `form:"relatedTradingPointIdOrder"  search:"type:order;column:related_trading_point_id;table:information_news"`
	Level                 string `form:"levelOrder"  search:"type:order;column:level;table:information_news"`
	Sentiment             string `form:"sentimentOrder"  search:"type:order;column:sentiment;table:information_news"`
	IsFeatured            string `form:"isFeaturedOrder"  search:"type:order;column:is_featured;table:information_news"`
	ViewCount             string `form:"viewCountOrder"  search:"type:order;column:view_count;table:information_news"`
	LikeCount             string `form:"likeCountOrder"  search:"type:order;column:like_count;table:information_news"`
	ShareCount            string `form:"shareCountOrder"  search:"type:order;column:share_count;table:information_news"`
	Status                string `form:"statusOrder"  search:"type:order;column:status;table:information_news"`
	PublishedAt           string `form:"publishedAtOrder"  search:"type:order;column:published_at;table:information_news"`
	CreatedAt             string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_news"`
	UpdatedAt             string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_news"`
	DeletedAt             string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_news"`
	CreateBy              string `form:"createByOrder"  search:"type:order;column:create_by;table:information_news"`
	UpdateBy              string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_news"`
}

func (m *InformationNewsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationNewsInsertReq struct {
	Id                    int       `json:"-" comment:"主键编码"` // 主键编码
	Title                 string    `json:"title" comment:"标题"`
	EnTitle               string    `json:"enTitle" comment:"英文的标题"`
	Content               string    `json:"content" comment:"内容"`
	EnContent             string    `json:"enContent" comment:"英文的内容"`
	Summary               string    `json:"summary" comment:"摘要"`
	CoverImage            string    `json:"coverImage" comment:"封面图片"`
	Source                string    `json:"source" comment:"来源"`
	SourceId              int64     `json:"sourceId" comment:"来源id"`
	LogUrl                string    `json:"logUrl" comment:"logo图标"`
	CategoryId            int64     `json:"categoryId" comment:"分类ID"`
	RelatedCoinId         string    `json:"relatedCoinId" comment:"相关币种id"`
	RelatedTradingPointId string    `json:"relatedTradingPointId" comment:"相关点位id"`
	Level                 string    `json:"level" comment:"级别：1-普通，2-重要，3-紧急"`
	Sentiment             string    `json:"sentiment" comment:"情绪：1-中性，2-利好，3-利空"`
	IsFeatured            string    `json:"isFeatured" comment:"是否头条：2-否，1-是"`
	ViewCount             int64     `json:"viewCount" comment:"浏览次数"`
	LikeCount             int64     `json:"likeCount" comment:"点赞次数"`
	ShareCount            int64     `json:"shareCount" comment:"分享次数"`
	Status                string    `json:"status" comment:"状态：1-草稿，2-已发布，3-已下架"`
	PublishedAt           time.Time `json:"publishedAt" comment:"发布时间"`
	common.ControlBy
}

func (s *InformationNewsInsertReq) Generate(model *models.InformationNews) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.Level != "" {
		level, _ := strconv.Atoi(s.Level)
		model.Level = int64(level)
	}
	if s.Sentiment != "" {
		sentiment, _ := strconv.Atoi(s.Sentiment)
		model.Sentiment = int64(sentiment)
	}
	if s.IsFeatured != "" {
		isFeatured, _ := strconv.Atoi(s.IsFeatured)
		model.IsFeatured = int64(isFeatured)
	}
	if s.RelatedCoinId != "" {
		relatedCoinId, _ := strconv.Atoi(s.RelatedCoinId)
		model.RelatedCoinId = int64(relatedCoinId)
	}
	if s.RelatedTradingPointId != "" {
		relatedTradingPointId, _ := strconv.Atoi(s.RelatedTradingPointId)
		model.RelatedTradingPointId = int64(relatedTradingPointId)
	}
	model.Title = s.Title
	model.EnTitle = s.EnTitle
	model.Content = s.Content
	model.EnContent = s.EnContent
	model.Summary = s.Summary
	model.CoverImage = s.CoverImage
	model.Source = s.Source
	model.SourceId = s.SourceId
	model.LogoUrl = s.LogUrl
	model.CategoryId = s.CategoryId
	model.ViewCount = s.ViewCount
	model.LikeCount = s.LikeCount
	model.ShareCount = s.ShareCount
	model.PublishedAt = s.PublishedAt
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationNewsInsertReq) GetId() interface{} {
	return s.Id
}

type InformationNewsUpdateReq struct {
	Id                    int       `uri:"id" comment:"主键编码"` // 主键编码
	Title                 string    `json:"title" comment:"标题"`
	EnTitle               string    `json:"enTitle" comment:"英文的标题"`
	Content               string    `json:"content" comment:"内容"`
	EnContent             string    `json:"enContent" comment:"英文的内容"`
	Summary               string    `json:"summary" comment:"摘要"`
	CoverImage            string    `json:"coverImage" comment:"封面图片"`
	Source                string    `json:"source" comment:"来源"`
	SourceId              int64     `json:"sourceId" comment:"来源id"`
	LogUrl                string    `json:"logUrl" comment:"logo图标"`
	CategoryId            int64     `json:"categoryId" comment:"分类ID"`
	RelatedCoinId         string    `json:"relatedCoinId" comment:"相关币种id"`
	RelatedTradingPointId string    `json:"relatedTradingPointId" comment:"相关点位id"`
	Level                 string    `json:"level" comment:"级别：1-普通，2-重要，3-紧急"`
	Sentiment             string    `json:"sentiment" comment:"情绪：1-中性，2-利好，3-利空"`
	IsFeatured            string    `json:"isFeatured" comment:"是否头条：2-否，1-是"`
	ViewCount             int64     `json:"viewCount" comment:"浏览次数"`
	LikeCount             int64     `json:"likeCount" comment:"点赞次数"`
	ShareCount            int64     `json:"shareCount" comment:"分享次数"`
	Status                string    `json:"status" comment:"状态：1-草稿，2-已发布，3-已下架"`
	PublishedAt           time.Time `json:"publishedAt" comment:"发布时间"`
	common.ControlBy
}

func (s *InformationNewsUpdateReq) Generate(model *models.InformationNews) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.Level != "" {
		level, _ := strconv.Atoi(s.Level)
		model.Level = int64(level)
	}
	if s.Sentiment != "" {
		sentiment, _ := strconv.Atoi(s.Sentiment)
		model.Sentiment = int64(sentiment)
	}
	if s.IsFeatured != "" {
		isFeatured, _ := strconv.Atoi(s.IsFeatured)
		model.IsFeatured = int64(isFeatured)
	}
	if s.RelatedCoinId != "" {
		relatedCoinId, _ := strconv.Atoi(s.RelatedCoinId)
		model.RelatedCoinId = int64(relatedCoinId)
	}
	if s.RelatedTradingPointId != "" {
		relatedTradingPointId, _ := strconv.Atoi(s.RelatedTradingPointId)
		model.RelatedTradingPointId = int64(relatedTradingPointId)
	}
	model.Title = s.Title
	model.EnTitle = s.EnTitle
	model.Content = s.Content
	model.EnContent = s.EnContent
	model.Summary = s.Summary
	model.CoverImage = s.CoverImage
	model.Source = s.Source
	model.SourceId = s.SourceId
	model.LogoUrl = s.LogUrl
	model.CategoryId = s.CategoryId

	model.ViewCount = s.ViewCount
	model.LikeCount = s.LikeCount
	model.ShareCount = s.ShareCount
	model.PublishedAt = s.PublishedAt
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationNewsUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationNewsGetReq 功能获取请求参数
type InformationNewsGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationNewsGetReq) GetId() interface{} {
	return s.Id
}

// InformationNewsDeleteReq 功能删除请求参数
type InformationNewsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationNewsDeleteReq) GetId() interface{} {
	return s.Ids
}
