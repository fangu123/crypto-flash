package web

import (
	"errors"
	webService "go-admin/app/other/service"
	fault "go-admin/common/errors"
	"strconv"
	"strings"
	"time"

	con "go-admin/config"

	"github.com/gin-gonic/gin"
)

// OAuth 处理OAuth相关请求
type OAuth struct {
	AppApi
}

// GoogleAuthReq Google认证请求
type GoogleAuthReq struct {
	RedirectURL string `json:"redirectUrl" binding:"required"` // 重定向URL
}

// MicrosoftAuthReq Microsoft认证请求
type MicrosoftAuthReq struct {
	RedirectURL string `json:"redirectUrl" binding:"required"` // 重定向URL
}

// GoogleCallbackReq Google回调请求
type GoogleCallbackReq struct {
	Code  string `json:"code" binding:"required"`  // 授权码
	State string `json:"state" binding:"required"` // 状态参数
}

// MicrosoftCallbackReq Microsoft回调请求
type MicrosoftCallbackReq struct {
	Code  string `json:"code" binding:"required"`  // 授权码
	State string `json:"state" binding:"required"` // 状态参数
}

// GoogleAuthURL 获取Google认证URL
// @Summary 获取Google认证URL
// @Tags APP用户端-OAuth
// @Accept application/json
// @Param data body GoogleAuthReq true "data"
// @Router /api/v1/app/oauth/google/auth [post]
func (e OAuth) GoogleAuthURL(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var req GoogleAuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	// 验证重定向URL
	if !strings.Contains(req.RedirectURL, "http://") && !strings.Contains(req.RedirectURL, "https://") {
		e.Error(400, errors.New(""), fault.GetErrorMessage("invalid_redirect_url", lang))
		return
	}

	// 获取OAuth服务
	oauthService := webService.OAuthService{}

	// 生成状态参数
	state, err := oauthService.GenerateState()
	if err != nil {
		e.Error(500, err, fault.GetErrorMessage("get_google_auth_url_failed", lang))
		return
	}

	u, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 将状态存储到Redis，5分钟过期
	if err := oauthService.StoreStateToRedis(state, strconv.Itoa(u.Id), 5*time.Minute); err != nil {
		e.Error(500, err, fault.GetErrorMessage("get_google_auth_url_failed", lang))
		return
	}

	// 获取Google OAuth配置
	googleConfig := oauthService.GoogleOAuthConfig(
		con.ExtConfig.OAuth.Google.ClientID,     // 从配置文件读取Google Client ID
		con.ExtConfig.OAuth.Google.ClientSecret, // 从配置文件读取Google Client Secret
		req.RedirectURL,                         // 使用客户端提供的重定向URL
	)

	// 生成Google认证URL
	authURL := oauthService.GetGoogleAuthURL(googleConfig, state)

	e.OK(gin.H{"authUrl": authURL}, fault.GetErrorMessage("acquire_success", lang))
}

// MicrosoftAuthURL 获取Microsoft认证URL
// @Summary 获取Microsoft认证URL
// @Tags APP用户端-OAuth
// @Accept application/json
// @Param data body MicrosoftAuthReq true "data"
// @Router /api/v1/app/oauth/microsoft/auth [post]
func (e OAuth) MicrosoftAuthURL(c *gin.Context) {

	u, err := e.currentUser(c)
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var req MicrosoftAuthReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	// 验证重定向URL
	if !strings.Contains(req.RedirectURL, "http://") && !strings.Contains(req.RedirectURL, "https://") {
		e.Error(400, errors.New(""), fault.GetErrorMessage("invalid_redirect_url", lang))
		return
	}

	// 获取OAuth服务
	oauthService := webService.OAuthService{}

	// 生成状态参数
	state, err := oauthService.GenerateState()
	if err != nil {
		e.Error(500, err, fault.GetErrorMessage("get_microsoft_auth_url_failed", lang))
		return
	}

	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 将状态存储到Redis，5分钟过期
	if err := oauthService.StoreStateToRedis(state, strconv.Itoa(u.Id), 5*time.Minute); err != nil {
		e.Error(500, err, fault.GetErrorMessage("get_microsoft_auth_url_failed", lang))
		return
	}

	// 获取Microsoft OAuth配置
	microsoftConfig := oauthService.MicrosoftOAuthConfig(
		con.ExtConfig.OAuth.Microsoft.ClientID,     // 从配置文件读取Microsoft Client ID
		con.ExtConfig.OAuth.Microsoft.ClientSecret, // 从配置文件读取Microsoft Client Secret
		req.RedirectURL, // 使用客户端提供的重定向URL
	)

	// 生成Microsoft认证URL
	authURL := oauthService.GetMicrosoftAuthURL(microsoftConfig, state)

	e.OK(gin.H{"authUrl": authURL}, fault.GetErrorMessage("acquire_success", lang))
}

