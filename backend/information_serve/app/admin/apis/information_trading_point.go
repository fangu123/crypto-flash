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

type InformationTradingPoint struct {
	api.Api
}

// GetPage 获取点位列表
// @Summary 获取点位列表
// @Description 获取点位列表
// @Tags 点位
// @Param coinId query int64 false "币种ID"
// @Param type query int64 false "类型：1-做多，2-做空"
// @Param status query int64 false "状态：1-进行中，2-已失效，3-已撤单"
// @Param isNotified query int64 false "是否已通知：1-未通知，2-已通知"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationTradingPoint}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-trading-point [get]
// @Security Bearer
func (e InformationTradingPoint) GetPage(c *gin.Context) {
    req := dto.InformationTradingPointGetPageReq{}
    s := service.InformationTradingPoint{}
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
	list := make([]models.InformationTradingPoint, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取点位失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取点位
// @Summary 获取点位
// @Description 获取点位
// @Tags 点位
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationTradingPoint} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-trading-point/{id} [get]
// @Security Bearer
func (e InformationTradingPoint) Get(c *gin.Context) {
	req := dto.InformationTradingPointGetReq{}
	s := service.InformationTradingPoint{}
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
	var object models.InformationTradingPoint

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取点位失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建点位
// @Summary 创建点位
// @Description 创建点位
// @Tags 点位
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationTradingPointInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-trading-point [post]
// @Security Bearer
func (e InformationTradingPoint) Insert(c *gin.Context) {
    req := dto.InformationTradingPointInsertReq{}
    s := service.InformationTradingPoint{}
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
		e.Error(500, err, fmt.Sprintf("创建点位失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改点位
// @Summary 修改点位
// @Description 修改点位
// @Tags 点位
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationTradingPointUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-trading-point/{id} [put]
// @Security Bearer
func (e InformationTradingPoint) Update(c *gin.Context) {
    req := dto.InformationTradingPointUpdateReq{}
    s := service.InformationTradingPoint{}
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
		e.Error(500, err, fmt.Sprintf("修改点位失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除点位
// @Summary 删除点位
// @Description 删除点位
// @Tags 点位
// @Param data body dto.InformationTradingPointDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-trading-point [delete]
// @Security Bearer
func (e InformationTradingPoint) Delete(c *gin.Context) {
    s := service.InformationTradingPoint{}
    req := dto.InformationTradingPointDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除点位失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
