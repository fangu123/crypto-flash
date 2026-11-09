package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	fault "go-admin/common/errors"
	"go-admin/common/redishelper"
	"go-admin/common/service"
	"go-admin/common/util"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	coreService "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	adminModels "go-admin/app/admin/models"
)

type AppProfile struct {
	coreService.Service
}

type NotificationResp struct {
	Id        int64     `json:"notificationId"` // 通知ID
	Title     string    `json:"title"`          // 标题
	Content   string    `json:"content"`        // 内容
	EnTitle   string    `json:"enTitle"`        // 标题
	EnContent string    `json:"enContent"`      // 内容
	Type      int64     `json:"type"`           // 类型: 1-系统, 2-交易, 3-安全
	Link      string    `json:"link"`           // 跳转链接
	IsRead    int64     `json:"isRead"`         // 阅读状态: 1-未读, 2-已读
	ReadAt    time.Time `json:"readAt"`         // 阅读时间
	CreatedAt time.Time `json:"createdAt"`      // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`      // 更新时间
}

// UpdateProfile 更新个人资料
func (e *AppProfile) UpdateProfile(userId int, nickName, email, avatar, code, lang string) error {
	updates := map[string]interface{}{}
	if nickName != "" {
		updates["nick_name"] = nickName
	}
	if code != "" {
		// 从Redis获取验证码
		cacheKey := fmt.Sprintf("email_code:%s", email+"_4")
		cachedCode, err := redishelper.DefaultRedis.Get(cacheKey)
		if err != nil || cachedCode == "" {
			return errors.New(fault.GetErrorMessage("please_send_the_verification_code", lang))
		}

		// 验证码校验
		if cachedCode != code {
			return errors.New(fault.GetErrorMessage("invalid_verification_code", lang))
		}
		if email != "" {
			updates["email"] = email
		}
	}

	if avatar != "" {
		updates["head"] = avatar
	}

	return e.Orm.Model(&adminModels.InformationUser{}).
		Where("id = ?", userId).
		Updates(updates).Error
}

// UpdatePassword 检查旧密码并更新新密码
func (e *AppProfile) UpdatePassword(user *adminModels.InformationUser, currentPassword, newPassword, confirmPassword, lang string) error {
	if _, err := pkg.CompareHashAndPassword(user.Password, currentPassword); err != nil {
		return err
	}
	if err := util.ValidatePassword(newPassword, lang); err != nil {
		return err
	}
	if newPassword != confirmPassword {
		return errors.New(fault.GetErrorMessage("new_password_and_confirm_password_do_not_match", lang))
	}
	user.Password = newPassword
	return e.Orm.Session(&gorm.Session{}).
		Model(user).
		Select("password").
		Updates(user).Error
}

func (e *AppProfile) SendEmailCode(email, lang string, t int) error {

	var count int64
	if err := e.Orm.Model(&adminModels.InformationUser{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return err
	}
	// 检查邮箱是否已存在（注册类型）修改邮箱  修改邮箱
	if t == 1 {
		if count > 0 {
			return errors.New(fault.GetErrorMessage("the_email_has_been_registered", lang))
		}
	} else {
		if count == 0 {
			return errors.New(fault.GetErrorMessage("email_does_not_exist", lang))
		}
	}

	// 检查邮箱是否存在（忘记密码类型）
	//if t == 2  || t == 4 {
	//
	//}

	// 生成6位随机验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// 将验证码存入Redis，5分钟过期
	cacheKey := fmt.Sprintf("email_code:%s", email+"_"+strconv.Itoa(t))
	if err := redishelper.DefaultRedis.Set(cacheKey, code, 5*time.Minute); err != nil {
		return err
	}

	// 发送邮件
	subject := fault.GetErrorMessage("verification_code", lang)
	//body := fmt.Sprintf("Your verification code is: %s. This code will expire in 5 minutes.", code)
	body := fmt.Sprintf("%s%s%s", fault.GetErrorMessage("email_information_before", lang), code, fault.GetErrorMessage("email_information_after", lang))
	if err := service.SendTextEmail([]string{email}, subject, body); err != nil {
		return err
	}

	return nil
}

func (e *AppProfile) ForgetPassword(email, password, code, lang string) error {

	// 从Redis获取验证码
	cacheKey := fmt.Sprintf("email_code:%s", email+"_2")
	cachedCode, err := redishelper.DefaultRedis.Get(cacheKey)
	if err != nil || cachedCode == "" {
		return errors.New(fault.GetErrorMessage("please_send_the_verification_code", lang))
	}

	// 验证码校验
	if cachedCode != code {
		return errors.New(fault.GetErrorMessage("invalid_verification_code", lang))
	}

	// 查找用户
	var user adminModels.InformationUser
	if err := e.Orm.Model(&adminModels.InformationUser{}).Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New(fault.GetErrorMessage("user_not_exist", lang))
	}

	// 更新密码
	user.Password = password
	if err := e.Orm.Model(&user).Select("password").Updates(&user).Error; err != nil {
		return err
	}

	// 删除已使用的验证码
	redishelper.DefaultRedis.Delete(cacheKey)
	return nil
}

