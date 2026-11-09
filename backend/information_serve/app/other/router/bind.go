package router

import (
	webService "go-admin/app/other/apis/app"

	"github.com/gin-gonic/gin"
)

// InitBindRouter 初始化绑定账号路由
func InitBindRouter(Router *gin.RouterGroup) {
	// 绑定账号相关路由
	bindRouter := Router.Group("bind")
	{
		// Google绑定路由
		googleRouter := bindRouter.Group("google")
		{
			googleRouter.POST("/binding", webService.Bind{}.BindGoogle)       // 绑定Google账号
			googleRouter.DELETE("/unbinding", webService.Bind{}.UnbindGoogle) // 解绑Google账号
		}

		// Microsoft绑定路由
		microsoftRouter := bindRouter.Group("microsoft")
		{
			microsoftRouter.POST("/binding", webService.Bind{}.BindMicrosoft)       // 绑定Microsoft账号
			microsoftRouter.DELETE("/unbinding", webService.Bind{}.UnbindMicrosoft) // 解绑Microsoft账号
		}

		// 获取绑定状态
		bindRouter.GET("/status", webService.Bind{}.GetBindStatus) // 获取账号绑定状态
	}
}
