package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type InformationUser struct {
	api.Api
}

// GetPage 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags 用户
// @Param username query string false "用户名"
// @Param deviceType query int64 false "设备类型：1-iOS，2-Android，3-Web"
// @Param walletAddress query string false "绑定的钱包地址"
// @Param pointNotificationSettings query int64 false "点位通知设置：2-全部，1-部分"
// @Param status query int64 false "状态：3-禁用，1-正常，2-锁定"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationUser}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-user [get]
// @Security Bearer
func (e InformationUser) GetPage(c *gin.Context) {
    req := dto.InformationUserGetPageReq{}
    s := service.InformationUser{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.InformationUser, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户
// @Summary 获取用户
// @Description 获取用户
// @Tags 用户
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-user/{id} [get]
// @Security Bearer
func (e InformationUser) Get(c *gin.Context) {
	req := dto.InformationUserGetReq{}
	s := service.InformationUser{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.InformationUser

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags 用户
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationUserInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-user [post]
// @Security Bearer
func (e InformationUser) Insert(c *gin.Context) {
    req := dto.InformationUserInsertReq{}
    s := service.InformationUser{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建用户失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户
// @Summary 修改用户
// @Description 修改用户
// @Tags 用户
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationUserUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-user/{id} [put]
// @Security Bearer
func (e InformationUser) Update(c *gin.Context) {
    req := dto.InformationUserUpdateReq{}
    s := service.InformationUser{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改用户失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户
// @Param data body dto.InformationUserDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-user [delete]
// @Security Bearer
func (e InformationUser) Delete(c *gin.Context) {
    s := service.InformationUser{}
    req := dto.InformationUserDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除用户失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
