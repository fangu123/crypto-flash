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

type InformationAppVersion struct {
	service.Service
}

// GetPage 获取InformationAppVersion列表
func (e *InformationAppVersion) GetPage(c *dto.InformationAppVersionGetPageReq, p *actions.DataPermission, list *[]models.InformationAppVersion, count *int64) error {
	var err error
	var data models.InformationAppVersion

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
		e.Log.Errorf("InformationAppVersionService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationAppVersion对象
func (e *InformationAppVersion) Get(d *dto.InformationAppVersionGetReq, p *actions.DataPermission, model *models.InformationAppVersion) error {
	var data models.InformationAppVersion

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationAppVersion error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationAppVersion对象
func (e *InformationAppVersion) Insert(c *dto.InformationAppVersionInsertReq) error {
	var err error
	var data models.InformationAppVersion
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationAppVersionService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationAppVersion对象
func (e *InformationAppVersion) Update(c *dto.InformationAppVersionUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InformationAppVersion{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationAppVersionService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除InformationAppVersion
func (e *InformationAppVersion) Remove(d *dto.InformationAppVersionDeleteReq, p *actions.DataPermission) error {
	var data models.InformationAppVersion

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationAppVersion error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
