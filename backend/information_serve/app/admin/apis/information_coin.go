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

type InformationCoin struct {
	api.Api
}

// GetPage 获取币种列表
// @Summary 获取币种列表
// @Description 获取币种列表
// @Tags 币种
// @Param blockchainId query int64 false "所属链ID"
// @Param symbol query string false "币种标识"
// @Param status query int64 false "状态：2-禁用，1-启用"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationCoin}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-coin [get]
// @Security Bearer
func (e InformationCoin) GetPage(c *gin.Context) {
    req := dto.InformationCoinGetPageReq{}
    s := service.InformationCoin{}
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
	list := make([]models.InformationCoin, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取币种失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取币种
// @Summary 获取币种
// @Description 获取币种
// @Tags 币种
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationCoin} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-coin/{id} [get]
// @Security Bearer
func (e InformationCoin) Get(c *gin.Context) {
	req := dto.InformationCoinGetReq{}
	s := service.InformationCoin{}
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
	var object models.InformationCoin

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取币种失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建币种
// @Summary 创建币种
// @Description 创建币种
// @Tags 币种
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationCoinInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-coin [post]
// @Security Bearer
func (e InformationCoin) Insert(c *gin.Context) {
    req := dto.InformationCoinInsertReq{}
    s := service.InformationCoin{}
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
		e.Error(500, err, fmt.Sprintf("创建币种失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改币种
// @Summary 修改币种
// @Description 修改币种
// @Tags 币种
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationCoinUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-coin/{id} [put]
// @Security Bearer
func (e InformationCoin) Update(c *gin.Context) {
    req := dto.InformationCoinUpdateReq{}
    s := service.InformationCoin{}
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
		e.Error(500, err, fmt.Sprintf("修改币种失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除币种
// @Summary 删除币种
// @Description 删除币种
// @Tags 币种
// @Param data body dto.InformationCoinDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-coin [delete]
// @Security Bearer
func (e InformationCoin) Delete(c *gin.Context) {
    s := service.InformationCoin{}
    req := dto.InformationCoinDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除币种失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
