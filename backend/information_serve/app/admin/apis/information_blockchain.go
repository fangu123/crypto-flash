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

type InformationBlockchain struct {
	api.Api
}

// GetPage 获取链列表
// @Summary 获取链列表
// @Description 获取链列表
// @Tags 链
// @Param status query int64 false "状态：2-禁用，1-启用"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationBlockchain}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-blockchain [get]
// @Security Bearer
func (e InformationBlockchain) GetPage(c *gin.Context) {
    req := dto.InformationBlockchainGetPageReq{}
    s := service.InformationBlockchain{}
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
	list := make([]models.InformationBlockchain, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取链失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取链
// @Summary 获取链
// @Description 获取链
// @Tags 链
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationBlockchain} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-blockchain/{id} [get]
// @Security Bearer
func (e InformationBlockchain) Get(c *gin.Context) {
	req := dto.InformationBlockchainGetReq{}
	s := service.InformationBlockchain{}
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
	var object models.InformationBlockchain

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取链失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建链
// @Summary 创建链
// @Description 创建链
// @Tags 链
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationBlockchainInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-blockchain [post]
// @Security Bearer
func (e InformationBlockchain) Insert(c *gin.Context) {
    req := dto.InformationBlockchainInsertReq{}
    s := service.InformationBlockchain{}
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
		e.Error(500, err, fmt.Sprintf("创建链失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改链
// @Summary 修改链
// @Description 修改链
// @Tags 链
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationBlockchainUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-blockchain/{id} [put]
// @Security Bearer
func (e InformationBlockchain) Update(c *gin.Context) {
    req := dto.InformationBlockchainUpdateReq{}
    s := service.InformationBlockchain{}
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
		e.Error(500, err, fmt.Sprintf("修改链失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除链
// @Summary 删除链
// @Description 删除链
// @Tags 链
// @Param data body dto.InformationBlockchainDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-blockchain [delete]
// @Security Bearer
func (e InformationBlockchain) Delete(c *gin.Context) {
    s := service.InformationBlockchain{}
    req := dto.InformationBlockchainDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除链失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
