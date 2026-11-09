package web

import (
	"errors"
	fault "go-admin/common/errors"
	"strconv"

	"github.com/gin-gonic/gin"

	webService "go-admin/app/other/service"
	"go-admin/common/util"
)

type Profile struct {
	AppApi
}

// GetProfile 个人信息
// @Summary Web用户端-个人信息
// @Tags Web用户端
// @Router /api/v1/web/profile [get]
// @Security Bearer
func (e Profile) GetProfile(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	e.OK(gin.H{
		"userId":        u.Id,
		"username":      u.Username,
		"email":         u.Email,
		"avatar":        u.Head,
		"factorSecret":  u.TwoFactorSecret,
		"factorEnabled": u.TwoFactorEnabled,
	}, fault.GetErrorMessage("acquire_success", lang))
}

type updateProfileReq struct {
	NickName  string `json:"nickName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Remark    string `json:"remark"`
	EmailCode string `json:"emailCode"`
}

// UpdateProfile 更新个人信息（姓名/邮箱/头像/收款类型）
// @Summary Web用户端-更新个人信息
// @Tags Web用户端
// @Accept application/json
// @Param data body updateProfileReq true "data"
// @Router /api/v1/web/profile [put]
// @Security Bearer
func (e Profile) UpdateProfile(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req updateProfileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if req.Email != "" && !emailRe.MatchString(req.Email) {
		e.Error(400, errors.New(""), fault.GetErrorMessage("the_email_format_is_incorrect", lang))
		return
	}

	if err := s.UpdateProfile(u.Id, req.NickName, req.Email, req.Avatar, req.EmailCode, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("update_successfully", lang))
}

type updatePasswordReq struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

// UpdatePassword 修改密码
// @Summary Web用户端-修改密码
// @Tags Web用户端
// @Accept application/json
// @Param data body updatePasswordReq true "data"
// @Router /api/v1/web/password [put]
// @Security Bearer
func (e Profile) UpdatePassword(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req updatePasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}
	if err := util.ValidatePassword(req.CurrentPassword, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	if err := util.ValidatePassword(req.NewPassword, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.UpdatePassword(&u, req.CurrentPassword, req.NewPassword, req.ConfirmPassword, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("update_successfully", lang))
}

type sendEmailCodeReq struct {
	Email string `json:"email" binding:"required"`
	Type  int    `json:"type" binding:"required"` // 1:注册 2:忘记密码 3:登录 4:绑定 5:解绑
}

// SendEmailCode 发送邮箱验证码
// @Summary Web用户端-发送邮箱验证码
// @Tags Web用户端
// @Accept application/json
// @Param data body sendEmailCodeReq true "data"
// @Router /api/v1/profile/sendEmailCode [post]
func (e Profile) SendEmailCode(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var req sendEmailCodeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	// 校验邮箱格式
	//req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if !emailRe.MatchString(req.Email) {
		e.Error(400, errors.New(""), fault.GetErrorMessage("the_email_format_is_incorrect", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.SendEmailCode(req.Email, lang, req.Type); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("verification_code_sent_successfully", lang))
}

type forgetPasswordReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

// ForgetPassword 忘记密码
// @Summary Web用户端-忘记密码
// @Tags Web用户端
// @Accept application/json
// @Param data body forgetPasswordReq true "data"
// @Router /api/v1/profile/forgetPassword [post]
func (e Profile) ForgetPassword(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var req forgetPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	// 校验邮箱格式
	//req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if !emailRe.MatchString(req.Email) {
		e.Error(400, errors.New(""), fault.GetErrorMessage("the_email_format_is_incorrect", lang))
		return
	}

	// 校验密码格式
	if err := util.ValidatePassword(req.Password, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.ForgetPassword(req.Email, req.Password, req.Code, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("password_reset_successfully", lang))
}

// Get2FA 获取2FA
// @Summary 获取2FA
// @Description 获取2FA
// @Tags 用户2FA
// @Accept application/json
// @Param data body faReq true "data"
// @Router /api/v1/web/profile/get2fa [get]
// @Security Bearer
func (e Profile) Get2FA(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	url, secret, err := s.Get2FA(user.Username, lang)
	if err != nil {
		e.Error(400, err, "failed to get 2FA")
		return
	}

	e.OK(gin.H{"url": url, "secret": secret}, fault.GetErrorMessage("acquire_success", lang))

}

// GetMessageList 获取消息通知列表
// @Summary 获取消息通知列表
// @Description 获取消息通知列表
// @Tags 用户消息
// @Router /api/v1/app/profile/messageList [get]
// @Security Bearer
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页数量" default(10)
// @Param type query int false "消息类型：1-系统通知，2-点位提醒，3-快讯推送， 4-点位状态"
// @Param isRead query int false "是否已读：1-已读，2-未读"
func (e Profile) GetMessageList(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	// 获取分页参数
	page := 1
	pageSize := 10
	if v := c.Query("page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			page = n
		}
	}
	if v := c.Query("size"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 100 {
			pageSize = n
		}
	}

	// 获取筛选参数
	msgType := int64(0) // 0表示不筛选
	if v := c.Query("type"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 1 && n <= 3 {
			msgType = int64(n)
		}
	}

	isRead := int64(0) // 0表示不筛选
	if v := c.Query("isRead"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && (n == 1 || n == 2) {
			isRead = int64(n)
		}
	}

	result, total, err := s.GetMessageList(user.Id, page, pageSize, msgType, isRead)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_message_list_fail", lang))
		return
	}

	e.OK(gin.H{
		"list":     result,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}, fault.GetErrorMessage("acquire_success", lang))
}

// MarkMessageRead 标记消息为已读
// @Summary 标记消息为已读
// @Description 标记消息为已读
// @Tags 用户消息
// @Accept application/json
// @Param data body markMessageReadReq true "data"
// @Router /api/v1/app/profile/markMessageRead [post]
// @Security Bearer
func (e Profile) MarkMessageRead(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 通知ID
	notificationId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("message_id_invalid", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.MarkMessageRead(user.Id, notificationId, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("mark_success", lang))
}

// GetMyCollection 获取我的收藏
// @Summary 获取我的收藏
// @Description 获取我的收藏
// @Tags 用户收藏
// @Router /api/v1/app/profile/myCollection [get]
// @Security Bearer
func (e Profile) GetMyCollection(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetMyCollection(user.Id)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_collection_list_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// GetContact 获取客户联系方式
// @Summary 获取客户联系方式
// @Description 获取客户联系方式
// @Tags 客服
// @Router /api/v1/app/profile/contact [get]
func (e Profile) GetContact(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	//_, err = e.currentUser(c)
	//if err != nil {
	//	e.Error(401, err, fault.GetErrorMessage("not_login", lang))
	//	return
	//}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetContact()
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_contact_information_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// GetAppVersion 获取关于我们版本
// @Summary 获取关于我们版本
// @Description 获取关于我们版本
// @Tags 关于我们
// @Param platform query int true "平台：1-iOS，2-Android" default(1)
// @Router /api/v1/app/profile/appVersion [get]
func (e Profile) GetAppVersion(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	//_, err = e.currentUser(c)
	//if err != nil {
	//	e.Error(401, err, fault.GetErrorMessage("not_login", lang))
	//	return
	//}

	platform := int64(1) // 默认iOS
	if v := c.Query("platform"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && (n == 1 || n == 2) {
			platform = int64(n)
		}
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetAppVersion(platform)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_version_information_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

type enabled2FAReq struct {
	Type          int    `json:"type" binding:"required"`
	TwoFactorCode string `json:"twoFactorCode" binding:"required"`
}

// Enabled2FA 设置2FA状态
// @Summary 设置2FA状态
// @Description 设置2FA状态
// @Tags 用户2FA
// @Accept application/json
// @Param data body faReq true "data"
// @Router /api/v1/web/profile/enabled2fa [post]
// @Security Bearer
func (e Profile) Enabled2FA(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req enabled2FAReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	err = s.Enabled2FA(u, req.Type, req.TwoFactorCode, lang)
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("success_to_set_2FA", lang))

}

// Query2FA 获取2FA状态
// @Summary 获取2FA状态
// @Description 获取2FA状态
// @Tags 用户2FA
// @Param email query string true "用户邮箱"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/profile/query2fa [get]
func (e Profile) Query2FA(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	// 从请求中获取用户邮箱
	email := c.Query("email")
	if email == "" {
		e.Error(400, nil, fault.GetErrorMessage("email_parameter_is_required", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	u, _ := s.Get2FAStatus(email, lang)

	e.OK(u, fault.GetErrorMessage("acquire_success", lang))
}

type update2FAReq struct {
	Type          int    `json:"type" binding:"required"` //1绑定2解绑
	Secret        string `json:"secret"`                  //绑定时必传
	TwoFactorCode string `json:"twoFactorCode" binding:"required"`
	EmailCode     string `json:"emailCode" binding:"required"`
}

type bindWalletReq struct {
	WalletAddress string `json:"walletAddress" binding:"required"` // 钱包地址
}

type updateMessageSettingReq struct {
	Type   int `json:"type" binding:"required"`   // 1-资讯, 2-新点位, 3-点位状态, 4-系统
	Status int `json:"status" binding:"required"` // 1-开启, 2-关闭
}

func (e Profile) Update2FA(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}
	var req update2FAReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}
	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	err2 := s.Update2FA(u, req.Type, req.Secret, req.TwoFactorCode, req.EmailCode, lang)
	if err2 != nil {
		e.Error(400, err2, err2.Error())
		return
	}
	e.OK(nil, fault.GetErrorMessage("update_successfully", lang))
}

// BindWallet 绑定钱包地址
// @Summary 绑定钱包地址
// @Description 绑定钱包地址
// @Tags 用户钱包
// @Accept application/json
// @Param data body bindWalletReq true "data"
// @Router /api/v1/app/profile/bindWallet [post]
// @Security Bearer
func (e Profile) BindWallet(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req bindWalletReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.BindWallet(int64(user.Id), req.WalletAddress, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("bind_success", lang))
}

// GetMessageSetting 获取消息通知设置列表
// @Summary 获取消息通知设置列表
// @Description 获取消息通知设置列表
// @Tags 用户消息
// @Router /api/v1/app/profile/messageSettingList [get]
// @Security Bearer
func (e Profile) GetMessageSetting(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetMessageSetting(user.Id, lang)
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// UpdateMessageSetting 更新消息通知设置
// @Summary 更新消息通知设置
// @Description 更新消息通知设置
// @Tags 用户消息
// @Accept application/json
// @Param data body updateMessageSettingReq true "data"
// @Router /api/v1/app/profile/messageSettingUpdate [post]
// @Security Bearer
func (e Profile) UpdateMessageSetting(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req updateMessageSettingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.UpdateMessageSetting(user.Id, req.Type, req.Status, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("update_successfully", lang))
}

type personalInfoReq struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Head          string `json:"head"`
	WalletAddress string `json:"walletAddress"`
	EmailCode     string `json:"emailCode"`
}

type refreshTokenReq struct {
	IsRemember bool `json:"isRemember"`
}

// RefreshToken 刷新token
// @Summary 刷新token
// @Description 刷新token
// @Tags 用户认证
// @Accept application/json
// @Param data body refreshTokenReq true "data"
// @Router /api/v1/app/profile/refreshToken [post]
// @Security Bearer
func (e Profile) RefreshToken(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req refreshTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	// 根据IsRemember参数设置不同的过期时间
	var expireTime int
	if req.IsRemember {
		expireTime = 30 * 24 * 60 // 30天
	} else {
		expireTime = 24 * 60 // 24小时
	}

	// 生成新的token
	err = s.GenerateToken(user.Id, expireTime)
	if err != nil {
		e.Error(500, err, fault.GetErrorMessage("token_refreshed_failed", lang))
		return
	}

	e.OK(gin.H{
		"expireTime": expireTime,
	}, fault.GetErrorMessage("token_refreshed_successfully", lang))
}

// GetPersonalInfo 获取个人信息
// @Summary 获取个人信息
// @Description 获取个人信息
// @Tags 个人中心
// @Router /api/v1/app/profile/personalInfo [get]
// @Security Bearer
func (e Profile) GetPersonalInfo(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetPersonalInfo(user.Id)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("acquire_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// UpdatePersonalInfo 更新个人信息
// @Summary 更新个人信息
// @Description 更新个人信息
// @Tags 个人中心
// @Accept application/json
// @Param data body personalInfoReq true "data"
// @Router /api/v1/app/profile/updatePersonalInfo [post]
// @Security Bearer
func (e Profile) UpdatePersonalInfo(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req personalInfoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	s := webService.AppProfile{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.UpdatePersonalInfo(user.Id, req.Username, req.Email, req.Head, req.WalletAddress, req.EmailCode, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("update_successfully", lang))
}
