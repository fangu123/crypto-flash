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

type InformationBlockchain struct {
	service.Service
}

// GetPage 获取InformationBlockchain列表
func (e *InformationBlockchain) GetPage(c *dto.InformationBlockchainGetPageReq, p *actions.DataPermission, list *[]models.InformationBlockchain, count *int64) error {
	var err error
	var data models.InformationBlockchain

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("InformationBlockchainService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationBlockchain对象
func (e *InformationBlockchain) Get(d *dto.InformationBlockchainGetReq, p *actions.DataPermission, model *models.InformationBlockchain) error {
	var data models.InformationBlockchain

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationBlockchain error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationBlockchain对象
func (e *InformationBlockchain) Insert(c *dto.InformationBlockchainInsertReq) error {
	var err error
	var data models.InformationBlockchain
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationBlockchainService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationBlockchain对象
func (e *InformationBlockchain) Update(c *dto.InformationBlockchainUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InformationBlockchain{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationBlockchainService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除InformationBlockchain
func (e *InformationBlockchain) Remove(d *dto.InformationBlockchainDeleteReq, p *actions.DataPermission) error {
	var data models.InformationBlockchain

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationBlockchain error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
