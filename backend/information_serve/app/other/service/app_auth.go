package service

import (
	"errors"
	"fmt"
	fault "go-admin/common/errors"
	"go-admin/common/redishelper"
	"strings"
	"time"

	adminModels "go-admin/app/admin/models"

	coreService "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
)

type AppAuth struct {
	coreService.Service
}

// Register 注册用户（已在上层做邮箱、密码规则校验）
func (e *AppAuth) Register(email, password, code, deviceId, systemVersion, appVersion, lang string, deviceType int64) (int, error) {
	//email = strings.TrimSpace(strings.ToLower(email))
	// 从Redis获取验证码
	cacheKey := fmt.Sprintf("email_code:%s", email+"_1")
	cachedCode, er := redishelper.DefaultRedis.Get(cacheKey)
	if er != nil || cachedCode == "" {
		return 0, errors.New(fault.GetErrorMessage("please_send_the_verification_code", lang))
	}

	// 验证码校验
	if cachedCode != code {
		return 0, errors.New(fault.GetErrorMessage("invalid_verification_code", lang))
	}

	var exists adminModels.InformationUser
	err := e.Orm.Model(&adminModels.InformationUser{}).
		Where("email = ?", email).
		First(&exists).Error
	if err == nil {
		return 0, errors.New(fault.GetErrorMessage("the_email_has_been_registered", lang))
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	u := adminModels.InformationUser{
		Username:                  email,
		Email:                     email,
		Status:                    1,
		Password:                  password,
		PointNotificationSettings: 1,
		TwoFactorEnabled:          2,
		Language:                  "cn",
		DeviceId:                  deviceId, SystemVersion: systemVersion, AppVersion: appVersion, DeviceType: deviceType,
		NotificationSettings: `["1", "2", "3", "4"]`,
		LastLoginAt:          time.Now(),
	}
	if err := e.Orm.Create(&u).Error; err != nil {
		return 0, err
	}
	// 删除已使用的验证码
	redishelper.DefaultRedis.Delete(cacheKey)
	return u.Id, nil
}

// WalletLogin 通过钱包地址注册和登录
func (e *AppAuth) WalletLogin(walletAddress, deviceId, systemVersion, appVersion, lang string, deviceType int64) (int, error) {
	walletAddress = strings.TrimSpace(walletAddress)

	// 查询用户是否已存在
	var exists adminModels.InformationUser
	err := e.Orm.Model(&adminModels.InformationUser{}).
		Where("wallet_address = ?", walletAddress).
		First(&exists).Error

	// 如果用户不存在，则创建新用户
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新用户
			u := adminModels.InformationUser{
				Username:                  walletAddress,
				Email:                     "",
				WalletAddress:             walletAddress,
				Status:                    1,
				Password:                  "",
				PointNotificationSettings: 1,
				TwoFactorEnabled:          2,
				Language:                  "cn",
				DeviceId:                  deviceId,
				SystemVersion:             systemVersion,
				AppVersion:                appVersion,
				DeviceType:                deviceType,
				NotificationSettings:      `["1", "2", "3", "4"]`,
				LastLoginAt:               time.Now(),
			}

			if err := e.Orm.Create(&u).Error; err != nil {
				return 0, err
			}

			// 更新最后登录时间
			e.Orm.Model(&u).Update("last_login_at", time.Now())
			return u.Id, nil
		}

		// 其他错误
		return 0, err
	}

	// 用户已存在，更新最后登录时间
	err = e.Orm.Model(&exists).Updates(map[string]interface{}{
		"last_login_at":  time.Now(),
		"device_id":      deviceId,
		"device_type":    deviceType,
		"system_version": systemVersion,
		"app_version":    appVersion,
	}).Error

	if err != nil {
		return 0, err
	}

	return exists.Id, nil
}
