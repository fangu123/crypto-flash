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

type InformationUser struct {
	service.Service
}

// GetPage 获取InformationUser列表
func (e *InformationUser) GetPage(c *dto.InformationUserGetPageReq, p *actions.DataPermission, list *[]models.InformationUser, count *int64) error {
	var err error
	var data models.InformationUser

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
		e.Log.Errorf("InformationUserService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationUser对象
func (e *InformationUser) Get(d *dto.InformationUserGetReq, p *actions.DataPermission, model *models.InformationUser) error {
	var data models.InformationUser

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationUser error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationUser对象
func (e *InformationUser) Insert(c *dto.InformationUserInsertReq) error {
	var err error
	var data models.InformationUser
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationUserService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationUser对象
func (e *InformationUser) Update(c *dto.InformationUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InformationUser{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationUserService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除InformationUser
func (e *InformationUser) Remove(d *dto.InformationUserDeleteReq, p *actions.DataPermission) error {
	var data models.InformationUser

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationUser error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
