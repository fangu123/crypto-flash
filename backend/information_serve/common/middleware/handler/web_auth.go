package handler

import (
	"context"
	"errors"
	"fmt"
	fault "go-admin/common/errors"
	"go-admin/common/redishelper"
	"go-admin/common/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"gorm.io/gorm"
)

type WebLogin struct {
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	TwoFactorCode string `json:"twoFactorCode" binding:"omitempty"`
	IsRemember    bool   `json:"isRemember" binding:"omitempty"`
}

type WebLoginCaptcha struct {
	Email         string `json:"email" binding:"required"`
	Code          string `json:"code" binding:"required"`
	TwoFactorCode string `json:"twoFactorCode" binding:"omitempty"`
	IsRemember    bool   `json:"isRemember" binding:"omitempty"`
}

type WebLoginGoogle struct {
	GoogleId          string `json:"googleId" binding:"required"`
	GoogleEmail       string `json:"googleEmail" binding:"required"`
	GoogleName        string `json:"googleName" binding:"required"`
	GooglePicture     string `json:"googlePicture" binding:"required"`
	GoogleAccessToken string `json:"googleAccessToken" binding:"required"`
	//DeviceId          string `json:"deviceId" binding:"required"`
	//DeviceType        int64  `json:"deviceType" binding:"required"`
	//SystemVersion     string `json:"systemVersion" binding:"required"`
	//AppVersion        string `json:"appVersion" binding:"required"`
	IsRemember bool `json:"isRemember" binding:"omitempty"`
}

type WebLoginMicrosoft struct {
	MicrosoftId          string `json:"microsoftId" binding:"required"`
	MicrosoftEmail       string `json:"microsoftEmail" binding:"required"`
	MicrosoftName        string `json:"microsoftName" binding:"required"`
	MicrosoftPicture     string `json:"microsoftPicture" binding:"required"`
	MicrosoftAccessToken string `json:"microsoftAccessToken" binding:"required"`
	//DeviceId             string `json:"deviceId" binding:"required"`
	//DeviceType           int64  `json:"deviceType" binding:"required"`
	//SystemVersion        string `json:"systemVersion" binding:"required"`
	//AppVersion           string `json:"appVersion" binding:"required"`
	IsRemember bool `json:"isRemember" binding:"omitempty"`
}

// WebAuthenticatorEmail 用邮箱+密码获取token（无验证码）
func WebAuthenticatorEmail(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}

	var loginVals WebLogin
	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	var u User
	err = db.Table("information_user").Where("email = ? and status = 1", strings.ToLower(loginVals.Email)).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, jwt.ErrFailedAuthentication
		}
		log.Errorf("get user error, %s", err.Error())
		return nil, jwt.ErrFailedAuthentication
	}

	if _, err := pkg.CompareHashAndPassword(u.Password, loginVals.Password); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	// 检查是否启用了2FA
	if u.TwoFactorEnabled == 1 {
		// 如果用户启用了2FA，必须提供验证码
		if loginVals.TwoFactorCode == "" {
			return nil, jwt.ErrFailedAuthentication
		}

		// 验证2FA验证码
		valid := util.ValidateTOTP(u.TwoFactorSecret, loginVals.TwoFactorCode, 6, 30, 1)
		if !valid {
			return nil, jwt.ErrFailedAuthentication
		}
	}

	// 生成新的token ID
	tokenId := fmt.Sprintf("%d_%d", u.Id, time.Now().UnixNano())

	// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
	ctx := context.Background()
	tokenKey := fmt.Sprintf("web_token:%d", u.Id)

	// 获取Redis客户端
	redisClient := redishelper.DefaultRedis.GetClient()

	// 存储新的token ID，过期时间设置为24小时
	if err := redisClient.Set(ctx, tokenKey, tokenId, 24*time.Hour).Err(); err != nil {
		log.Errorf("store token to redis error, %s", err.Error())
		// 不影响登录流程，继续执行
	}

	return map[string]interface{}{"user": u, "tokenId": tokenId}, nil
}

