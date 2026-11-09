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

type InformationNews struct {
	service.Service
}

// StatisticsResp 统计数据返回结构
type StatisticsResp struct {
	News         NewsStatistics      `json:"news"`
	TradingPoint PointStatistics     `json:"tradingPoint"`
	User         UserStatistics      `json:"user"`
	ReadCount    ReadCountStatistics `json:"readCount"`
}

// NewsStatistics 快讯统计数据
type NewsStatistics struct {
	Total int64   `json:"total"` // 总数量
	Today int64   `json:"today"` // 今日增加数量
	Trend []int64 `json:"trend"` // 最近7天的趋势数据
}

// PointStatistics 点位统计数据
type PointStatistics struct {
	Total int64   `json:"total"` // 总数量
	Today int64   `json:"today"` // 今日增加数量
	Trend []int64 `json:"trend"` // 最近7天的趋势数据
}

// UserStatistics 用户统计数据
type UserStatistics struct {
	Total int64 `json:"total"` // 总数量
	Today int64 `json:"today"` // 今日增加数量
}

// ReadCountStatistics 阅读次数统计数据
type ReadCountStatistics struct {
	Total int64 `json:"total"` // 总阅读次数
	Today int64 `json:"today"` // 今日阅读次数
}

// GetPage 获取InformationNews列表
func (e *InformationNews) GetPage(c *dto.InformationNewsGetPageReq, p *actions.DataPermission, list *[]models.InformationNews, count *int64) error {
	var err error
	var data models.InformationNews

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
		e.Log.Errorf("InformationNewsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationNews对象
func (e *InformationNews) Get(d *dto.InformationNewsGetReq, p *actions.DataPermission, model *models.InformationNews) error {
	var data models.InformationNews

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationNews error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationNews对象
func (e *InformationNews) Insert(c *dto.InformationNewsInsertReq) error {
	var err error
	var data models.InformationNews
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationNewsService Insert error:%s \r\n", err)
		return err
	}
	//当前状态为2已发布 就需要把标题 内容需要写入到information_notification表type为3  并且向information_notification_user表 插入信息 用户id是开启了资讯通知设置的用户NotificationSettings有含1的用户
	// 如果状态为2已发布，则发送通知
	if data.Status == 2 {
		err = e.sendNotification(data)
		if err != nil {
			e.Log.Errorf("Failed to send notification: %v", err)
		}
	}
	return nil
}

// Update 修改InformationNews对象
func (e *InformationNews) Update(c *dto.InformationNewsUpdateReq, p *actions.DataPermission) error {

	var err error
	var data = models.InformationNews{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	if data.Status == 2 && c.Status == "1" {
		return errors.New("已发布不能修改为草稿状态")
	}
	if data.Status == 3 && (c.Status == "1" || c.Status == "2") {
		return errors.New("已下架不能修改状态")
	}
	c.Generate(&data)
	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationNewsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	//当前状态改为2已发布 就需要把标题 内容需要写入到information_notification表type为3  并且向information_notification_user表 插入信息 用户id是开启了资讯通知设置的用户NotificationSettings有含1的用户
	// 如果状态为2已发布，则发送通知
	if data.Status == 2 {
		err = e.sendNotification(data)
		if err != nil {
			e.Log.Errorf("Failed to send notification: %v", err)
		}
	}
	return nil
}

// Remove 删除InformationNews
func (e *InformationNews) Remove(d *dto.InformationNewsDeleteReq, p *actions.DataPermission) error {
	var data models.InformationNews

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationNews error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *InformationNews) sendNotification(news models.InformationNews) error {
	notification := models.InformationNotification{
		Title:     news.Title,
		EnTitle:   news.EnTitle,
		Content:   news.Content,
		EnContent: news.EnContent,
		Type:      3, // 资讯通知类型
	}
	err := e.Orm.Create(&notification).Error
	if err != nil {
		e.Log.Errorf("Failed to create notification : %v", err)
		return err
	}

	// 查询开启了资讯通知的用户
	var Users []models.InformationUser
	err2 := e.Orm.Where("JSON_CONTAINS(notification_settings, JSON_QUOTE(?))", "1").Find(&Users).Error
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
			NotificationId: int64(notification.Id),
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
