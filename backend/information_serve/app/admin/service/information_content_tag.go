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

type InformationContentTag struct {
	service.Service
}

// GetPage 获取InformationContentTag列表
func (e *InformationContentTag) GetPage(c *dto.InformationContentTagGetPageReq, p *actions.DataPermission, list *[]models.InformationContentTag, count *int64) error {
	var err error
	var data models.InformationContentTag

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("InformationContentTagService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationContentTag对象
func (e *InformationContentTag) Get(d *dto.InformationContentTagGetReq, p *actions.DataPermission, model *models.InformationContentTag) error {
	var data models.InformationContentTag

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationContentTag error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationContentTag对象
func (e *InformationContentTag) Insert(c *dto.InformationContentTagInsertReq) error {
    var err error
    var data models.InformationContentTag
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationContentTagService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationContentTag对象
func (e *InformationContentTag) Update(c *dto.InformationContentTagUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.InformationContentTag{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("InformationContentTagService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除InformationContentTag
func (e *InformationContentTag) Remove(d *dto.InformationContentTagDeleteReq, p *actions.DataPermission) error {
	var data models.InformationContentTag

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveInformationContentTag error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