// WebAuthenticatorEmailWith2FA 用邮箱+密码获取token（强制2FA验证）
func WebAuthenticatorEmailWith2FA(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var loginVals WebLogin
	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, errors.New(fault.GetErrorMessage("parameter_error", lang))
	}

	var u User
	err = db.Table("information_user").Where("email = ? and status = 1", strings.ToLower(loginVals.Email)).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		log.Errorf("get user error, %s", err.Error())
		return nil, errors.New(fault.GetErrorMessage("user_not_exist", lang))
	}

	if _, err := pkg.CompareHashAndPassword(u.Password, loginVals.Password); err != nil {
		return nil, errors.New(fault.GetErrorMessage("incorrect_username_or_password", lang))
	}

	// 启用了2FA
	if u.TwoFactorEnabled == 1 {
		// 如果用户启用了2FA，必须提供验证码
		if loginVals.TwoFactorCode == "" {
			return nil, errors.New(fault.GetErrorMessage("please_enter_the_Google_verification_code", lang))
		}

		// 验证2FA验证码
		valid := util.ValidateTOTP(u.TwoFactorSecret, loginVals.TwoFactorCode, 6, 30, 1)
		if !valid {
			return nil, errors.New(fault.GetErrorMessage("google_verification_code_error", lang))
		}
	}

	// 生成新的token ID
	tokenId := fmt.Sprintf("%d_%d", u.Id, time.Now().UnixNano())

	// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
	ctx := context.Background()
	tokenKey := fmt.Sprintf("web_token:%d", u.Id)

	// 获取Redis客户端
	redisClient := redishelper.DefaultRedis.GetClient()
	var t = 24 * time.Hour
	if loginVals.IsRemember {
		// 存储新的token ID，过期时间设置为30天
		t = 30 * 24 * time.Hour
	}

	// 存储新的token ID，过期时间设置为24小时
	if err := redisClient.Set(ctx, tokenKey, tokenId, t).Err(); err != nil {
		log.Errorf("store token to redis error, %s", err.Error())
		// 不影响登录流程，继续执行
	}

	return map[string]interface{}{"user": u, "tokenId": tokenId}, nil
}

// WebAuthenticatorCaptcha 用邮箱+验证码获取token（如果开启2FA就需要输入谷歌验证码）
func WebAuthenticatorCaptcha(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var loginVals WebLoginCaptcha
	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, errors.New(fault.GetErrorMessage("parameter_error", lang))
	}

	var u User
	err = db.Table("information_user").Where("email = ? and status = 1", strings.ToLower(loginVals.Email)).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		log.Errorf("get user error, %s", err.Error())
		return nil, errors.New(fault.GetErrorMessage("user_not_exist", lang))
	}

	// 从Redis获取验证码
	cacheKey := fmt.Sprintf("email_code:%s", strings.ToLower(loginVals.Email)+"_3") // 3代表登录类型
	cachedCode, er := redishelper.DefaultRedis.Get(cacheKey)
	if er != nil || cachedCode == "" {
		return nil, errors.New(fault.GetErrorMessage("please_send_the_verification_code", lang))
	}

	// 验证码校验
	if cachedCode != loginVals.Code {
		return nil, errors.New(fault.GetErrorMessage("invalid_verification_code", lang))
	}

	// 启用了2FA
	if u.TwoFactorEnabled == 1 {
		// 如果用户启用了2FA，必须提供验证码
		if loginVals.TwoFactorCode == "" {
			return nil, errors.New(fault.GetErrorMessage("please_enter_the_Google_verification_code", lang))
		}

		// 验证2FA验证码
		valid := util.ValidateTOTP(u.TwoFactorSecret, loginVals.TwoFactorCode, 6, 30, 1)
		if !valid {
			return nil, errors.New(fault.GetErrorMessage("google_verification_code_error", lang))
		}
	}

	// 生成新的token ID
	tokenId := fmt.Sprintf("%d_%d", u.Id, time.Now().UnixNano())

	// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
	ctx := context.Background()
	tokenKey := fmt.Sprintf("web_token:%d", u.Id)

	// 获取Redis客户端
	redisClient := redishelper.DefaultRedis.GetClient()

	var t = 24 * time.Hour
	if loginVals.IsRemember {
		// 存储新的token ID，过期时间设置为30天
		t = 30 * 24 * time.Hour
	}

	// 存储新的token ID，过期时间设置为24小时
	if err := redisClient.Set(ctx, tokenKey, tokenId, t).Err(); err != nil {
		log.Errorf("store token to redis error, %s", err.Error())
		// 不影响登录流程，继续执行
	}

	// 删除已使用的验证码
	redishelper.DefaultRedis.Delete(cacheKey)

	return map[string]interface{}{"user": u, "tokenId": tokenId}, nil
}

