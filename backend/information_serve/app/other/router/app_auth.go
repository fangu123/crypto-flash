package router

import (
	"github.com/gin-gonic/gin"

	appApis "go-admin/app/other/apis/app"
	common "go-admin/common/middleware"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerWebAuthRouter)
}

func registerWebAuthRouter(v1 *gin.RouterGroup) {
	// 注册
	auth := appApis.Auth{}
	v1.POST("/app/register", auth.Register)

	// 登录（邮箱+密码，强制2FA验证）
	webAuthMiddleware, err := common.AuthInitWebEmailWith2FA()
	if err != nil {
		// 初始化失败时不注册登录路由，避免启动报错
		return
	}
	v1.POST("/app/login", webAuthMiddleware.LoginHandler)
	//验证码登录(邮箱+验证码，如果开启2FA就需要输入谷歌验证码)
	webAuthCaptchaMiddleware, err := common.AuthInitWebEmailCodeWith2FA()
	v1.POST("/app/login/captcha", webAuthCaptchaMiddleware.LoginHandler)

	//直接通过钱包地址 注册 登录
	webAuthWalletMiddleware, err := common.AuthInitWebWallet()
	if err != nil {
		// 初始化失败时不注册登录路由，避免启动报错
		return
	}
	v1.POST("/app/login/wallet", webAuthWalletMiddleware.LoginHandler)

	//google登录
	webAuthGoogleMiddleware, err := common.AuthInitWebGoogle()
	if err != nil {
		// 初始化失败时不注册登录路由，避免启动报错
		return
	}
	v1.POST("/app/login/google", webAuthGoogleMiddleware.LoginHandler)

	//microsoft登录
	webAuthMicrosoftMiddleware, err := common.AuthInitWebMicrosoft()
	if err != nil {
		// 初始化失败时不注册登录路由，避免启动报错
		return
	}
	v1.POST("/app/login/microsoft", webAuthMicrosoftMiddleware.LoginHandler)
}
