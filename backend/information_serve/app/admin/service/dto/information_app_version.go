package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"strconv"
)

type InformationAppVersionGetPageReq struct {
	dto.Pagination `search:"-"`
	Platform       int64 `form:"platform"  search:"type:exact;column:platform;table:information_app_version" comment:"平台：1-iOS，2-Android"`
	ForceUpdate    int64 `form:"forceUpdate"  search:"type:exact;column:force_update;table:information_app_version" comment:"是否强制更新：2-否，1-是"`
	UpdateType     int64 `form:"updateType"  search:"type:exact;column:update_type;table:information_app_version" comment:"类型 1-功能更新 2-问题修复 3-性能优化"`
	Status         int64 `form:"status"  search:"type:exact;column:status;table:information_app_version" comment:"状态 1-发布 2-草稿 3-撤回"`
	InformationAppVersionOrder
}

type InformationAppVersionOrder struct {
	Id                string `form:"idOrder"  search:"type:order;column:id;table:information_app_version"`
	Version           string `form:"versionOrder"  search:"type:order;column:version;table:information_app_version"`
	BuildNumber       string `form:"buildNumberOrder"  search:"type:order;column:build_number;table:information_app_version"`
	Platform          string `form:"platformOrder"  search:"type:order;column:platform;table:information_app_version"`
	DownloadUrl       string `form:"downloadUrlOrder"  search:"type:order;column:download_url;table:information_app_version"`
	ForceUpdate       string `form:"forceUpdateOrder"  search:"type:order;column:force_update;table:information_app_version"`
	Content           string `form:"contentOrder"  search:"type:order;column:content;table:information_app_version"`
	UpdateType        string `form:"updateTypeOrder"  search:"type:order;column:update_type;table:information_app_version"`
	Channel           string `form:"channelOrder"  search:"type:order;column:channel;table:information_app_version"`
	Status            string `form:"statusOrder"  search:"type:order;column:status;table:information_app_version"`
	File              string `form:"fileOrder"  search:"type:order;column:file;table:information_app_version"`
	VersionHistoryLog string `form:"versionHistoryLogOrder"  search:"type:order;column:version_history_log;table:information_app_version"`
	CreatedAt         string `form:"createdAtOrder"  search:"type:order;column:created_at;table:information_app_version"`
	UpdatedAt         string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:information_app_version"`
	DeletedAt         string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:information_app_version"`
	CreateBy          string `form:"createByOrder"  search:"type:order;column:create_by;table:information_app_version"`
	UpdateBy          string `form:"updateByOrder"  search:"type:order;column:update_by;table:information_app_version"`
}

func (m *InformationAppVersionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InformationAppVersionInsertReq struct {
	Id                  int    `json:"-" comment:"主键编码"` // 主键编码
	Version             string `json:"version" comment:"版本号"`
	BuildNumber         string `json:"buildNumber" comment:"构建号"`
	Platform            string `json:"platform" comment:"平台：1-iOS，2-Android"`
	DownloadUrl         string `json:"downloadUrl" comment:"下载地址"`
	ForceUpdate         string `json:"forceUpdate" comment:"是否强制更新：2-否，1-是"`
	Content             string `json:"content" comment:"当前更新内容"`
	EnContent           string `json:"enContent" comment:"英文的当前更新内容"`
	UpdateType          string `json:"updateType" comment:"类型 1-功能更新 2-问题修复 3-性能优化"`
	Channel             string `json:"channel" comment:"发布渠道"`
	Status              string `json:"status" comment:"状态 1-发布 2-草稿 3-撤回"`
	File                string `json:"file" comment:"文件"`
	VersionHistoryLog   string `json:"versionHistoryLog" comment:"版本历史日志"`
	EnVersionHistoryLog string `json:"enVersionHistoryLog" comment:"英文的版本历史日志"`
	common.ControlBy
}

func (s *InformationAppVersionInsertReq) Generate(model *models.InformationAppVersion) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.ForceUpdate != "" {
		forceUpdate, _ := strconv.Atoi(s.ForceUpdate)
		model.ForceUpdate = int64(forceUpdate)
	}
	if s.UpdateType != "" {
		updateType, _ := strconv.Atoi(s.UpdateType)
		model.UpdateType = int64(updateType)
	}
	if s.Platform != "" {
		platform, _ := strconv.Atoi(s.Platform)
		model.Platform = int64(platform)
	}
	//model.Status = s.Status
	//model.UpdateType = s.UpdateType
	//model.ForceUpdate = s.ForceUpdate
	//model.Platform = s.Platform
	model.Version = s.Version
	model.BuildNumber = s.BuildNumber
	model.DownloadUrl = s.DownloadUrl
	model.Content = s.Content
	model.EnContent = s.EnContent
	model.Channel = s.Channel
	model.File = s.File
	model.VersionHistoryLog = s.VersionHistoryLog
	model.EnVersionHistoryLog = s.EnVersionHistoryLog
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *InformationAppVersionInsertReq) GetId() interface{} {
	return s.Id
}

