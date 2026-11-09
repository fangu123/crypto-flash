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

type InformationContentTag struct {
	api.Api
}

// GetPage 获取内容分类标签列表
// @Summary 获取内容分类标签列表
// @Description 获取内容分类标签列表
// @Tags 内容分类标签
// @Param name query string false "标签名称"
// @Param status query int64 false "状态：2-禁用，1-启用"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationContentTag}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-content-tag [get]
// @Security Bearer
func (e InformationContentTag) GetPage(c *gin.Context) {
    req := dto.InformationContentTagGetPageReq{}
    s := service.InformationContentTag{}
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
	list := make([]models.InformationContentTag, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取内容分类标签失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取内容分类标签
// @Summary 获取内容分类标签
// @Description 获取内容分类标签
// @Tags 内容分类标签
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationContentTag} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-content-tag/{id} [get]
// @Security Bearer
func (e InformationContentTag) Get(c *gin.Context) {
	req := dto.InformationContentTagGetReq{}
	s := service.InformationContentTag{}
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
	var object models.InformationContentTag

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取内容分类标签失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建内容分类标签
// @Summary 创建内容分类标签
// @Description 创建内容分类标签
// @Tags 内容分类标签
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationContentTagInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-content-tag [post]
// @Security Bearer
func (e InformationContentTag) Insert(c *gin.Context) {
    req := dto.InformationContentTagInsertReq{}
    s := service.InformationContentTag{}
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
		e.Error(500, err, fmt.Sprintf("创建内容分类标签失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改内容分类标签
// @Summary 修改内容分类标签
// @Description 修改内容分类标签
// @Tags 内容分类标签
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationContentTagUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-content-tag/{id} [put]
// @Security Bearer
func (e InformationContentTag) Update(c *gin.Context) {
    req := dto.InformationContentTagUpdateReq{}
    s := service.InformationContentTag{}
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
		e.Error(500, err, fmt.Sprintf("修改内容分类标签失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除内容分类标签
// @Summary 删除内容分类标签
// @Description 删除内容分类标签
// @Tags 内容分类标签
// @Param data body dto.InformationContentTagDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-content-tag [delete]
// @Security Bearer
func (e InformationContentTag) Delete(c *gin.Context) {
    s := service.InformationContentTag{}
    req := dto.InformationContentTagDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除内容分类标签失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
