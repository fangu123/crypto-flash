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

type InformationComment struct {
	service.Service
}

// GetPage 获取InformationComment列表
func (e *InformationComment) GetPage(c *dto.InformationCommentGetPageReq, p *actions.DataPermission, list *[]models.InformationComment, count *int64) error {
	var err error
	var data models.InformationComment

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
		e.Log.Errorf("InformationCommentService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationComment对象
func (e *InformationComment) Get(d *dto.InformationCommentGetReq, p *actions.DataPermission, model *models.InformationComment) error {
	var data models.InformationComment

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationComment error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationComment对象
func (e *InformationComment) Insert(c *dto.InformationCommentInsertReq) error {
	var err error
	var data models.InformationComment
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationCommentService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改InformationComment对象
func (e *InformationComment) Update(c *dto.InformationCommentUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InformationComment{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationCommentService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除InformationComment
func (e *InformationComment) Remove(d *dto.InformationCommentDeleteReq, p *actions.DataPermission) error {
	var data models.InformationComment

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationComment error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
