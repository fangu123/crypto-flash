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

type InformationNotificationUser struct {
	api.Api
}

// GetPage 获取消息通知用户关联列表
// @Summary 获取消息通知用户关联列表
// @Description 获取消息通知用户关联列表
// @Tags 消息通知用户关联
// @Param userId query int64 false "用户ID"
// @Param notificationId query int64 false "通知ID，关联 information_notification 表"
// @Param isRead query int64 false "是否已读：2-未读，1-已读"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InformationNotificationUser}} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-notification-user [get]
// @Security Bearer
func (e InformationNotificationUser) GetPage(c *gin.Context) {
    req := dto.InformationNotificationUserGetPageReq{}
    s := service.InformationNotificationUser{}
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
	list := make([]models.InformationNotificationUser, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取消息通知用户关联失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取消息通知用户关联
// @Summary 获取消息通知用户关联
// @Description 获取消息通知用户关联
// @Tags 消息通知用户关联
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InformationNotificationUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/information-notification-user/{id} [get]
// @Security Bearer
func (e InformationNotificationUser) Get(c *gin.Context) {
	req := dto.InformationNotificationUserGetReq{}
	s := service.InformationNotificationUser{}
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
	var object models.InformationNotificationUser

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取消息通知用户关联失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建消息通知用户关联
// @Summary 创建消息通知用户关联
// @Description 创建消息通知用户关联
// @Tags 消息通知用户关联
// @Accept application/json
// @Product application/json
// @Param data body dto.InformationNotificationUserInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/information-notification-user [post]
// @Security Bearer
func (e InformationNotificationUser) Insert(c *gin.Context) {
    req := dto.InformationNotificationUserInsertReq{}
    s := service.InformationNotificationUser{}
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
		e.Error(500, err, fmt.Sprintf("创建消息通知用户关联失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改消息通知用户关联
// @Summary 修改消息通知用户关联
// @Description 修改消息通知用户关联
// @Tags 消息通知用户关联
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InformationNotificationUserUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/information-notification-user/{id} [put]
// @Security Bearer
func (e InformationNotificationUser) Update(c *gin.Context) {
    req := dto.InformationNotificationUserUpdateReq{}
    s := service.InformationNotificationUser{}
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
		e.Error(500, err, fmt.Sprintf("修改消息通知用户关联失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除消息通知用户关联
// @Summary 删除消息通知用户关联
// @Description 删除消息通知用户关联
// @Tags 消息通知用户关联
// @Param data body dto.InformationNotificationUserDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/information-notification-user [delete]
// @Security Bearer
func (e InformationNotificationUser) Delete(c *gin.Context) {
    s := service.InformationNotificationUser{}
    req := dto.InformationNotificationUserDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除消息通知用户关联失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