// GoogleCallback Google回调处理
// @Summary Google回调处理
// @Tags APP用户端-OAuth
// @Accept application/json
// @Param data body GoogleCallbackReq true "data"
// @Router /api/v1/app/oauth/google/callback [post]
func (e OAuth) GoogleCallback(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var req GoogleCallbackReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	// 获取OAuth服务
	oauthService := webService.OAuthService{}

	// 验证状态参数
	_, err := oauthService.VerifyStateFromRedis(req.State, lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("invalid_state", lang))
		return
	}

	// 获取Google OAuth配置（实际项目中应该从配置中读取）
	googleConfig := oauthService.GoogleOAuthConfig(
		con.ExtConfig.OAuth.Google.ClientID,     // 从配置文件读取Google Client ID
		con.ExtConfig.OAuth.Google.ClientSecret, // 从配置文件读取Google Client Secret
		con.ExtConfig.OAuth.Google.RedirectURL,  // 回调URL，实际项目中应该配置为固定的回调地址
	)

	// 使用授权码交换访问令牌
	token, err := oauthService.ExchangeGoogleCodeForToken(googleConfig, req.Code, lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("token_exchange_failed", lang))
		return
	}

	// 获取Google用户信息
	userInfo, err := oauthService.GetGoogleUserInfo(googleConfig, token, lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_user_info_failed", lang))
		return
	}

	// 返回用户信息和访问令牌
	e.OK(gin.H{
		"googleId":          userInfo.ID,
		"googleEmail":       userInfo.Email,
		"googleName":        userInfo.Name,
		"googlePicture":     userInfo.Picture,
		"googleAccessToken": token.AccessToken,
		"googleExpiresAt":   time.Now().Add(time.Duration(token.ExpiresIn) * time.Second).Unix(),
	}, fault.GetErrorMessage("acquire_success", lang))
}

// MicrosoftCallback Microsoft回调处理
// @Summary Microsoft回调处理
// @Tags APP用户端-OAuth
// @Accept application/json
// @Param data body MicrosoftCallbackReq true "data"
// @Router /api/v1/app/oauth/microsoft/callback [post]
func (e OAuth) MicrosoftCallback(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var req MicrosoftCallbackReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}

	// 获取OAuth服务
	oauthService := webService.OAuthService{}

	// 验证状态参数
	_, err := oauthService.VerifyStateFromRedis(req.State, lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("invalid_state", lang))
		return
	}

	// 获取Microsoft OAuth配置（实际项目中应该从配置中读取）
	microsoftConfig := oauthService.MicrosoftOAuthConfig(
		con.ExtConfig.OAuth.Microsoft.ClientID,     // 替换为实际的Microsoft Client ID
		con.ExtConfig.OAuth.Microsoft.ClientSecret, // 替换为实际的Microsoft Client Secret
		con.ExtConfig.OAuth.Microsoft.RedirectURL,  // 回调URL，实际项目中应该配置为固定的回调地址
	)

	// 使用授权码交换访问令牌
	token, err := oauthService.ExchangeMicrosoftCodeForToken(microsoftConfig, req.Code, lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("token_exchange_failed", lang))
		return
	}

	// 获取Microsoft用户信息
	userInfo, err := oauthService.GetMicrosoftUserInfo(microsoftConfig, token, lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_user_info_failed", lang))
		return
	}

	// 处理用户照片URL
	var pictureURL string
	if userInfo.Photo != nil {
		pictureURL = userInfo.Photo.Value
	}

	// 返回用户信息和访问令牌
	e.OK(gin.H{
		"microsoftId":          userInfo.ID,
		"microsoftEmail":       userInfo.Email,
		"microsoftName":        userInfo.DisplayName,
		"microsoftPicture":     pictureURL,
		"microsoftAccessToken": token.AccessToken,
		"microsoftExpiresAt":   time.Now().Add(time.Duration(token.ExpiresIn) * time.Second).Unix(),
	}, fault.GetErrorMessage("acquire_success", lang))
}
