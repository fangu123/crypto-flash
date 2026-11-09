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

type InformationNewsTagRelation struct {
	api.Api
}

// GetPage 获取新闻内容标签关联表列表
// @Summary 获取新闻内容标签关联表列表
// @Description 获取新闻内容标签关联表列表
// @Tags 新闻内容标签关联表
// @Param newsId query int64 false "新闻ID"
// @Param tagId query int64 false "标签ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationNewsTagRelation}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-news-tag-relation [get]
// @Security Bearer
func (e InformationNewsTagRelation) GetPage(c *gin.Context) {
    req := dto.InformationNewsTagRelationGetPageReq{}
    s := service.InformationNewsTagRelation{}
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
	list := make([]models.InformationNewsTagRelation, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取新闻内容标签关联表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取新闻内容标签关联表
// @Summary 获取新闻内容标签关联表
// @Description 获取新闻内容标签关联表
// @Tags 新闻内容标签关联表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationNewsTagRelation} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-news-tag-relation/{id} [get]
// @Security Bearer
func (e InformationNewsTagRelation) Get(c *gin.Context) {
	req := dto.InformationNewsTagRelationGetReq{}
	s := service.InformationNewsTagRelation{}
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
	var object models.InformationNewsTagRelation

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取新闻内容标签关联表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建新闻内容标签关联表
// @Summary 创建新闻内容标签关联表
// @Description 创建新闻内容标签关联表
// @Tags 新闻内容标签关联表
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationNewsTagRelationInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-news-tag-relation [post]
// @Security Bearer
func (e InformationNewsTagRelation) Insert(c *gin.Context) {
    req := dto.InformationNewsTagRelationInsertReq{}
    s := service.InformationNewsTagRelation{}
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
		e.Error(500, err, fmt.Sprintf("创建新闻内容标签关联表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改新闻内容标签关联表
// @Summary 修改新闻内容标签关联表
// @Description 修改新闻内容标签关联表
// @Tags 新闻内容标签关联表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationNewsTagRelationUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-news-tag-relation/{id} [put]
// @Security Bearer
func (e InformationNewsTagRelation) Update(c *gin.Context) {
    req := dto.InformationNewsTagRelationUpdateReq{}
    s := service.InformationNewsTagRelation{}
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
		e.Error(500, err, fmt.Sprintf("修改新闻内容标签关联表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除新闻内容标签关联表
// @Summary 删除新闻内容标签关联表
// @Description 删除新闻内容标签关联表
// @Tags 新闻内容标签关联表
// @Param data body dto.InformationNewsTagRelationDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-news-tag-relation [delete]
// @Security Bearer
func (e InformationNewsTagRelation) Delete(c *gin.Context) {
    s := service.InformationNewsTagRelation{}
    req := dto.InformationNewsTagRelationDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除新闻内容标签关联表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
