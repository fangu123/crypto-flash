package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	fault "go-admin/common/errors"
	"go-admin/common/redishelper"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/microsoft"
)

// OAuthService 处理OAuth相关服务
type OAuthService struct {
	service.Service
}

// GoogleOAuthConfig Google OAuth配置
func (o *OAuthService) GoogleOAuthConfig(clientID, clientSecret, redirectURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

// MicrosoftOAuthConfig Microsoft OAuth配置
func (o *OAuthService) MicrosoftOAuthConfig(clientID, clientSecret, redirectURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"User.Read"},
		Endpoint:     microsoft.AzureADEndpoint("common"), // 使用"common"作为租户ID
	}
}

// GenerateState 生成随机状态字符串
func (o *OAuthService) GenerateState() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// StoreStateToRedis 将状态存储到Redis
func (o *OAuthService) StoreStateToRedis(state, userID string, expires time.Duration) error {
	ctx := context.Background()
	redisClient := redishelper.DefaultRedis.GetClient()
	key := fmt.Sprintf("oauth_state:%s", state)
	return redisClient.Set(ctx, key, userID, expires).Err()
}

// VerifyStateFromRedis 从Redis验证状态
func (o *OAuthService) VerifyStateFromRedis(state, lang string) (string, error) {
	ctx := context.Background()
	redisClient := redishelper.DefaultRedis.GetClient()
	key := fmt.Sprintf("oauth_state:%s", state)
	userID, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", errors.New(fault.GetErrorMessage("invalid_state_parameter", lang))
	}

	// 验证后删除状态
	redisClient.Del(ctx, key)
	return userID, nil
}

// GetGoogleAuthURL 获取Google认证URL
func (o *OAuthService) GetGoogleAuthURL(config *oauth2.Config, state string) string {
	return config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// GetMicrosoftAuthURL 获取Microsoft认证URL
func (o *OAuthService) GetMicrosoftAuthURL(config *oauth2.Config, state string) string {
	return config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

// ExchangeGoogleCodeForToken 使用授权码交换Google访问令牌
func (o *OAuthService) ExchangeGoogleCodeForToken(config *oauth2.Config, code, lang string) (*oauth2.Token, error) {
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, errors.New(fault.GetErrorMessage("get_google_access_token_failed", lang))
	}
	return token, nil
}

// ExchangeMicrosoftCodeForToken 使用授权码交换Microsoft访问令牌
func (o *OAuthService) ExchangeMicrosoftCodeForToken(config *oauth2.Config, code, lang string) (*oauth2.Token, error) {
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, errors.New(fault.GetErrorMessage("get_microsoft_access_token_failed", lang))
	}
	return token, nil
}

// GetGoogleUserInfo 获取Google用户信息
func (o *OAuthService) GetGoogleUserInfo(config *oauth2.Config, token *oauth2.Token, lang string) (*GoogleUserInfo, error) {
	// 使用config和token创建客户端
	client := config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, errors.New(fault.GetErrorMessage("get_google_user_info_failed", lang))
	}
	defer resp.Body.Close()

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, errors.New(fault.GetErrorMessage("parse_google_user_info_failed", lang))
	}

	return &userInfo, nil
}

// GetMicrosoftUserInfo 获取Microsoft用户信息
func (o *OAuthService) GetMicrosoftUserInfo(config *oauth2.Config, token *oauth2.Token, lang string) (*MicrosoftUserInfo, error) {
	// 使用config和token创建客户端
	client := config.Client(context.Background(), token)
	resp, err := client.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		return nil, errors.New(fault.GetErrorMessage("get_microsoft_user_info_failed", lang))
	}
	defer resp.Body.Close()

	var userInfo MicrosoftUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, errors.New(fault.GetErrorMessage("parse_microsoft_user_info_failed", lang))
	}

	return &userInfo, nil
}

// GoogleUserInfo Google用户信息
type GoogleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

// MicrosoftUserInfo Microsoft用户信息
type MicrosoftUserInfo struct {
	ID          string          `json:"id"`
	Email       string          `json:"userPrincipalName"`
	DisplayName string          `json:"displayName"`
	Photo       *MicrosoftPhoto `json:"photo"`
}

// MicrosoftPhoto Microsoft用户照片
type MicrosoftPhoto struct {
	Value string `json:"value"`
}
