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

type InformationTradingPointFavorite struct {
	service.Service
}

// GetPage 获取InformationTradingPointFavorite列表
func (e *InformationTradingPointFavorite) GetPage(c *dto.InformationTradingPointFavoriteGetPageReq, p *actions.DataPermission, list *[]models.InformationTradingPointFavorite, count *int64) error {
	var err error
	var data models.InformationTradingPointFavorite

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("InformationTradingPointFavoriteService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationTradingPointFavorite对象
func (e *InformationTradingPointFavorite) Get(d *dto.InformationTradingPointFavoriteGetReq, p *actions.DataPermission, model *models.InformationTradingPointFavorite) error {
	var data models.InformationTradingPointFavorite

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationTradingPointFavorite error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationTradingPointFavorite对象
func (e *InformationTradingPointFavorite) Insert(c *dto.InformationTradingPointFavoriteInsertReq) error {
    var err error
    var data models.InformationTradingPointFavorite
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationTradingPointFavoriteService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationTradingPointFavorite对象
func (e *InformationTradingPointFavorite) Update(c *dto.InformationTradingPointFavoriteUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.InformationTradingPointFavorite{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("InformationTradingPointFavoriteService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除InformationTradingPointFavorite
func (e *InformationTradingPointFavorite) Remove(d *dto.InformationTradingPointFavoriteDeleteReq, p *actions.DataPermission) error {
	var data models.InformationTradingPointFavorite

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveInformationTradingPointFavorite error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
