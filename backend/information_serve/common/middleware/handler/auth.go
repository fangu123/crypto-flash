package handler

import (
	"fmt"
	"go-admin/app/admin/models"
	"go-admin/common"
	"net/http"
	"time"

	"go-admin/common/global"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/captcha"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"github.com/mssola/user_agent"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(SysUser)
		r, _ := v["role"].(SysRole)
		return jwt.MapClaims{
			jwt.IdentityKey:  u.UserId,
			jwt.RoleIdKey:    r.RoleId,
			jwt.RoleKey:      r.RoleKey,
			jwt.NiceKey:      u.Username,
			jwt.DataScopeKey: r.DataScope,
			jwt.RoleNameKey:  r.RoleName,
		}
	}
	return jwt.MapClaims{}
}

func WebPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(User)
		// 生成唯一的token ID
		tokenId := fmt.Sprintf("%d_%d", u.Id, time.Now().UnixNano())
		return jwt.MapClaims{
			jwt.IdentityKey: u.Id,
			jwt.NiceKey:     u.Username,
			"tokenId":       tokenId,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"RoleKey":     claims["rolekey"],
		"UserId":      claims["identity"],
		"RoleIds":     claims["roleid"],
		"DataScope":   claims["datascope"],
	}
}

func WebIdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"UserId":      claims["identity"],
	}
}

// Authenticator 获取token
// @Summary 登陆
// @Description 获取token
// @Description LoginHandler can be used by clients to get a jwt token.
// @Description Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// @Description Reply will be of the form {"token": "TOKEN"}.
// @Description dev mode：It should be noted that all fields cannot be empty, and a value of 0 can be passed in addition to the account password
// @Description 注意：开发模式：需要注意全部字段不能为空，账号密码外可以传入0值
// @Tags 登陆
// @Accept  application/json
// @Product application/json
// @Param account body Login  true "account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /api/v1/login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	log := api.GetRequestLogger(c)
	db, err := pkg.GetOrm(c)
	if err != nil {
		log.Errorf("get db error, %s", err.Error())
		response.Error(c, 500, err, "数据库连接获取失败")
		return nil, jwt.ErrFailedAuthentication
	}

	var loginVals Login
	var status = "2"
	var msg = "登录成功"
	var username = ""
	defer func() {
		LoginLogToDB(c, status, msg, username)
	}()

	if err = c.ShouldBind(&loginVals); err != nil {
		username = loginVals.Username
		msg = "数据解析失败"
		status = "1"

		return nil, jwt.ErrMissingLoginValues
	}
	//if config.ApplicationConfig.Mode != "dev" {
	if !captcha.Verify(loginVals.UUID, loginVals.Code, true) {
		username = loginVals.Username
		msg = "验证码错误"
		status = "1"

		return nil, jwt.ErrInvalidVerificationode
	}
	//}
	sysUser, role, e := loginVals.GetUser(db)
	if e == nil {
		username = loginVals.Username

		// 检查是否启用了2FA
		//if sysUser.TwoFactorEnabled {
		//	// 如果用户启用了2FA，必须提供验证码
		//	if loginVals.GoogleCode == "" {
		//		msg = "请输入双因素验证码"
		//		status = "1"
		//		return nil, errors.New("请输入谷歌验证码")
		//	}
		//
		//	// 验证2FA验证码
		//	valid := util.ValidateTOTP(sysUser.TwoFactorSecret, loginVals.GoogleCode, 6, 30, 1)
		//	if !valid {
		//		// 如果验证码无效，尝试备用码
		//		if sysUser.TwoFactorBackupCodes != "" {
		//			// 这里简化处理，实际应该验证备用码
		//			msg = "双因素验证码错误"
		//			status = "1"
		//			return nil, jwt.ErrFailedAuthentication
		//		} else {
		//			msg = "双因素验证码错误"
		//			status = "1"
		//			return nil, errors.New("谷歌验证码错误")
		//		}
		//	}
		//}

		return map[string]interface{}{"user": sysUser, "role": role}, nil
	} else {
		msg = "登录失败"
		status = "1"
		log.Warnf("%s login failed!", loginVals.Username)
	}
	return nil, jwt.ErrFailedAuthentication
}

// LoginLogToDB Write log to database
func LoginLogToDB(c *gin.Context, status string, msg string, username string) {
	if !config.LoggerConfig.EnabledDB {
		return
	}
	log := api.GetRequestLogger(c)
	l := make(map[string]interface{})

	ua := user_agent.New(c.Request.UserAgent())
	l["ipaddr"] = common.GetClientIP(c)
	l["loginLocation"] = "" // pkg.GetLocation(common.GetClientIP(c),gaConfig.ExtConfig.AMap.Key)
	l["loginTime"] = pkg.GetCurrentTime()
	l["status"] = status
	l["remark"] = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	l["browser"] = browserName + " " + browserVersion
	l["os"] = ua.OS()
	l["platform"] = ua.Platform()
	l["username"] = username
	l["msg"] = msg

	q := sdk.Runtime.GetMemoryQueue(c.Request.Host)
	message, err := sdk.Runtime.GetStreamMessage("", global.LoginLog, l)
	if err != nil {
		log.Errorf("GetStreamMessage error, %s", err.Error())
		//日志报错错误，不中断请求
	} else {
		err = q.Append(message)
		if err != nil {
			log.Errorf("Append message error, %s", err.Error())
		}
	}
}

// LogOut
// @Summary 退出登录
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "msg": "成功退出系统" }"
// @Router /logout [post]
// @Security Bearer
func LogOut(c *gin.Context) {
	LoginLogToDB(c, "2", "退出成功", user.GetUserName(c))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}

func Authorizator(data interface{}, c *gin.Context) bool {

	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(models.SysUser)
		r, _ := v["role"].(models.SysRole)
		c.Set("role", r.RoleName)
		c.Set("roleIds", r.RoleId)
		c.Set("userId", u.UserId)
		c.Set("userName", u.Username)
		c.Set("dataScope", r.DataScope)
		return true
	}
	return false
}

func WebAuthorization(data interface{}, c *gin.Context) bool {

	if _, ok := data.(map[string]interface{}); ok {
		//u, _ := v["user"].(models.User)
		//c.Set("userId", u.UserId)
		//c.Set("userName", u.Username)
		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

// middleware/auth.go
/*func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从token中获取用户信息
		claims := c.MustGet("claims").(*jwt.CustomClaims)

		// 查询用户语言设置
		var user models.InformationUser
		if err := global.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"code": 401, "message": "User not found"})
			c.Abort()
			return
		}

		// 设置用户语言到上下文
		utils.SetUserLanguage(c, user.Language)
		c.Next()
	}
}*/
