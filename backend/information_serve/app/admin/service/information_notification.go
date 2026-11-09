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

type InformationNotification struct {
	service.Service
}

// GetPage 获取InformationNotification列表
func (e *InformationNotification) GetPage(c *dto.InformationNotificationGetPageReq, p *actions.DataPermission, list *[]models.InformationNotification, count *int64) error {
	var err error
	var data models.InformationNotification

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
		e.Log.Errorf("InformationNotificationService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationNotification对象
func (e *InformationNotification) Get(d *dto.InformationNotificationGetReq, p *actions.DataPermission, model *models.InformationNotification) error {
	var data models.InformationNotification

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationNotification error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationNotification对象
func (e *InformationNotification) Insert(c *dto.InformationNotificationInsertReq) error {
	var err error
	var data models.InformationNotification
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationNotificationService Insert error:%s \r\n", err)
		return err
	}

	//当新增的type为1时 需要向information_notification_user表 插入信息 用户id是开启了系统通知设置的用户NotificationSettings有含4的用户
	// 如果状态为1已发布，则发送通知
	if data.Type == 1 {
		err = e.sendNotification(data)
		if err != nil {
			e.Log.Errorf("Failed to send notification: %v", err)
		}
	}
	return nil
}

// Update 修改InformationNotification对象
func (e *InformationNotification) Update(c *dto.InformationNotificationUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InformationNotification{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationNotificationService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除InformationNotification
func (e *InformationNotification) Remove(d *dto.InformationNotificationDeleteReq, p *actions.DataPermission) error {
	var data models.InformationNotification

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationNotification error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *InformationNotification) sendNotification(news models.InformationNotification) error {
	// 查询开启了资讯通知的用户
	var Users []models.InformationUser
	err2 := e.Orm.Where("JSON_CONTAINS(notification_settings, JSON_QUOTE(?))", "4").Find(&Users).Error
	if err2 != nil {
		e.Log.Errorf("Failed to get notification settings: %v", err2)
		return err2
	}

	if len(Users) == 0 {
		return nil
	}

	// 批量创建通知关联
	var notificationUsers []models.InformationNotificationUser
	for _, user := range Users {
		notificationUsers = append(notificationUsers, models.InformationNotificationUser{
			UserId:         int64(user.Id),
			NotificationId: int64(news.Id),
			IsRead:         2,
		})
	}

	// 批量插入通知关联
	err3 := e.Orm.CreateInBatches(notificationUsers, 100).Error
	if err3 != nil {
		e.Log.Errorf("Failed to batch create notification users: %v", err3)
		return err3
	}

	return nil
}
