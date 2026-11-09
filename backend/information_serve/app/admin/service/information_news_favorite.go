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

type InformationNewsFavorite struct {
	service.Service
}

// GetPage 获取InformationNewsFavorite列表
func (e *InformationNewsFavorite) GetPage(c *dto.InformationNewsFavoriteGetPageReq, p *actions.DataPermission, list *[]models.InformationNewsFavorite, count *int64) error {
	var err error
	var data models.InformationNewsFavorite

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
		e.Log.Errorf("InformationNewsFavoriteService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationNewsFavorite对象
func (e *InformationNewsFavorite) Get(d *dto.InformationNewsFavoriteGetReq, p *actions.DataPermission, model *models.InformationNewsFavorite) error {
	var data models.InformationNewsFavorite

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationNewsFavorite error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationNewsFavorite对象
func (e *InformationNewsFavorite) Insert(c *dto.InformationNewsFavoriteInsertReq) error {
	var err error
	var data models.InformationNewsFavorite
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationNewsFavoriteService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationNewsFavorite对象
func (e *InformationNewsFavorite) Update(c *dto.InformationNewsFavoriteUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InformationNewsFavorite{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationNewsFavoriteService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除InformationNewsFavorite
func (e *InformationNewsFavorite) Remove(d *dto.InformationNewsFavoriteDeleteReq, p *actions.DataPermission) error {
	var data models.InformationNewsFavorite

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationNewsFavorite error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