// WebAuthenticatorWallet 用钱包地址获取token，如果用户不存在则自动注册
func WebAuthenticatorWallet(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var walletLoginReq struct {
		WalletAddress string `json:"walletAddress" binding:"required"`
		BundleId      string `json:"chainId" binding:"required"`
		//DeviceType    int64  `json:"deviceType" binding:"required"`
		//SystemVersion string `json:"systemVersion" binding:"required"`
		//AppVersion    string `json:"appVersion" binding:"required"`
		IsRemember bool `json:"isRemember"`
	}

	if err := c.ShouldBindJSON(&walletLoginReq); err != nil {
		return nil, errors.New(fault.GetErrorMessage("parameter_error", lang))
	}

	// 钱包地址格式校验（简单的以太坊地址格式校验）
	if len(walletLoginReq.WalletAddress) < 20 || len(walletLoginReq.WalletAddress) > 50 {
		return nil, errors.New(fault.GetErrorMessage("invalid_wallet_address", lang))
	}

	// 查询用户是否已存在
	var u User
	err = db.Table("information_user").Where("wallet_address = ? and status = 1", walletLoginReq.WalletAddress).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新用户
			newUser := User{
				Username:                  walletLoginReq.WalletAddress,
				Email:                     "",
				WalletAddress:             walletLoginReq.WalletAddress,
				Status:                    1,
				Password:                  "",
				PointNotificationSettings: 1,
				TwoFactorEnabled:          2,
				Language:                  lang,
				BundleId:                  walletLoginReq.BundleId,
				//DeviceId:                  walletLoginReq.DeviceId,
				//SystemVersion:             walletLoginReq.SystemVersion,
				//AppVersion:                walletLoginReq.AppVersion,
				//DeviceType:                walletLoginReq.DeviceType,
				NotificationSettings: `["1", "2", "3", "4"]`,
				LastLoginAt:          time.Now(),
			}

			if err := db.Table("information_user").Create(&newUser).Error; err != nil {
				return nil, errors.New(fault.GetErrorMessage("user_creation_failed", lang))
			}

			// 更新最后登录时间
			db.Table("information_user").Model(&newUser).Update("last_login_at", time.Now())

			// 生成新的token ID
			tokenId := fmt.Sprintf("%d_%d", newUser.Id, time.Now().UnixNano())

			// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
			ctx := context.Background()
			tokenKey := fmt.Sprintf("web_token:%d", newUser.Id)

			// 获取Redis客户端
			redisClient := redishelper.DefaultRedis.GetClient()

			var t = 24 * time.Hour
			if walletLoginReq.IsRemember {
				// 存储新的token ID，过期时间设置为30天
				t = 30 * 24 * time.Hour
			}

			// 存储新的token ID，过期时间设置为24小时
			if err := redisClient.Set(ctx, tokenKey, tokenId, t).Err(); err != nil {
				log.Errorf("store token to redis error, %s", err.Error())
				// 不影响登录流程，继续执行
			}

			return map[string]interface{}{"user": newUser, "tokenId": tokenId}, nil
		}

		// 其他错误
		return nil, jwt.ErrFailedAuthentication
	}

	// 用户已存在，更新最后登录时间
	err = db.Table("information_user").Model(&u).Updates(map[string]interface{}{
		"last_login_at": time.Now(),
		//"device_id":      walletLoginReq.DeviceId,
		//"device_type":    walletLoginReq.DeviceType,
		//"system_version": walletLoginReq.SystemVersion,
		//"app_version":    walletLoginReq.AppVersion,
	}).Error

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	// 生成新的token ID
	tokenId := fmt.Sprintf("%d_%d", u.Id, time.Now().UnixNano())

	// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
	ctx := context.Background()
	tokenKey := fmt.Sprintf("web_token:%d", u.Id)

	// 获取Redis客户端
	redisClient := redishelper.DefaultRedis.GetClient()

	var t = 24 * time.Hour
	if walletLoginReq.IsRemember {
		// 存储新的token ID，过期时间设置为30天
		t = 30 * 24 * time.Hour
	}

	// 存储新的token ID，过期时间设置为24小时
	if err := redisClient.Set(ctx, tokenKey, tokenId, t).Err(); err != nil {
		log.Errorf("store token to redis error, %s", err.Error())
		// 不影响登录流程，继续执行
	}

	return map[string]interface{}{"user": u, "tokenId": tokenId}, nil
}

