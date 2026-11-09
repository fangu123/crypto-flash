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

type InformationCoin struct {
	service.Service
}

// GetPage 获取InformationCoin列表
func (e *InformationCoin) GetPage(c *dto.InformationCoinGetPageReq, p *actions.DataPermission, list *[]models.InformationCoin, count *int64) error {
	var err error
	var data models.InformationCoin

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("InformationCoinService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationCoin对象
func (e *InformationCoin) Get(d *dto.InformationCoinGetReq, p *actions.DataPermission, model *models.InformationCoin) error {
	var data models.InformationCoin

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationCoin error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationCoin对象
func (e *InformationCoin) Insert(c *dto.InformationCoinInsertReq) error {
    var err error
    var data models.InformationCoin
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationCoinService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationCoin对象
func (e *InformationCoin) Update(c *dto.InformationCoinUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.InformationCoin{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("InformationCoinService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除InformationCoin
func (e *InformationCoin) Remove(d *dto.InformationCoinDeleteReq, p *actions.DataPermission) error {
	var data models.InformationCoin

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveInformationCoin error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
