package service

import (
	"errors"
	"go-admin/app/admin/models"
	fault "go-admin/common/errors"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
)

// BindService 处理账号绑定相关服务
type BindService struct {
	service.Service
}

// BindGoogle 绑定Google账号
func (s *BindService) BindGoogle(userID int, googleId, googleEmail, googleName, googlePicture, googleAccessToken, lang string) error {

	// 检查该Google账号是否已被其他用户绑定
	var count int64
	err := s.Orm.Model(&models.InformationUser{}).
		Where("google_id = ? and id != ?", googleId, userID).
		Count(&count).Error
	if err != nil {
		return errors.New(fault.GetErrorMessage("check_account_binding_status_failed", lang))
	}

	if count > 0 {
		return errors.New(fault.GetErrorMessage("google_account_has_been_bound_by_other_users", lang))
	}

	// 更新用户信息，绑定Google账号
	err2 := s.Orm.Model(&models.InformationUser{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"google_id":           googleId,
			"google_email":        googleEmail,
			"google_name":         googleName,
			"google_picture":      googlePicture,
			"google_access_token": googleAccessToken,
			"google_expires_at":   time.Now().Add(time.Hour * 1), // 假设令牌1小时后过期
			"last_login_at":       time.Now(),
		}).Error

	if err2 != nil {
		return errors.New(fault.GetErrorMessage("bind_google_account_failed", lang))
	}

	return nil
}

// BindMicrosoft 绑定Microsoft账号
func (s *BindService) BindMicrosoft(userID int, microsoftId, microsoftEmail, microsoftName, microsoftPicture, microsoftAccessToken, lang string) error {

	// 检查该Microsoft账号是否已被其他用户绑定
	var count int64
	err := s.Orm.Model(&models.InformationUser{}).
		Where("microsoft_id = ? and id != ?", microsoftId, userID).
		Count(&count).Error
	if err != nil {
		return errors.New(fault.GetErrorMessage("check_account_binding_status_failed", lang))
	}

	if count > 0 {
		return errors.New(fault.GetErrorMessage("microsoft_account_has_been_bound_by_other_users", lang))
	}

	// 更新用户信息，绑定Microsoft账号
	err2 := s.Orm.Model(&models.InformationUser{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"microsoft_id":           microsoftId,
			"microsoft_email":        microsoftEmail,
			"microsoft_name":         microsoftName,
			"microsoft_picture":      microsoftPicture,
			"microsoft_access_token": microsoftAccessToken,
			"microsoft_expires_at":   time.Now().Add(time.Hour * 1), // 假设令牌1小时后过期
			"last_login_at":          time.Now(),
		}).Error

	if err2 != nil {
		return errors.New(fault.GetErrorMessage("bind_microsoft_account_failed", lang))
	}

	return nil
}

// UnbindGoogle 解绑Google账号
func (s *BindService) UnbindGoogle(userID int, lang string) error {
	// 检查用户是否已绑定Google账号
	var hasGoogle bool
	err := s.Orm.Model(&models.InformationUser{}).
		Where("id = ? and google_id != ''", userID).
		Select("count(*)").Scan(&hasGoogle).Error
	if err != nil {
		return errors.New(fault.GetErrorMessage("check_account_binding_status_failed", lang))
	}

	if !hasGoogle {
		return errors.New(fault.GetErrorMessage("not_bound_google_account", lang))
	}

	// 更新用户信息，解绑Google账号
	err2 := s.Orm.Model(&models.InformationUser{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"google_id":           "",
			"google_email":        "",
			"google_name":         "",
			"google_picture":      "",
			"google_access_token": "",
			"google_expires_at":   time.Time{},
		}).Error

	if err2 != nil {
		return errors.New(fault.GetErrorMessage("unbind_google_account_failed", lang))
	}

	return nil
}

// UnbindMicrosoft 解绑Microsoft账号
func (s *BindService) UnbindMicrosoft(userID int, lang string) error {

	// 检查用户是否已绑定Microsoft账号
	var hasMicrosoft bool
	err := s.Orm.Model(&models.InformationUser{}).
		Where("id = ? and microsoft_id != ''", userID).
		Select("count(*)").Scan(&hasMicrosoft).Error
	if err != nil {
		return errors.New(fault.GetErrorMessage("check_account_binding_status_failed", lang))
	}

	if !hasMicrosoft {
		return errors.New(fault.GetErrorMessage("not_bound_microsoft_account", lang))
	}

	// 更新用户信息，解绑Microsoft账号
	err2 := s.Orm.Model(&models.InformationUser{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"microsoft_id":           "",
			"microsoft_email":        "",
			"microsoft_name":         "",
			"microsoft_picture":      "",
			"microsoft_access_token": "",
			"microsoft_expires_at":   time.Time{},
		}).Error

	if err2 != nil {
		return errors.New(fault.GetErrorMessage("unbind_microsoft_account_failed", lang))
	}

	return nil
}

// GetBindStatus 获取账号绑定状态
func (s *BindService) GetBindStatus(userID int, lang string) (map[string]interface{}, error) {
	// 获取用户信息
	var user models.InformationUser
	err := s.Orm.Model(&models.InformationUser{}).
		Where("id = ?", userID).
		Select("google_id, microsoft_id, email, username, head").
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		return nil, errors.New(fault.GetErrorMessage("get_user_info_failed", lang))
	}

	// 返回绑定状态
	return map[string]interface{}{
		"googleBound":    user.GoogleId != "",
		"microsoftBound": user.MicrosoftId != "",
	}, nil
}
