package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strconv"
)

type InformationBlockchainGetPageReq struct {
	dto.Pagination `search:"-"`
	Status         int64 `form:"status"  search:"type:exact;column:status;table:information_blockchain" comment:"状态：2-禁用，1-启用"`
	InformationBlockchainOrder
}

type InformationBlockchainOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:information_blockchain"`
	Name        string `form:"nameOrder"  search:"type:order;column:name;table:information_blockchain"`
	Symbol      string `form:"symbolOrder"  search:"type:order;column:symbol;table:information_blockchain"`
	LogoUrl     string `form:"logoUrlOrder"  search:"type:order;column:logo_url;table:information_blockchain"`
	Description string `form:"descriptionOrder"  search:"type:order;column:description;table:information_blockchain"`
	Website     string `form:"websiteOrder"  search:"type:order;column:website;table:information_blockchain"`
	Whitepaper  string `form:"whitepaperOrder"  search:"type:order;column:whitepaper;table:information_blockchain"`
	Status      string `form:"statusOrder"  search:"type:order;column:status;table:information_blockchain"`
	SortOrder   string `form:"sortOrderOrder"  search:"type:order;column:sort_order;table:information_blockchain"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_blockchain"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_blockchain"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_blockchain"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:information_blockchain"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_blockchain"`
}

func (m *InformationBlockchainGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationBlockchainInsertReq struct {
	Id          int    `json:"-" comment:"主键编码"` // 主键编码
	Name        string `json:"name" comment:"链名称"`
	Symbol      string `json:"symbol" comment:"链标识"`
	LogoUrl     string `json:"logoUrl" comment:"链Logo URL"`
	Description string `json:"description" comment:"链描述"`
	Website     string `json:"website" comment:"官方网站"`
	Whitepaper  string `json:"whitepaper" comment:"白皮书链接"`
	Status      string `json:"status" comment:"状态：2-禁用，1-启用"`
	SortOrder   int64  `json:"sortOrder" comment:"排序"`
	common.ControlBy
}

func (s *InformationBlockchainInsertReq) Generate(model *models.InformationBlockchain) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	model.Name = s.Name
	model.Symbol = s.Symbol
	model.LogoUrl = s.LogoUrl
	model.Description = s.Description
	model.Website = s.Website
	model.Whitepaper = s.Whitepaper
	model.SortOrder = s.SortOrder
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationBlockchainInsertReq) GetId() interface{} {
	return s.Id
}

type InformationBlockchainUpdateReq struct {
	Id          int    `uri:"id" comment:"主键编码"` // 主键编码
	Name        string `json:"name" comment:"链名称"`
	Symbol      string `json:"symbol" comment:"链标识"`
	LogoUrl     string `json:"logoUrl" comment:"链Logo URL"`
	Description string `json:"description" comment:"链描述"`
	Website     string `json:"website" comment:"官方网站"`
	Whitepaper  string `json:"whitepaper" comment:"白皮书链接"`
	Status      string `json:"status" comment:"状态：2-禁用，1-启用"`
	SortOrder   int64  `json:"sortOrder" comment:"排序"`
	common.ControlBy
}

func (s *InformationBlockchainUpdateReq) Generate(model *models.InformationBlockchain) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	model.Name = s.Name
	model.Symbol = s.Symbol
	model.LogoUrl = s.LogoUrl
	model.Description = s.Description
	model.Website = s.Website
	model.Whitepaper = s.Whitepaper
	model.SortOrder = s.SortOrder
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationBlockchainUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationBlockchainGetReq 功能获取请求参数
type InformationBlockchainGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationBlockchainGetReq) GetId() interface{} {
	return s.Id
}

// InformationBlockchainDeleteReq 功能删除请求参数
type InformationBlockchainDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationBlockchainDeleteReq) GetId() interface{} {
	return s.Ids
}
