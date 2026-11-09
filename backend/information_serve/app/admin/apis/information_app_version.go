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

type InformationAppVersion struct {
	api.Api
}

// GetPage 获取版本管理列表
// @Summary 获取版本管理列表
// @Description 获取版本管理列表
// @Tags 版本管理
// @Param platform query int64 false "平台：1-iOS，2-Android"
// @Param forceUpdate query int64 false "是否强制更新：2-否，1-是"
// @Param updateType query int64 false "类型 1-功能更新 2-问题修复 3-性能优化"
// @Param status query int64 false "状态 1-发布 2-草稿 3-撤回"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationAppVersion}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-app-version [get]
// @Security Bearer
func (e InformationAppVersion) GetPage(c *gin.Context) {
    req := dto.InformationAppVersionGetPageReq{}
    s := service.InformationAppVersion{}
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
	list := make([]models.InformationAppVersion, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取版本管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取版本管理
// @Summary 获取版本管理
// @Description 获取版本管理
// @Tags 版本管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationAppVersion} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-app-version/{id} [get]
// @Security Bearer
func (e InformationAppVersion) Get(c *gin.Context) {
	req := dto.InformationAppVersionGetReq{}
	s := service.InformationAppVersion{}
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
	var object models.InformationAppVersion

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取版本管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建版本管理
// @Summary 创建版本管理
// @Description 创建版本管理
// @Tags 版本管理
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationAppVersionInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-app-version [post]
// @Security Bearer
func (e InformationAppVersion) Insert(c *gin.Context) {
    req := dto.InformationAppVersionInsertReq{}
    s := service.InformationAppVersion{}
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
		e.Error(500, err, fmt.Sprintf("创建版本管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改版本管理
// @Summary 修改版本管理
// @Description 修改版本管理
// @Tags 版本管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationAppVersionUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-app-version/{id} [put]
// @Security Bearer
func (e InformationAppVersion) Update(c *gin.Context) {
    req := dto.InformationAppVersionUpdateReq{}
    s := service.InformationAppVersion{}
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
		e.Error(500, err, fmt.Sprintf("修改版本管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除版本管理
// @Summary 删除版本管理
// @Description 删除版本管理
// @Tags 版本管理
// @Param data body dto.InformationAppVersionDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-app-version [delete]
// @Security Bearer
func (e InformationAppVersion) Delete(c *gin.Context) {
    s := service.InformationAppVersion{}
    req := dto.InformationAppVersionDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除版本管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
