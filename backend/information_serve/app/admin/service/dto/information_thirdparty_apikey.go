package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strconv"
)

type InformationThirdpartyApikeyGetPageReq struct {
	dto.Pagination `search:"-"`
	Type           int64 `form:"type"  search:"type:exact;column:type;table:information_thirdparty_apikey" comment:"类型：1-coingecko，2-bybit，3-odaily，4-blockBeats"`
	Status         int64 `form:"status"  search:"type:exact;column:status;table:information_thirdparty_apikey" comment:"状态 1-启用 2-停用"`
	InformationThirdpartyApikeyOrder
}

type InformationThirdpartyApikeyOrder struct {
	Id         string `form:"idOrder"  search:"type:order;column:id;table:information_thirdparty_apikey"`
	Type       string `form:"typeOrder"  search:"type:order;column:type;table:information_thirdparty_apikey"`
	BaseUrl    string `form:"baseUrlOrder"  search:"type:order;column:base_url;table:information_thirdparty_apikey"`
	Apikey     string `form:"apikeyOrder"  search:"type:order;column:apikey;table:information_thirdparty_apikey"`
	Secretkey  string `form:"secretkeyOrder"  search:"type:order;column:secretkey;table:information_thirdparty_apikey"`
	Passphrase string `form:"passphraseOrder"  search:"type:order;column:passphrase;table:information_thirdparty_apikey"`
	Status     string `form:"statusOrder"  search:"type:order;column:status;table:information_thirdparty_apikey"`
	CreateBy   string `form:"createByOrder"  search:"type:order;column:create_by;table:information_thirdparty_apikey"`
	UpdateBy   string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_thirdparty_apikey"`
	CreatedAt  string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_thirdparty_apikey"`
	UpdatedAt  string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_thirdparty_apikey"`
	DeletedAt  string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_thirdparty_apikey"`
}

func (m *InformationThirdpartyApikeyGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationThirdpartyApikeyInsertReq struct {
	Id         int    `json:"-" comment:"ID"` // ID
	Type       string `json:"type" comment:"类型：1-coingecko，2-bybit，3-odaily，4-blockBeats"`
	BaseUrl    string `json:"baseUrl" comment:"url地址"`
	Apikey     string `json:"apikey" comment:"apikey"`
	Secretkey  string `json:"secretkey" comment:"secretkey"`
	Passphrase string `json:"passphrase" comment:"passphrase"`
	Status     string `json:"status" comment:"状态 1-启用 2-停用"`
	common.ControlBy
}

func (s *InformationThirdpartyApikeyInsertReq) Generate(model *models.InformationThirdpartyApikey) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.Type != "" {
		t, _ := strconv.Atoi(s.Type)
		model.Type = int64(t)
	}
	model.BaseUrl = s.BaseUrl
	model.Apikey = s.Apikey
	model.Secretkey = s.Secretkey
	model.Passphrase = s.Passphrase
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationThirdpartyApikeyInsertReq) GetId() interface{} {
	return s.Id
}

type InformationThirdpartyApikeyUpdateReq struct {
	Id         int    `uri:"id" comment:"ID"` // ID
	Type       string `json:"type" comment:"类型：1-coingecko，2-bybit，3-odaily，4-blockBeats"`
	BaseUrl    string `json:"baseUrl" comment:"url地址"`
	Apikey     string `json:"apikey" comment:"apikey"`
	Secretkey  string `json:"secretkey" comment:"secretkey"`
	Passphrase string `json:"passphrase" comment:"passphrase"`
	Status     string `json:"status" comment:"状态 1-启用 2-停用"`
	common.ControlBy
}

func (s *InformationThirdpartyApikeyUpdateReq) Generate(model *models.InformationThirdpartyApikey) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.Type != "" {
		t, _ := strconv.Atoi(s.Type)
		model.Type = int64(t)
	}
	model.BaseUrl = s.BaseUrl
	model.Apikey = s.Apikey
	model.Secretkey = s.Secretkey
	model.Passphrase = s.Passphrase
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationThirdpartyApikeyUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationThirdpartyApikeyGetReq 功能获取请求参数
type InformationThirdpartyApikeyGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationThirdpartyApikeyGetReq) GetId() interface{} {
	return s.Id
}

// InformationThirdpartyApikeyDeleteReq 功能删除请求参数
type InformationThirdpartyApikeyDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationThirdpartyApikeyDeleteReq) GetId() interface{} {
	return s.Ids
}
