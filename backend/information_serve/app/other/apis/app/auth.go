package web

import (
	"errors"
	webService "go-admin/app/other/service"
	fault "go-admin/common/errors"
	"go-admin/common/middleware/handler"
	"go-admin/common/util"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
)

type Auth struct {
	AppApi
}

type registerReq struct {
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	Code          string `json:"code" binding:"required"`
	DeviceId      string `json:"deviceId" binding:"required"`      //设备id
	DeviceType    int64  `json:"deviceType" binding:"required"`    //设备类型：1-iOS，2-Android，3-Web
	SystemVersion string `json:"systemVersion" binding:"required"` //系统版本
	AppVersion    string `json:"appVersion" binding:"required"`    //app版本
}

var (
	emailRe = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)
)

// Register APP用户端注册（邮箱 + 两次密码一致 + 密码强度校验）
// @Summary APP用户端-注册
// @Tags APP用户端
// @Accept application/json
// @Param data body registerReq true "data"
// @Router /api/v1/app/register [post]
func (e Auth) Register(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}

	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}
	//req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if !emailRe.MatchString(req.Email) {
		e.Error(400, errors.New(""), fault.GetErrorMessage("the_email_format_is_incorrect", lang))
		return
	}

	if err := util.ValidatePassword(req.Password, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	s := webService.AppAuth{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	userId, err := s.Register(req.Email, req.Password, req.Code, req.DeviceId, req.SystemVersion, req.AppVersion, lang, req.DeviceType)
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(gin.H{"userId": userId}, fault.GetErrorMessage("registered_successfully", lang))
}

// Logout APP用户端退出登录（需携带有效 token）
// @Summary APP用户端-退出登录
// @Tags APP用户端
// @Router /api/v1/app/logout [post]
// @Security Bearer
func (e Auth) Logout(c *gin.Context) {
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	e.MakeContext(c)
	handler.LoginLogToDB(c, "2", fault.GetErrorMessage("exit_successful", lang), user.GetUserName(c))
	e.OK(nil, fault.GetErrorMessage("exit_successful", lang))
}
