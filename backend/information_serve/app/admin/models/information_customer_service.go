package models

import (

	"go-admin/common/models"

)

type InformationCustomerService struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(64);comment:渠道名称"` 
    Icon string `json:"icon" gorm:"type:varchar(128);comment:图标地址"` 
    CustomerName string `json:"customerName" gorm:"type:varchar(64);comment:客服名称"` 
    Avatar string `json:"avatar" gorm:"type:varchar(256);comment:头像"` 
    Email string `json:"email" gorm:"type:varchar(128);comment:邮箱"` 
    Phone string `json:"phone" gorm:"type:varchar(32);comment:电话"` 
    Sort int64 `json:"sort" gorm:"type:int;comment:排序序号"` 
    Status int64 `json:"status" gorm:"type:tinyint;comment:状态 1-启用 2-停用"` 
    models.ModelTime
    models.ControlBy
}

func (InformationCustomerService) TableName() string {
    return "information_customer_service"
}

func (e *InformationCustomerService) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InformationCustomerService) GetId() interface{} {
	return e.Id
}