// Get2FA 获取2FA
func (e *AppProfile) Get2FA(userName, lang string) (string, string, error) {

	// 生成2FA密钥
	secret, err := util.GenerateSecret(32)
	if err != nil {
		return "", "", errors.New(fault.GetErrorMessage("failed_to_generate_2FA_key", lang))
	}

	// 生成备用码
	backupCodes := make([]string, 10)
	for i := 0; i < 10; i++ {
		backupCodes[i] = strconv.Itoa(100000 + i)
	}

	// 生成二维码URL
	config := &util.TOTPConfig{
		Secret:  secret,
		Issuer:  "information",
		Account: userName,
		Digits:  6,
		Period:  30,
	}
	qrCodeURL := util.GenerateQRCodeURL(config)

	return qrCodeURL, secret, nil

}

// Get2FAStatus 获取2FA状态
func (e *AppProfile) Get2FAStatus(email, lang string) (bool, error) {
	// 查询用户信息
	var u adminModels.InformationUser
	if err := e.Orm.Model(&adminModels.InformationUser{}).Where("email = ?", email).First(&u).Error; err != nil {
		return false, errors.New(fault.GetErrorMessage("user_not_exist", lang))
	}
	if u.TwoFactorEnabled == 1 {
		return true, nil
	}

	return false, nil
}

func (e *AppProfile) Enabled2FA(user adminModels.InformationUser, t int, code, lang string) error {

	if user.TwoFactorSecret == "" {
		errors.New(fault.GetErrorMessage("please_binding_2FA_verification", lang))
	}

	if code == "" {
		errors.New(fault.GetErrorMessage("please_enter_the_Google_verification_code", lang))
	}

	// 验证2FA验证码
	valid := util.ValidateTOTP(user.TwoFactorSecret, code, 6, 30, 1)
	if !valid {
		return errors.New(fault.GetErrorMessage("google_verification_code_error", lang))
	}
	if t == 1 { //开启
		//设置two_factor_enabled值
		return e.Orm.Model(&adminModels.InformationUser{}).
			Where("id = ?", user.Id).
			Update("two_factor_enabled", true).Error
	}
	//关闭
	//设置two_factor_enabled值
	return e.Orm.Model(&adminModels.InformationUser{}).
		Where("id = ?", user.Id).
		Update("two_factor_enabled", false).Error

}

func (e *AppProfile) Update2FA(user adminModels.InformationUser, t int, secret, code, ecode, lang string) error {

	//校验验证码
	if code == "" {
		return errors.New(fault.GetErrorMessage("please_enter_the_Google_verification_code", lang))
	}
	if ecode == "" {
		return errors.New(fault.GetErrorMessage("please_enter_the_email_code", lang))
	}
	var cacheKey = ""
	var factorEnabled = "0"
	var s = ""
	if t == 1 { //绑定
		// 验证2FA验证码
		valid := util.ValidateTOTP(secret, code, 6, 30, 1)
		if !valid {
			return errors.New(fault.GetErrorMessage("google_verification_code_error", lang))
		}
		cacheKey = fmt.Sprintf("email_code:%s", user.Email+"_5") //绑定
		factorEnabled = "1"
		s = secret
	} else if t == 2 { //解绑
		// 验证2FA验证码
		valid := util.ValidateTOTP(user.TwoFactorSecret, code, 6, 30, 1)
		if !valid {
			return errors.New(fault.GetErrorMessage("google_verification_code_error", lang))
		}
		cacheKey = fmt.Sprintf("email_code:%s", user.Email+"_6") //解绑
	} else {
		return errors.New(fault.GetErrorMessage("please_enter_the_correct_type", lang))
	}
	// 从Redis获取验证码
	cachedCode, err := redishelper.DefaultRedis.Get(cacheKey)
	if err != nil || cachedCode == "" {
		return errors.New(fault.GetErrorMessage("please_send_the_verification_code", lang))
	}

	// 验证码校验
	if cachedCode != ecode {
		return errors.New(fault.GetErrorMessage("invalid_verification_code", lang))
	}
	// 更新用户信息
	updates := map[string]interface{}{
		"two_factor_enabled": factorEnabled,
		"two_factor_secret":  s,
	}

	if err := e.Orm.Model(&adminModels.InformationUser{}).
		Where("id = ?", user.Id).
		Updates(updates).Error; err != nil {
		return errors.New(fault.GetErrorMessage("failed_to_enable_2FA", lang))
	}
	return nil
}

