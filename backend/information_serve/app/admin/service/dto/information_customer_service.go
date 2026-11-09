package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strconv"
)

type InformationCustomerServiceGetPageReq struct {
	dto.Pagination `search:"-"`
	Status         int64 `form:"status"  search:"type:exact;column:status;table:information_customer_service" comment:"状态 1-启用 2-停用"`
	InformationCustomerServiceOrder
}

type InformationCustomerServiceOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:information_customer_service"`
	Name         string `form:"nameOrder"  search:"type:order;column:name;table:information_customer_service"`
	Icon         string `form:"iconOrder"  search:"type:order;column:icon;table:information_customer_service"`
	CustomerName string `form:"customerNameOrder"  search:"type:order;column:customer_name;table:information_customer_service"`
	Avatar       string `form:"avatarOrder"  search:"type:order;column:avatar;table:information_customer_service"`
	Email        string `form:"emailOrder"  search:"type:order;column:email;table:information_customer_service"`
	Phone        string `form:"phoneOrder"  search:"type:order;column:phone;table:information_customer_service"`
	Sort         string `form:"sortOrder"  search:"type:order;column:sort;table:information_customer_service"`
	Status       string `form:"statusOrder"  search:"type:order;column:status;table:information_customer_service"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_customer_service"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_customer_service"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_customer_service"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:information_customer_service"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_customer_service"`
}

func (m *InformationCustomerServiceGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationCustomerServiceInsertReq struct {
	Id           int    `json:"-" comment:"主键编码"` // 主键编码
	Name         string `json:"name" comment:"渠道名称"`
	Icon         string `json:"icon" comment:"图标地址"`
	CustomerName string `json:"customerName" comment:"客服名称"`
	Avatar       string `json:"avatar" comment:"头像"`
	Email        string `json:"email" comment:"邮箱"`
	Phone        string `json:"phone" comment:"电话"`
	Sort         int64  `json:"sort" comment:"排序序号"`
	Status       string `json:"status" comment:"状态 1-启用 2-停用"`
	common.ControlBy
}

func (s *InformationCustomerServiceInsertReq) Generate(model *models.InformationCustomerService) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	model.Name = s.Name
	model.Icon = s.Icon
	model.CustomerName = s.CustomerName
	model.Avatar = s.Avatar
	model.Email = s.Email
	model.Phone = s.Phone
	model.Sort = s.Sort
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationCustomerServiceInsertReq) GetId() interface{} {
	return s.Id
}

type InformationCustomerServiceUpdateReq struct {
	Id           int    `uri:"id" comment:"主键编码"` // 主键编码
	Name         string `json:"name" comment:"渠道名称"`
	Icon         string `json:"icon" comment:"图标地址"`
	CustomerName string `json:"customerName" comment:"客服名称"`
	Avatar       string `json:"avatar" comment:"头像"`
	Email        string `json:"email" comment:"邮箱"`
	Phone        string `json:"phone" comment:"电话"`
	Sort         int64  `json:"sort" comment:"排序序号"`
	Status       string `json:"status" comment:"状态 1-启用 2-停用"`
	common.ControlBy
}

func (s *InformationCustomerServiceUpdateReq) Generate(model *models.InformationCustomerService) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	model.Name = s.Name
	model.Icon = s.Icon
	model.CustomerName = s.CustomerName
	model.Avatar = s.Avatar
	model.Email = s.Email
	model.Phone = s.Phone
	model.Sort = s.Sort
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationCustomerServiceUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationCustomerServiceGetReq 功能获取请求参数
type InformationCustomerServiceGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationCustomerServiceGetReq) GetId() interface{} {
	return s.Id
}

// InformationCustomerServiceDeleteReq 功能删除请求参数
type InformationCustomerServiceDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationCustomerServiceDeleteReq) GetId() interface{} {
	return s.Ids
}