// WebAuthenticatorGoogle 用Google账号获取token，如果用户不存在则自动注册
func WebAuthenticatorGoogle(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var loginVals WebLoginGoogle
	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, errors.New(fault.GetErrorMessage("parameter_error", lang))
	}

	// 查询用户是否已存在
	var u User
	err = db.Table("information_user").Where("google_id = ?", loginVals.GoogleId).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新用户
			newUser := User{
				Username:                  loginVals.GoogleEmail,
				Email:                     loginVals.GoogleEmail,
				Head:                      loginVals.GooglePicture,
				Status:                    1,
				PointNotificationSettings: 1,
				TwoFactorEnabled:          2, // 默认不启用2FA
				Language:                  lang,
				//DeviceId:                  loginVals.DeviceId,
				//SystemVersion:             loginVals.SystemVersion,
				//AppVersion:                loginVals.AppVersion,
				//DeviceType:                loginVals.DeviceType,
				NotificationSettings: `["1", "2", "3", "4"]`,
				LastLoginAt:          time.Now(),
				GoogleId:             loginVals.GoogleId,
				GoogleEmail:          loginVals.GoogleEmail,
				GoogleName:           loginVals.GoogleName,
				GooglePicture:        loginVals.GooglePicture,
				GoogleAccessToken:    loginVals.GoogleAccessToken,
				GoogleExpiresAt:      time.Now().Add(time.Hour * 1), // 假设令牌1小时后过期
			}

			if err := db.Table("information_user").Create(&newUser).Error; err != nil {
				return nil, errors.New(fault.GetErrorMessage("user_creation_failed", lang))
			}

			// 更新最后登录时间
			db.Table("information_user").Model(&newUser).Update("last_login_at", time.Now())

			// 生成新的token ID
			tokenId := fmt.Sprintf("%d_%d", newUser.Id, time.Now().UnixNano())

			// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
			ctx := context.Background()
			tokenKey := fmt.Sprintf("web_token:%d", newUser.Id)

			// 获取Redis客户端
			redisClient := redishelper.DefaultRedis.GetClient()

			var t = 24 * time.Hour
			if loginVals.IsRemember {
				// 存储新的token ID，过期时间设置为30天
				t = 30 * 24 * time.Hour
			}

			// 存储新的token ID，过期时间设置为24小时
			if err := redisClient.Set(ctx, tokenKey, tokenId, t).Err(); err != nil {
				log.Errorf("store token to redis error, %s", err.Error())
				// 不影响登录流程，继续执行
			}

			return map[string]interface{}{"user": newUser, "tokenId": tokenId}, nil
		}

		// 其他错误
		return nil, jwt.ErrFailedAuthentication
	}

	// 用户已存在，更新最后登录时间和Google信息
	err = db.Table("information_user").Model(&u).Updates(map[string]interface{}{
		"last_login_at": time.Now(),
		//"device_id":           loginVals.DeviceId,
		//"device_type":         loginVals.DeviceType,
		//"system_version":      loginVals.SystemVersion,
		//"app_version":         loginVals.AppVersion,
		"google_id":           loginVals.GoogleId,
		"google_email":        loginVals.GoogleEmail,
		"google_name":         loginVals.GoogleName,
		"google_picture":      loginVals.GooglePicture,
		"google_access_token": loginVals.GoogleAccessToken,
		"google_expires_at":   time.Now().Add(time.Hour * 1), // 假设令牌1小时后过期
	}).Error

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	// 生成新的token ID
	tokenId := fmt.Sprintf("%d_%d", u.Id, time.Now().UnixNano())

	// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
	ctx := context.Background()
	tokenKey := fmt.Sprintf("web_token:%d", u.Id)

	// 获取Redis客户端
	redisClient := redishelper.DefaultRedis.GetClient()

	var t = 24 * time.Hour
	if loginVals.IsRemember {
		// 存储新的token ID，过期时间设置为30天
		t = 30 * 24 * time.Hour
	}

	// 存储新的token ID，过期时间设置为24小时
	if err := redisClient.Set(ctx, tokenKey, tokenId, t).Err(); err != nil {
		log.Errorf("store token to redis error, %s", err.Error())
		// 不影响登录流程，继续执行
	}

	return map[string]interface{}{"user": u, "tokenId": tokenId}, nil
}