// GetMessageList 获取消息通知列表
func (e *AppProfile) GetMessageList(userId, page, pageSize int, msgType, isRead int64) ([]NotificationResp, int64, error) {
	// 查询通知地址关联表
	var notificationUser []adminModels.InformationNotificationUser
	var total int64

	query := e.Orm.Model(&adminModels.InformationNotificationUser{}).
		Where("user_id = ?", userId)

	if isRead > 0 {
		query = query.Where("is_read = ?", isRead)
	}

	if msgType > 0 {
		query = query.Where("type = ?", msgType)
	}

	err := query.
		Joins("JOIN information_notification ON information_notification.id = information_notification_user.notification_id").
		Order("information_notification.created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&notificationUser).
		Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	// 提取通知ID
	var notificationIds []int64
	for _, addr := range notificationUser {
		notificationIds = append(notificationIds, addr.NotificationId)
	}

	if len(notificationIds) == 0 {
		return []NotificationResp{}, 0, nil
	}

	// 查询通知详情
	var notifications []adminModels.InformationNotification
	err2 := e.Orm.Where("id IN ?", notificationIds).Find(&notifications).Error
	if err2 != nil {
		return nil, 0, err2
	}

	// 构建返回结果
	var result []NotificationResp
	for _, notification := range notifications {
		// 查找对应的阅读状态
		var notificationAddress *adminModels.InformationNotificationUser
		for _, addr := range notificationUser {
			if addr.NotificationId == int64(notification.Id) {
				notificationAddress = &addr
				break
			}
		}

		result = append(result, NotificationResp{
			Id:        int64(notification.Id),
			Title:     notification.Title,
			Content:   notification.Content,
			EnTitle:   notification.EnTitle,
			EnContent: notification.EnContent,
			Type:      notification.Type,
			Link:      notification.Link,
			IsRead:    notificationAddress.IsRead,
			CreatedAt: notification.CreatedAt,
			UpdatedAt: notification.UpdatedAt,
			ReadAt: func() time.Time {
				if notificationAddress.ReadAt != nil {
					return *notificationAddress.ReadAt
				}
				return time.Time{} // 返回零值时间
			}(),
		})
	}

	return result, total, nil
}

// MarkMessageRead 标记消息为已读
func (e *AppProfile) MarkMessageRead(userId int, notificationId int64, lang string) error {
	// 更新消息为已读
	result := e.Orm.Model(&adminModels.InformationNotificationUser{}).
		Where("user_id = ? AND notification_id  = ?", userId, notificationId).
		Updates(map[string]interface{}{
			"is_read": 1,
			"read_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New(fault.GetErrorMessage("message_not_found", lang))
	}

	return nil
}

// GetMyCollection 获取我的收藏
func (e *AppProfile) GetMyCollection(userId int) (map[string]interface{}, error) {
	// 查询新闻收藏
	var newsFavorites []adminModels.InformationNewsFavorite
	err := e.Orm.Where("user_id = ?", userId).Find(&newsFavorites).Error
	if err != nil {
		return nil, err
	}

	// 查询点位收藏
	var pointFavorites []adminModels.InformationTradingPointFavorite
	err = e.Orm.Where("user_id = ?", userId).Find(&pointFavorites).Error
	if err != nil {
		return nil, err
	}

	// 查询新闻详情
	var newsList []map[string]interface{}
	if len(newsFavorites) > 0 {
		newsIds := make([]int64, len(newsFavorites))
		for i, favorite := range newsFavorites {
			newsIds[i] = favorite.NewsId
		}

		err = e.Orm.Table("information_news").
			Select("id, title,en_title, content,en_content, created_at").
			Where("id IN ?", newsIds).
			Find(&newsList).Error
		if err != nil {
			return nil, err
		}
	}

	// 查询点位详情
	var pointList []map[string]interface{}
	if len(pointFavorites) > 0 {
		pointIds := make([]int64, len(pointFavorites))
		for i, favorite := range pointFavorites {
			pointIds[i] = favorite.TradingPointId
		}

		err = e.Orm.Table("information_trading_point").
			Joins("LEFT JOIN information_coin ON information_coin.id = information_trading_point.coin_id").
			Select("information_trading_point.id, information_coin.symbol, information_trading_point.entry_price, information_trading_point.type as direction, information_trading_point.take_profits as description, information_trading_point.created_at").
			Where("information_trading_point.id IN ?", pointIds).
			Find(&pointList).Error
		if err != nil {
			return nil, err
		}
	}

	result := map[string]interface{}{
		"newsList":  newsList,
		"pointList": pointList,
	}

	return result, nil
}

// GetContact 获取客户联系方式
func (e *AppProfile) GetContact() ([]adminModels.InformationCustomerService, error) {
	var contacts []adminModels.InformationCustomerService
	err := e.Orm.Where("status = ?", 1).
		Order("sort asc").
		Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

// GetAppVersion 获取关于我们版本
func (e *AppProfile) GetAppVersion(platform int64) (*adminModels.InformationAppVersion, error) {
	var version adminModels.InformationAppVersion
	err := e.Orm.Where("platform = ? AND status = ?", platform, 1).
		Order("created_at desc").
		First(&version).Error
	if err != nil {
		return nil, err
	}
	return &version, nil
}

// BindWallet 绑定钱包地址
// @Summary 绑定钱包地址
// @Description 绑定钱包地址
// @Tags 用户钱包
// @Accept application/json
// @Param data body bindWalletReq true "data"
// @Router /api/v1/app/profile/bindWallet [post]
// @Security Bearer
func (e *AppProfile) BindWallet(userId int64, walletAddress, lang string) error {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		return err
	}
	//查看地址是否存在已绑定
	var count int64
	if err := e.Orm.Model(&adminModels.InformationUser{}).Where("wallet_address = ?", walletAddress).Count(&count).Error; err != nil {
		return err
	}

	// 校验钱包地址格式
	if walletAddress == "" {
		return errors.New(fault.GetErrorMessage("wallet_address_cannot_be_empty", lang))
	}

	// 更新用户钱包地址
	if err := e.Orm.Model(&user).Update("wallet_address", walletAddress).Error; err != nil {
		return err
	}

	return nil
}

// GetMessageSetting 获取消息通知设置列表
func (e *AppProfile) GetMessageSetting(userId int, lang string) (map[string]bool, error) {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		return nil, errors.New(fault.GetErrorMessage("user_not_exist", lang))
	}

	// 初始化返回结果
	result := map[string]bool{
		"1": false, // 资讯
		"2": false, // 新点位
		"3": false, // 点位状态
		"4": false, // 系统
	}

	// 如果NotificationSettings不为空，解析JSON
	if user.NotificationSettings != "" {
		// 这里需要根据实际存储的JSON格式进行解析
		// 假设存储的是JSON数组，如["1", "3"]表示开启资讯和点位状态通知
		// 实际实现可能需要根据具体JSON结构调整
		// 这里仅作为示例
		var settings []string
		if err := json.Unmarshal([]byte(user.NotificationSettings), &settings); err == nil {
			for _, setting := range settings {
				if _, ok := result[setting]; ok {
					result[setting] = true
				}
			}
		}
	}

	return result, nil
}

// UpdateMessageSetting 更新消息通知设置
func (e *AppProfile) UpdateMessageSetting(userId, msgType, status int, lang string) error {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		return errors.New(fault.GetErrorMessage("user_not_exist", lang))
	}

	// 验证消息类型
	if msgType < 1 || msgType > 4 {
		return errors.New(fault.GetErrorMessage("invalid_message_type", lang))
	}

	// 验证状态
	if status != 1 && status != 2 {
		return errors.New(fault.GetErrorMessage("invalid_status_value", lang))
	}

	// 获取当前的通知设置
	var settings []string
	if user.NotificationSettings != "" {
		if err := json.Unmarshal([]byte(user.NotificationSettings), &settings); err != nil {
			// 如果解析失败，初始化为空数组
			settings = []string{}
		}
	}

	// 转换为字符串类型
	typeStr := strconv.Itoa(msgType)

	// 根据状态更新设置
	if status == 1 { // 开启
		// 检查是否已经存在
		found := false
		for _, setting := range settings {
			if setting == typeStr {
				found = true
				break
			}
		}
		// 如果不存在，添加到数组
		if !found {
			settings = append(settings, typeStr)
		}
	} else { // 关闭
		// 从数组中移除
		newSettings := []string{}
		for _, setting := range settings {
			if setting != typeStr {
				newSettings = append(newSettings, setting)
			}
		}
		settings = newSettings
	}

	// 将设置转换为JSON字符串
	newSettingsJSON, err := json.Marshal(settings)
	if err != nil {
		return errors.New(fault.GetErrorMessage("update_notification_settings_fail", lang))
	}

	// 更新用户的通知设置
	if err := e.Orm.Model(&user).Update("notification_settings", string(newSettingsJSON)).Error; err != nil {
		return errors.New(fault.GetErrorMessage("update_notification_settings_fail", lang))
	}

	return nil
}

