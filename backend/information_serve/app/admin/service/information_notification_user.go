package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type InformationNotificationUser struct {
	service.Service
}

// GetPage 获取InformationNotificationUser列表
func (e *InformationNotificationUser) GetPage(c *dto.InformationNotificationUserGetPageReq, p *actions.DataPermission, list *[]models.InformationNotificationUser, count *int64) error {
	var err error
	var data models.InformationNotificationUser

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Order("created_at DESC").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("InformationNotificationUserService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationNotificationUser对象
func (e *InformationNotificationUser) Get(d *dto.InformationNotificationUserGetReq, p *actions.DataPermission, model *models.InformationNotificationUser) error {
	var data models.InformationNotificationUser

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationNotificationUser error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationNotificationUser对象
func (e *InformationNotificationUser) Insert(c *dto.InformationNotificationUserInsertReq) error {
	var err error
	var data models.InformationNotificationUser
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationNotificationUserService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationNotificationUser对象
func (e *InformationNotificationUser) Update(c *dto.InformationNotificationUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InformationNotificationUser{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationNotificationUserService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除InformationNotificationUser
func (e *InformationNotificationUser) Remove(d *dto.InformationNotificationUserDeleteReq, p *actions.DataPermission) error {
	var data models.InformationNotificationUser

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationNotificationUser error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
