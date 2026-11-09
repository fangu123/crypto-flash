package web

import (
	fault "go-admin/common/errors"

	webService "go-admin/app/other/service"

	"github.com/gin-gonic/gin"
)

// Bind 处理账号绑定相关请求
type Bind struct {
	AppApi
}

// BindGoogleReq 绑定Google账号请求
type BindGoogleReq struct {
	GoogleId          string `json:"googleId" binding:"required"`          // Google用户ID
	GoogleEmail       string `json:"googleEmail" binding:"required"`       // Google邮箱
	GoogleName        string `json:"googleName" binding:"required"`        // Google用户名
	GooglePicture     string `json:"googlePicture" binding:"required"`     // Google头像URL
	GoogleAccessToken string `json:"googleAccessToken" binding:"required"` // Google访问令牌
}

// BindMicrosoftReq 绑定Microsoft账号请求
type BindMicrosoftReq struct {
	MicrosoftId          string `json:"microsoftId" binding:"required"`          // Microsoft用户ID
	MicrosoftEmail       string `json:"microsoftEmail" binding:"required"`       // Microsoft邮箱
	MicrosoftName        string `json:"microsoftName" binding:"required"`        // Microsoft用户名
	MicrosoftPicture     string `json:"microsoftPicture" binding:"required"`     // Microsoft头像URL
	MicrosoftAccessToken string `json:"microsoftAccessToken" binding:"required"` // Microsoft访问令牌
}

// BindGoogle 绑定Google账号
// @Summary 绑定Google账号
// @Tags APP用户端-账号绑定
// @Accept application/json
// @Param data body BindGoogleReq true "data"
// @Router /api/v1/app/bind/google [post]
func (e Bind) BindGoogle(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	// 检查用户是否已登录
	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req BindGoogleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	// 使用BindService处理数据库操作
	bindService := webService.BindService{}
	err = bindService.BindGoogle(u.Id, req.GoogleId, req.GoogleEmail, req.GoogleName, req.GooglePicture, req.GoogleAccessToken, lang)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(gin.H{"message": "Google账号绑定成功"}, fault.GetErrorMessage("bind_success", lang))
}

// BindMicrosoft 绑定Microsoft账号
// @Summary 绑定Microsoft账号
// @Tags APP用户端-账号绑定
// @Accept application/json
// @Param data body BindMicrosoftReq true "data"
// @Router /api/v1/app/bind/microsoft [post]
func (e Bind) BindMicrosoft(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	var req BindMicrosoftReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	// 使用BindService处理数据库操作
	bindService := webService.BindService{}
	err = bindService.BindMicrosoft(u.Id, req.MicrosoftId, req.MicrosoftEmail, req.MicrosoftName, req.MicrosoftPicture, req.MicrosoftAccessToken, lang)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(gin.H{"message": "ok"}, fault.GetErrorMessage("bind_success", lang))
}

// UnbindGoogle 解绑Google账号
// @Summary 解绑Google账号
// @Tags APP用户端-账号绑定
// @Router /api/v1/app/bind/google [delete]
func (e Bind) UnbindGoogle(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	// 检查用户是否已登录
	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 使用BindService处理数据库操作
	bindService := webService.BindService{}
	err = bindService.UnbindGoogle(u.Id, lang)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(gin.H{"message": "ok"}, fault.GetErrorMessage("unbind_success", lang))
}

// UnbindMicrosoft 解绑Microsoft账号
// @Summary 解绑Microsoft账号
// @Tags APP用户端-账号绑定
// @Router /api/v1/app/bind/microsoft [delete]
func (e Bind) UnbindMicrosoft(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 使用BindService处理数据库操作
	bindService := webService.BindService{}
	err = bindService.UnbindMicrosoft(u.Id, lang)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	e.OK(gin.H{"message": "ok"}, fault.GetErrorMessage("unbind_success", lang))
}

// GetBindStatus 获取账号绑定状态
// @Summary 获取账号绑定状态
// @Tags APP用户端-账号绑定
// @Router /api/v1/app/bind/status [get]
func (e Bind) GetBindStatus(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	// 检查用户是否已登录
	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 使用BindService处理数据库操作
	bindService := webService.BindService{}
	status, err := bindService.GetBindStatus(u.Id, lang)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}

	// 返回绑定状态
	e.OK(status, fault.GetErrorMessage("acquire_success", lang))
}