type InformationAppVersionUpdateReq struct {
	Id                  int    `uri:"id" comment:"主键编码"` // 主键编码
	Version             string `json:"version" comment:"版本号"`
	BuildNumber         string `json:"buildNumber" comment:"构建号"`
	Platform            string `json:"platform" comment:"平台：1-iOS，2-Android"`
	DownloadUrl         string `json:"downloadUrl" comment:"下载地址"`
	ForceUpdate         string `json:"forceUpdate" comment:"是否强制更新：2-否，1-是"`
	Content             string `json:"content" comment:"当前更新内容"`
	EnContent           string `json:"enContent" comment:"英文的当前更新内容"`
	UpdateType          string `json:"updateType" comment:"类型 1-功能更新 2-问题修复 3-性能优化"`
	Channel             string `json:"channel" comment:"发布渠道"`
	Status              string `json:"status" comment:"状态 1-发布 2-草稿 3-撤回"`
	File                string `json:"file" comment:"文件"`
	VersionHistoryLog   string `json:"versionHistoryLog" comment:"版本历史日志"`
	EnVersionHistoryLog string `json:"enVersionHistoryLog" comment:"英文的版本历史日志"`
	common.ControlBy
}

func (s *InformationAppVersionUpdateReq) Generate(model *models.InformationAppVersion) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	if s.Status != "" {
		status, _ := strconv.Atoi(s.Status)
		model.Status = int64(status)
	}
	if s.ForceUpdate != "" {
		forceUpdate, _ := strconv.Atoi(s.ForceUpdate)
		model.ForceUpdate = int64(forceUpdate)
	}
	if s.UpdateType != "" {
		updateType, _ := strconv.Atoi(s.UpdateType)
		model.UpdateType = int64(updateType)
	}
	if s.Platform != "" {
		platform, _ := strconv.Atoi(s.Platform)
		model.Platform = int64(platform)
	}
	//model.Status = s.Status
	//model.UpdateType = s.UpdateType
	//model.ForceUpdate = s.ForceUpdate
	//model.Platform = s.Platform
	model.Version = s.Version
	model.BuildNumber = s.BuildNumber
	model.DownloadUrl = s.DownloadUrl
	model.Content = s.Content
	model.EnContent = s.EnContent
	model.Channel = s.Channel
	model.File = s.File
	model.VersionHistoryLog = s.VersionHistoryLog
	model.EnVersionHistoryLog = s.EnVersionHistoryLog
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InformationAppVersionUpdateReq) GetId() interface{} {
	return s.Id
}

// InformationAppVersionGetReq 功能获取请求参数
type InformationAppVersionGetReq struct {
	Id int `uri:"id"`
}

func (s *InformationAppVersionGetReq) GetId() interface{} {
	return s.Id
}

// InformationAppVersionDeleteReq 功能删除请求参数
type InformationAppVersionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InformationAppVersionDeleteReq) GetId() interface{} {
	return s.Ids
}
