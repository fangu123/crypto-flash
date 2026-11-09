package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strconv"
)

type InformationContentTagGetPageReq struct {
	dto.Pagination `search:"-"`
	Name           string `form:"name"  search:"type:contains;column:name;table:information_content_tag" comment:"标签名称"`
	Status         int64  `form:"status"  search:"type:exact;column:status;table:information_content_tag" comment:"状态：2-禁用，1-启用"`
	InformationContentTagOrder
}

type InformationContentTagOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:information_content_tag"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:information_content_tag"`
	Code      string `form:"codeOrder"  search:"type:order;column:code;table:information_content_tag"`
	Status    string `form:"statusOrder"  search:"type:order;column:status;table:information_content_tag"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_content_tag"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_content_tag"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_content_tag"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:information_content_tag"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_content_tag"`
}

func (m *InformationContentTagGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationContentTagInsertReq struct {
	Id     int    `json:"-" comment:"主键编码"` // 主键编码
	Name   string `json:"name" comment:"标签名称"`
	EnName string `json:"enName" comment:"英文的标签名称"`
	Code   string `json:"code" comment:"标签编码"`
	Status string `json:"status" comment:"状态：2-禁用，1-启用"`
	common.ControlBy
}

func (s *InformationContentTagInsertReq) Generate(model *models.InformationContentTag) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	model.Name = s.Name
	model.EnName = s.EnName
	model.Code = s.Code
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationContentTagInsertReq) GetId() interface{} {
	return s.Id
}

type InformationContentTagUpdateReq struct {
	Id     int    `uri:"id" comment:"主键编码"` // 主键编码
	Name   string `json:"name" comment:"标签名称"`
	EnName string `json:"enName" comment:"英文的标签名称"`
	Code   string `json:"code" comment:"标签编码"`
	Status string `json:"status" comment:"状态：2-禁用，1-启用"`
	common.ControlBy
}

func (s *InformationContentTagUpdateReq) Generate(model *models.InformationContentTag) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	model.Name = s.Name
	model.EnName = s.EnName
	model.Code = s.Code
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationContentTagUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationContentTagGetReq 功能获取请求参数
type InformationContentTagGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationContentTagGetReq) GetId() interface{} {
	return s.Id
}

// InformationContentTagDeleteReq 功能删除请求参数
type InformationContentTagDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationContentTagDeleteReq) GetId() interface{} {
	return s.Ids
}