func (e *AppProfile) GenerateToken(userId int, expireTime int) error {
	//设置
	// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
	ctx := context.Background()
	tokenKey := fmt.Sprintf("web_token:%d", userId)

	// 获取Redis客户端
	redisClient := redishelper.DefaultRedis.GetClient()
	//更新tokenKey的过期时间
	redisClient.Expire(ctx, tokenKey, time.Duration(expireTime))
	return nil
}

// GetPersonalInfo 获取个人信息
func (e *AppProfile) GetPersonalInfo(userId int) (map[string]interface{}, error) {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		return nil, err
	}

	// 返回用户个人信息
	result := map[string]interface{}{
		"id":                 user.Id,
		"username":           user.Username,
		"email":              user.Email,
		"head":               user.Head,
		"wallet_address":     user.WalletAddress,
		"language":           user.Language,
		"two_factor_enabled": user.TwoFactorEnabled == 1,
		"status":             user.Status,
		"last_login_at":      user.LastLoginAt,
		"last_login_ip":      user.LastLoginIp,
	}

	return result, nil
}

// UpdatePersonalInfo 更新个人信息
func (e *AppProfile) UpdatePersonalInfo(userId int, username, email, head, walletAddress, emailCode, lang string) error {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		return errors.New(fault.GetErrorMessage("user_not_exist", lang))
	}

	// 准备更新数据
	updates := map[string]interface{}{}

	// 如果提供了邮箱验证码，需要验证邮箱
	if emailCode != "" {
		// 从Redis获取验证码
		cacheKey := fmt.Sprintf("email_code:%s", email+"_4")
		cachedCode, err := redishelper.DefaultRedis.Get(cacheKey)
		if err != nil || cachedCode == "" {
			return errors.New(fault.GetErrorMessage("please_send_the_verification_code", lang))
		}

		// 验证码校验
		if cachedCode != emailCode {
			return errors.New(fault.GetErrorMessage("invalid_verification_code", lang))
		}

		// 更新邮箱
		if email != "" {
			updates["email"] = email
		}
	} else if email != "" {
		// 没有提供邮箱验证码但尝试修改邮箱
		return errors.New(fault.GetErrorMessage("email_code_is_required", lang))
	}

	// 更新用户名
	if username != "" {
		updates["username"] = username
	}

	// 更新头像
	if head != "" {
		updates["head"] = head
	}

	// 更新钱包地址
	if walletAddress != "" {
		// 检查钱包地址是否已被其他用户使用
		var count int64
		if err := e.Orm.Model(&adminModels.InformationUser{}).Where("wallet_address = ? AND id != ?", walletAddress, userId).Count(&count).Error; err != nil {
			return err
		}

		if count > 0 {
			return errors.New(fault.GetErrorMessage("wallet_address_already_exists", lang))
		}

		updates["wallet_address"] = walletAddress
	}

	// 执行更新
	if len(updates) > 0 {
		if err := e.Orm.Model(&adminModels.InformationUser{}).Where("id = ?", userId).Updates(updates).Error; err != nil {
			return err
		}
	}

	return nil
}
