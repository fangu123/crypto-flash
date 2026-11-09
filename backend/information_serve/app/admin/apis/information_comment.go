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

type InformationComment struct {
	api.Api
}

// GetPage 获取资讯评论表列表
// @Summary 获取资讯评论表列表
// @Description 获取资讯评论表列表
// @Tags 资讯评论表
// @Param newsId query int64 false "资讯ID"
// @Param userId query int64 false "评论用户ID"
// @Param parentId query int64 false "父评论ID，1表示顶级评论"
// @Param replyToUserId query int64 false "被回复用户ID"
// @Param isTop query int64 false "是否置顶 2否 1是"
// @Param status query int64 false "状态 4待审核 1正常 2违规 3已删除"
// @Param createTime query time.Time false "发布时间"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationComment}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-comment [get]
// @Security Bearer
func (e InformationComment) GetPage(c *gin.Context) {
    req := dto.InformationCommentGetPageReq{}
    s := service.InformationComment{}
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
	list := make([]models.InformationComment, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资讯评论表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取资讯评论表
// @Summary 获取资讯评论表
// @Description 获取资讯评论表
// @Tags 资讯评论表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationComment} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-comment/{id} [get]
// @Security Bearer
func (e InformationComment) Get(c *gin.Context) {
	req := dto.InformationCommentGetReq{}
	s := service.InformationComment{}
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
	var object models.InformationComment

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取资讯评论表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建资讯评论表
// @Summary 创建资讯评论表
// @Description 创建资讯评论表
// @Tags 资讯评论表
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationCommentInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-comment [post]
// @Security Bearer
func (e InformationComment) Insert(c *gin.Context) {
    req := dto.InformationCommentInsertReq{}
    s := service.InformationComment{}
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
		e.Error(500, err, fmt.Sprintf("创建资讯评论表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改资讯评论表
// @Summary 修改资讯评论表
// @Description 修改资讯评论表
// @Tags 资讯评论表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationCommentUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-comment/{id} [put]
// @Security Bearer
func (e InformationComment) Update(c *gin.Context) {
    req := dto.InformationCommentUpdateReq{}
    s := service.InformationComment{}
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
		e.Error(500, err, fmt.Sprintf("修改资讯评论表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除资讯评论表
// @Summary 删除资讯评论表
// @Description 删除资讯评论表
// @Tags 资讯评论表
// @Param data body dto.InformationCommentDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-comment [delete]
// @Security Bearer
func (e InformationComment) Delete(c *gin.Context) {
    s := service.InformationComment{}
    req := dto.InformationCommentDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除资讯评论表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
