package dto

import (
	"strconv"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InformationCommentGetPageReq struct {
	dto.Pagination `search:"-"`
	NewsId         int64     `form:"newsId"  search:"type:exact;column:news_id;table:information_comment" comment:"资讯ID"`
	UserId         int64     `form:"userId"  search:"type:exact;column:user_id;table:information_comment" comment:"评论用户ID"`
	ParentId       int64     `form:"parentId"  search:"type:exact;column:parent_id;table:information_comment" comment:"父评论ID，1表示顶级评论"`
	ReplyToUserId  int64     `form:"replyToUserId"  search:"type:exact;column:reply_to_user_id;table:information_comment" comment:"被回复用户ID"`
	IsTop          int64     `form:"isTop"  search:"type:exact;column:is_top;table:information_comment" comment:"是否置顶 2否 1是"`
	Status         int64     `form:"status"  search:"type:exact;column:status;table:information_comment" comment:"状态 4待审核 1正常 2违规 3已删除"`
	CreateTime     time.Time `form:"createTime"  search:"type:exact;column:created_time;table:information_comment" comment:"发布时间"`
	InformationCommentOrder
}

type InformationCommentOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:information_comment"`
	NewsId        string `form:"newsIdOrder"  search:"type:order;column:news_id;table:information_comment"`
	UserId        string `form:"userIdOrder"  search:"type:order;column:user_id;table:information_comment"`
	ParentId      string `form:"parentIdOrder"  search:"type:order;column:parent_id;table:information_comment"`
	ReplyToUserId string `form:"replyToUserIdOrder"  search:"type:order;column:reply_to_user_id;table:information_comment"`
	Content       string `form:"contentOrder"  search:"type:order;column:content;table:information_comment"`
	LikeCount     string `form:"likeCountOrder"  search:"type:order;column:like_count;table:information_comment"`
	IsTop         string `form:"isTopOrder"  search:"type:order;column:is_top;table:information_comment"`
	Status        string `form:"statusOrder"  search:"type:order;column:status;table:information_comment"`
	IpLocation    string `form:"ipLocationOrder"  search:"type:order;column:ip_location;table:information_comment"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_comment"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_comment"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_comment"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:information_comment"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_comment"`
}

func (m *InformationCommentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationCommentInsertReq struct {
	Id            int       `json:"-" comment:"主键ID"` // 主键ID
	NewsId        string    `json:"newsId" comment:"资讯ID"`
	UserId        string    `json:"userId" comment:"评论用户ID"`
	ParentId      string    `json:"parentId" comment:"父评论ID，1表示顶级评论"`
	ReplyToUserId string    `json:"replyToUserId" comment:"被回复用户ID"`
	Content       string    `json:"content" comment:"评论内容"`
	LikeCount     string    `json:"likeCount" comment:"点赞数"`
	IsTop         string    `json:"isTop" comment:"是否置顶 2否 1是"`
	Status        string    `json:"status" comment:"状态 4待审核 1正常 2违规 3已删除"`
	IpLocation    string    `json:"ipLocation" comment:"IP归属地"`
	CreateTime    time.Time `json:"createTime" comment:"发布时间"`
	UpdateTime    time.Time `json:"updateTime" comment:"更新时间"`
	common.ControlBy
}

func (s *InformationCommentInsertReq) Generate(model *models.InformationComment) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.NewsId != "" {
		newsId, _ := strconv.Atoi(s.NewsId)
		model.NewsId = int64(newsId)
	}
	if s.UserId != "" {
		userId, _ := strconv.Atoi(s.UserId)
		model.UserId = int64(userId)
	}
	if s.ParentId != "" {
		parentId, _ := strconv.Atoi(s.ParentId)
		model.ParentId = int64(parentId)
	}
	if s.ReplyToUserId != "" {
		replyToUserId, _ := strconv.Atoi(s.ReplyToUserId)
		model.ReplyToUserId = int64(replyToUserId)
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.IsTop != "" {
		isTop, _ := strconv.Atoi(s.IsTop)
		model.IsTop = int64(isTop)
	}
	model.Content = s.Content
	if s.LikeCount != "" {
		likeCount, _ := strconv.Atoi(s.LikeCount)
		model.LikeCount = int64(likeCount)
	}
	model.IpLocation = s.IpLocation
	model.CreatedAt = s.CreateTime
	model.UpdatedAt = s.UpdateTime
}

func (s *InformationCommentInsertReq) GetId() interface{} {
	return s.Id
}

type InformationCommentUpdateReq struct {
	Id            int       `uri:"id" comment:"主键ID"` // 主键ID
	NewsId        string    `json:"newsId" comment:"资讯ID"`
	UserId        string    `json:"userId" comment:"评论用户ID"`
	ParentId      string    `json:"parentId" comment:"父评论ID，1表示顶级评论"`
	ReplyToUserId string    `json:"replyToUserId" comment:"被回复用户ID"`
	Content       string    `json:"content" comment:"评论内容"`
	LikeCount     string    `json:"likeCount" comment:"点赞数"`
	IsTop         string    `json:"isTop" comment:"是否置顶 2否 1是"`
	Status        string    `json:"status" comment:"状态 4待审核 1正常 2违规 3已删除"`
	IpLocation    string    `json:"ipLocation" comment:"IP归属地"`
	CreateTime    time.Time `json:"createTime" comment:"发布时间"`
	UpdateTime    time.Time `json:"updateTime" comment:"更新时间"`
	common.ControlBy
}

func (s *InformationCommentUpdateReq) Generate(model *models.InformationComment) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.NewsId != "" {
		newsId, _ := strconv.Atoi(s.NewsId)
		model.NewsId = int64(newsId)
	}
	if s.UserId != "" {
		userId, _ := strconv.Atoi(s.UserId)
		model.UserId = int64(userId)
	}
	if s.ParentId != "" {
		parentId, _ := strconv.Atoi(s.ParentId)
		model.ParentId = int64(parentId)
	}
	if s.ReplyToUserId != "" {
		replyToUserId, _ := strconv.Atoi(s.ReplyToUserId)
		model.ReplyToUserId = int64(replyToUserId)
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.IsTop != "" {
		isTop, _ := strconv.Atoi(s.IsTop)
		model.IsTop = int64(isTop)
	}
	model.Content = s.Content
	if s.LikeCount != "" {
		likeCount, _ := strconv.Atoi(s.LikeCount)
		model.LikeCount = int64(likeCount)
	}
	model.IpLocation = s.IpLocation
	model.CreatedAt = s.CreateTime
	model.UpdatedAt = s.UpdateTime
}

func (s *InformationCommentUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationCommentGetReq 功能获取请求参数
type InformationCommentGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationCommentGetReq) GetId() interface{} {
	return s.Id
}

// InformationCommentDeleteReq 功能删除请求参数
type InformationCommentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationCommentDeleteReq) GetId() interface{} {
	return s.Ids
}
