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

type InformationNews struct {
	api.Api
}

// GetPage 获取资讯列表
// @Summary 获取资讯列表
// @Description 获取资讯列表
// @Tags 资讯
// @Param title query string false "标题"
// @Param relatedCoinId query int64 false "相关币种id"
// @Param relatedTradingPointId query int64 false "相关点位id"
// @Param level query int64 false "级别：1-普通，2-重要，3-紧急"
// @Param sentiment query int64 false "情绪：1-中性，2-利好，3-利空"
// @Param isFeatured query int64 false "是否头条：2-否，1-是"
// @Param status query int64 false "状态：1-草稿，2-已发布，3-已下架"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationNews}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-news [get]
// @Security Bearer
func (e InformationNews) GetPage(c *gin.Context) {
    req := dto.InformationNewsGetPageReq{}
    s := service.InformationNews{}
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
	list := make([]models.InformationNews, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资讯失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资讯
// @Summary 获取资讯
// @Description 获取资讯
// @Tags 资讯
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationNews} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-news/{id} [get]
// @Security Bearer
func (e InformationNews) Get(c *gin.Context) {
	req := dto.InformationNewsGetReq{}
	s := service.InformationNews{}
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
	var object models.InformationNews

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资讯失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建资讯
// @Summary 创建资讯
// @Description 创建资讯
// @Tags 资讯
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationNewsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-news [post]
// @Security Bearer
func (e InformationNews) Insert(c *gin.Context) {
    req := dto.InformationNewsInsertReq{}
    s := service.InformationNews{}
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
		e.Error(500, err, fmt.Sprintf("创建资讯失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资讯
// @Summary 修改资讯
// @Description 修改资讯
// @Tags 资讯
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationNewsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-news/{id} [put]
// @Security Bearer
func (e InformationNews) Update(c *gin.Context) {
    req := dto.InformationNewsUpdateReq{}
    s := service.InformationNews{}
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
		e.Error(500, err, fmt.Sprintf("修改资讯失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除资讯
// @Summary 删除资讯
// @Description 删除资讯
// @Tags 资讯
// @Param data body dto.InformationNewsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-news [delete]
// @Security Bearer
func (e InformationNews) Delete(c *gin.Context) {
    s := service.InformationNews{}
    req := dto.InformationNewsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资讯失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
