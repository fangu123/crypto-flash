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

type InformationThirdpartyApikey struct {
	api.Api
}

// GetPage 获取第三方apikey列表
// @Summary 获取第三方apikey列表
// @Description 获取第三方apikey列表
// @Tags 第三方apikey
// @Param type query int64 false "类型：1-aiCoin，2-bybit，3-odaily，4-blockBeats"
// @Param status query int64 false "状态 1-启用 2-停用"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationThirdpartyApikey}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-thirdparty-apikey [get]
// @Security Bearer
func (e InformationThirdpartyApikey) GetPage(c *gin.Context) {
    req := dto.InformationThirdpartyApikeyGetPageReq{}
    s := service.InformationThirdpartyApikey{}
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
	list := make([]models.InformationThirdpartyApikey, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取第三方apikey失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取第三方apikey
// @Summary 获取第三方apikey
// @Description 获取第三方apikey
// @Tags 第三方apikey
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationThirdpartyApikey} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-thirdparty-apikey/{id} [get]
// @Security Bearer
func (e InformationThirdpartyApikey) Get(c *gin.Context) {
	req := dto.InformationThirdpartyApikeyGetReq{}
	s := service.InformationThirdpartyApikey{}
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
	var object models.InformationThirdpartyApikey

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取第三方apikey失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建第三方apikey
// @Summary 创建第三方apikey
// @Description 创建第三方apikey
// @Tags 第三方apikey
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationThirdpartyApikeyInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-thirdparty-apikey [post]
// @Security Bearer
func (e InformationThirdpartyApikey) Insert(c *gin.Context) {
    req := dto.InformationThirdpartyApikeyInsertReq{}
    s := service.InformationThirdpartyApikey{}
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
		e.Error(500, err, fmt.Sprintf("创建第三方apikey失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改第三方apikey
// @Summary 修改第三方apikey
// @Description 修改第三方apikey
// @Tags 第三方apikey
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationThirdpartyApikeyUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-thirdparty-apikey/{id} [put]
// @Security Bearer
func (e InformationThirdpartyApikey) Update(c *gin.Context) {
    req := dto.InformationThirdpartyApikeyUpdateReq{}
    s := service.InformationThirdpartyApikey{}
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
		e.Error(500, err, fmt.Sprintf("修改第三方apikey失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除第三方apikey
// @Summary 删除第三方apikey
// @Description 删除第三方apikey
// @Tags 第三方apikey
// @Param data body dto.InformationThirdpartyApikeyDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-thirdparty-apikey [delete]
// @Security Bearer
func (e InformationThirdpartyApikey) Delete(c *gin.Context) {
    s := service.InformationThirdpartyApikey{}
    req := dto.InformationThirdpartyApikeyDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除第三方apikey失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
