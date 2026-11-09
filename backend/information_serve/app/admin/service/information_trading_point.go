package service

import (
	"errors"
	"strconv"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type InformationTradingPoint struct {
	service.Service
}

// GetPage 获取InformationTradingPoint列表
func (e *InformationTradingPoint) GetPage(c *dto.InformationTradingPointGetPageReq, p *actions.DataPermission, list *[]models.InformationTradingPoint, count *int64) error {
	var err error
	var data models.InformationTradingPoint

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Order("created_at DESC").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	//拼接list中的信息返给name字段
	if *count > 0 {
		for i := 0; i < len(*list); i++ {
			coin := models.InformationCoin{}
			e.Orm.First(&coin, (*list)[i].CoinId)
			price := (*list)[i].EntryPrice
			if (*list)[i].Type == 1 {
				(*list)[i].Name = coin.Symbol + "_做多_" + price
			} else {
				(*list)[i].Name = coin.Symbol + "_做空_" + price
			}
		}

	}
	if err != nil {
		e.Log.Errorf("InformationTradingPointService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InformationTradingPoint对象
func (e *InformationTradingPoint) Get(d *dto.InformationTradingPointGetReq, p *actions.DataPermission, model *models.InformationTradingPoint) error {
	var data models.InformationTradingPoint

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInformationTradingPoint error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InformationTradingPoint对象
func (e *InformationTradingPoint) Insert(c *dto.InformationTradingPointInsertReq) error {
	var err error
	var data models.InformationTradingPoint
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("InformationTradingPointService Insert error:%s \r\n", err)
		return err
	}
	//当前status为1进行中且isNotified状态为2已通知 就需要把标题（点位币种名称 类型做多还是多空） 内容（点位币种名称 类型做多还是多空 入场价格 止盈价格数组 止损价格 有效时长 到期时间） 需要写入到information_notification表标题 内容 type为2。 并且向information_notification_user表 插入信息 用户id是开启了新点位设置且点位权限是所有的用户NotificationSettings有含2的用户且point_notification_settings为2
	//当前status为2或者3且isNotified状态为2已通知 就需要把标题（点位币种名称 类型做多还是多空） 内容（点位币种名称 类型做多还是多空 入场价格 止盈价格数组 止损价格 有效时长 到期时间） 需要写入到information_notification表标题 内容 type为4。 并且向information_notification_user表 插入信息 用户id是开启了点位状态设置且点位权限是所有的用户NotificationSettings有含3的用户且point_notification_settings为2
	if data.Status == 1 && data.IsNotified == 2 { //进行中且已通知
		err = e.sendNotification(data, 1)
		if err != nil {
			e.Log.Errorf("Failed to send notification: %v", err)
		}
	} else if (data.Status == 2 || data.Status == 3) && data.IsNotified == 2 { //结束且已通知
		err = e.sendNotification(data, 2)
		if err != nil {
			e.Log.Errorf("Failed to send notification: %v", err)
		}
	}
	return nil
}

// Update 修改InformationTradingPoint对象
func (e *InformationTradingPoint) Update(c *dto.InformationTradingPointUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InformationTradingPoint{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())

	if data.IsNotified == 2 && data.Status == 2 && (c.Status == "1" || c.Status == "2" || c.Status == "3") {
		return errors.New("已失效不能修改状态")
	}
	if data.IsNotified == 2 && data.Status == 3 && (c.Status == "1" || c.Status == "2" || c.Status == "3") {
		return errors.New("已撤销不能修改状态")
	}

	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InformationTradingPointService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	//当前status为1进行中且isNotified状态为2已通知 就需要把标题（点位币种名称 类型做多还是多空） 内容（点位币种名称 类型做多还是多空 入场价格 止盈价格数组 止损价格 有效时长 到期时间） 需要写入到information_notification表标题 内容 type为2。 并且向information_notification_user表 插入信息 用户id是开启了新点位设置且点位权限是所有的用户NotificationSettings有含2的用户且point_notification_settings为2
	//当前status为2或者3且isNotified状态为2已通知 就需要把标题（点位币种名称 类型做多还是多空） 内容（点位币种名称 类型做多还是多空 入场价格 止盈价格数组 止损价格 有效时长 到期时间） 需要写入到information_notification表标题 内容 type为4。 并且向information_notification_user表 插入信息 用户id是开启了点位状态设置且点位权限是所有的用户NotificationSettings有含3的用户且point_notification_settings为2
	if data.Status == 1 && data.IsNotified == 2 { //进行中且已通知
		err = e.sendNotification(data, 1)
		if err != nil {
			e.Log.Errorf("Failed to send notification: %v", err)
		}
	} else if (data.Status == 2 || data.Status == 3) && data.IsNotified == 2 { //结束且已通知
		err = e.sendNotification(data, 2)
		if err != nil {
			e.Log.Errorf("Failed to send notification: %v", err)
		}
	}
	return nil
}

// Remove 删除InformationTradingPoint
func (e *InformationTradingPoint) Remove(d *dto.InformationTradingPointDeleteReq, p *actions.DataPermission) error {
	var data models.InformationTradingPoint

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInformationTradingPoint error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func (e *InformationTradingPoint) sendNotification(data models.InformationTradingPoint, i int) error {
	var t = 2       //新点位
	var setting = 2 //新点位
	if i == 2 {
		t = 4       //点位状态
		setting = 3 //点位状态
	}
	var statsName string
	var enStatusName string
	if data.Status == 1 {
		statsName = "新增"
		enStatusName = "new"
	} else if data.Status == 2 {
		statsName = "失效"
		enStatusName = "expired"
	} else if data.Status == 3 {
		statsName = "撤销"
		enStatusName = "canceled"
	}
	coin := models.InformationCoin{}
	e.Orm.First(&coin, data.CoinId)
	var enName string
	if data.Type == 1 {
		data.Name = coin.Symbol + "_做多" + ":" + statsName
		enName = coin.Symbol + "_long" + ":" + enStatusName
	} else {
		data.Name = coin.Symbol + "_做空" + ":" + statsName
		enName = coin.Symbol + "_short" + ":" + enStatusName
	}
	//title 币名称+做多还是做空
	//content 币种名称 类型做多还是多空 入场价格 止盈价格数组 止损价格 有效时长 到期时间
	notification := models.InformationNotification{
		Title:     data.Name,
		EnTitle:   enName,
		Content:   data.Name + "，入场价格是" + data.EntryPrice + "，止盈价格是" + data.TakeProfits + "，止损价格是" + data.StopLoss + "，有效时长是" + strconv.FormatInt(data.ValidDuration, 10) + "小时，截止时间是" + data.ValidUntil.Format("2006-01-02 15:04:05"),
		EnContent: enName + "，entry price is " + data.EntryPrice + "，take profits is " + data.TakeProfits + "，stop loss is " + data.StopLoss + "，valid duration is " + strconv.FormatInt(data.ValidDuration, 10) + " hours，deadline is " + data.ValidUntil.Format("2006-01-02 15:04:05"),
		Type:      int64(t), // 点位通知类型
	}
	err := e.Orm.Create(&notification).Error
	if err != nil {
		e.Log.Errorf("Failed to create notification : %v", err)
		return err
	}

	// 查询开启了点位通知的用户
	var Users []models.InformationUser
	err2 := e.Orm.Where("JSON_CONTAINS(notification_settings, JSON_QUOTE(?))  AND point_notification_settings = 2", strconv.Itoa(setting)).Find(&Users).Error

	if err2 != nil {
		e.Log.Errorf("Failed to get User: %v", err2)
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
