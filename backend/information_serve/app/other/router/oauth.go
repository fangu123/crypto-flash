package router

import (
	webService "go-admin/app/other/apis/app"

	"github.com/gin-gonic/gin"
)

// InitOAuthRouter 初始化OAuth路由
func InitOAuthRouter(Router *gin.RouterGroup) {
	// OAuth相关路由
	oauthRouter := Router.Group("oauth")
	{
		// Google OAuth路由
		googleRouter := oauthRouter.Group("google")
		{
			googleRouter.POST("/auth", webService.OAuth{}.GoogleAuthURL)      // 获取Google认证URL
			googleRouter.POST("/callback", webService.OAuth{}.GoogleCallback) // Google回调处理
		}

		// Microsoft OAuth路由
		microsoftRouter := oauthRouter.Group("microsoft")
		{
			microsoftRouter.POST("/auth", webService.OAuth{}.MicrosoftAuthURL)      // 获取Microsoft认证URL
			microsoftRouter.POST("/callback", webService.OAuth{}.MicrosoftCallback) // Microsoft回调处理
		}
	}
}