// WebAuthenticatorMicrosoft 用Microsoft账号获取token，如果用户不存在则自动注册
func WebAuthenticatorMicrosoft(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var loginVals WebLoginMicrosoft
	if err := c.ShouldBind(&loginVals); err != nil {
		return nil, errors.New(fault.GetErrorMessage("parameter_error", lang))
	}

	// 查询用户是否已存在
	var u User
	err = db.Table("information_user").Where("microsoft_id = ?", loginVals.MicrosoftId).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新用户
			newUser := User{
				Username:                  loginVals.MicrosoftEmail,
				Email:                     loginVals.MicrosoftEmail,
				Head:                      loginVals.MicrosoftPicture,
				Status:                    1,
				PointNotificationSettings: 1,
				TwoFactorEnabled:          2, // 默认不启用2FA
				Language:                  lang,
				//DeviceId:                  loginVals.DeviceId,
				//SystemVersion:             loginVals.SystemVersion,
				//AppVersion:                loginVals.AppVersion,
				//DeviceType:                loginVals.DeviceType,
				NotificationSettings: `["1", "2", "3", "4"]`,
				LastLoginAt:          time.Now(),
				MicrosoftId:          loginVals.MicrosoftId,
				MicrosoftEmail:       loginVals.MicrosoftEmail,
				MicrosoftName:        loginVals.MicrosoftName,
				MicrosoftPicture:     loginVals.MicrosoftPicture,
				MicrosoftAccessToken: loginVals.MicrosoftAccessToken,
				MicrosoftExpiresAt:   time.Now().Add(time.Hour * 1), // 假设令牌1小时后过期
			}

			if err := db.Table("information_user").Create(&newUser).Error; err != nil {
				return nil, errors.New(fault.GetErrorMessage("user_creation_failed", lang))
			}

			// 更新最后登录时间
			db.Table("information_user").Model(&newUser).Update("last_login_at", time.Now())

			// 生成新的token ID
			tokenId := fmt.Sprintf("%d_%d", newUser.Id, time.Now().UnixNano())

			// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
			ctx := context.Background()
			tokenKey := fmt.Sprintf("web_token:%d", newUser.Id)

			// 获取Redis客户端
			redisClient := redishelper.DefaultRedis.GetClient()

			var t = 24 * time.Hour
			if loginVals.IsRemember {
				// 存储新的token ID，过期时间设置为30天
				t = 30 * 24 * time.Hour
			}

			// 存储新的token ID，过期时间设置为24小时
			if err := redisClient.Set(ctx, tokenKey, tokenId, t).Err(); err != nil {
				log.Errorf("store token to redis error, %s", err.Error())
				// 不影响登录流程，继续执行
			}

			return map[string]interface{}{"user": newUser, "tokenId": tokenId}, nil
		}

		// 其他错误
		return nil, jwt.ErrFailedAuthentication
	}

	// 用户已存在，更新最后登录时间和Microsoft信息
	err = db.Table("information_user").Model(&u).Updates(map[string]interface{}{
		"last_login_at": time.Now(),
		//"device_id":              loginVals.DeviceId,
		//"device_type":            loginVals.DeviceType,
		//"system_version":         loginVals.SystemVersion,
		//"app_version":            loginVals.AppVersion,
		"microsoft_id":           loginVals.MicrosoftId,
		"microsoft_email":        loginVals.MicrosoftEmail,
		"microsoft_name":         loginVals.MicrosoftName,
		"microsoft_picture":      loginVals.MicrosoftPicture,
		"microsoft_access_token": loginVals.MicrosoftAccessToken,
		"microsoft_expires_at":   time.Now().Add(time.Hour * 1), // 假设令牌1小时后过期
	}).Error

	if err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	// 生成新的token ID
	tokenId := fmt.Sprintf("%d_%d", u.Id, time.Now().UnixNano())

	// 将新的token ID存储到Redis，设置过期时间为JWT的过期时间
	ctx := context.Background()
	tokenKey := fmt.Sprintf("web_token:%d", u.Id)

	// 获取Redis客户端
	redisClient := redishelper.DefaultRedis.GetClient()

	var t = 24 * time.Hour
	if loginVals.IsRemember {
		// 存储新的token ID，过期时间设置为30天
		t = 30 * 24 * time.Hour
	}

	// 存储新的token ID，过期时间设置为24小时
	if err := redisClient.Set(ctx, tokenKey, tokenId, t).Err(); err != nil {
		log.Errorf("store token to redis error, %s", err.Error())
		// 不影响登录流程，继续执行
	}

	return map[string]interface{}{"user": u, "tokenId": tokenId}, nil
}
