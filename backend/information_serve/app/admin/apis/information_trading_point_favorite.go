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

type InformationTradingPointFavorite struct {
	api.Api
}

// GetPage 获取点位收藏列表
// @Summary 获取点位收藏列表
// @Description 获取点位收藏列表
// @Tags 点位收藏
// @Param userId query int64 false "用户ID"
// @Param tradingPointId query int64 false "点位ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationTradingPointFavorite}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-trading-point-favorite [get]
// @Security Bearer
func (e InformationTradingPointFavorite) GetPage(c *gin.Context) {
    req := dto.InformationTradingPointFavoriteGetPageReq{}
    s := service.InformationTradingPointFavorite{}
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
	list := make([]models.InformationTradingPointFavorite, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取点位收藏失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取点位收藏
// @Summary 获取点位收藏
// @Description 获取点位收藏
// @Tags 点位收藏
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationTradingPointFavorite} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-trading-point-favorite/{id} [get]
// @Security Bearer
func (e InformationTradingPointFavorite) Get(c *gin.Context) {
	req := dto.InformationTradingPointFavoriteGetReq{}
	s := service.InformationTradingPointFavorite{}
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
	var object models.InformationTradingPointFavorite

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取点位收藏失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建点位收藏
// @Summary 创建点位收藏
// @Description 创建点位收藏
// @Tags 点位收藏
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationTradingPointFavoriteInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-trading-point-favorite [post]
// @Security Bearer
func (e InformationTradingPointFavorite) Insert(c *gin.Context) {
    req := dto.InformationTradingPointFavoriteInsertReq{}
    s := service.InformationTradingPointFavorite{}
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
		e.Error(500, err, fmt.Sprintf("创建点位收藏失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改点位收藏
// @Summary 修改点位收藏
// @Description 修改点位收藏
// @Tags 点位收藏
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationTradingPointFavoriteUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-trading-point-favorite/{id} [put]
// @Security Bearer
func (e InformationTradingPointFavorite) Update(c *gin.Context) {
    req := dto.InformationTradingPointFavoriteUpdateReq{}
    s := service.InformationTradingPointFavorite{}
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
		e.Error(500, err, fmt.Sprintf("修改点位收藏失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除点位收藏
// @Summary 删除点位收藏
// @Description 删除点位收藏
// @Tags 点位收藏
// @Param data body dto.InformationTradingPointFavoriteDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-trading-point-favorite [delete]
// @Security Bearer
func (e InformationTradingPointFavorite) Delete(c *gin.Context) {
    s := service.InformationTradingPointFavorite{}
    req := dto.InformationTradingPointFavoriteDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除点位收藏失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
