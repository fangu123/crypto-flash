package web

import (
	webService "go-admin/app/other/service"
	fault "go-admin/common/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Points struct {
	AppApi
}

// GetTradingPointList 获取点位列表
// @Summary 获取点位列表
// @Description 获取点位列表，支持分类查询和收藏筛选
// @Tags 点位
// @Param type query int false "类型：1-做多，2-做空，3-自选 ，0-全部" default(0)
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Router /api/v1/app/trading-point/list [get]
// @Security Bearer
func (e Points) GetTradingPointList(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}
	userId := int64(user.Id)
	defaultValue := "10"
	if user.PointNotificationSettings == 1 { //如果是查看部分点位 就只查看两条数据
		defaultValue = "3"
	}

	// 解析查询参数
	pointType, _ := strconv.Atoi(c.DefaultQuery("type", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", defaultValue))

	s := webService.AppPoint{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, total, err := s.GetTradingPointList(pointType, page, size, int(userId))
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_point_list_fail", lang))
		return
	}

	e.OK(gin.H{
		"list": result,
		"pagination": gin.H{
			"page":  page,
			"size":  size,
			"total": total,
		},
	}, fault.GetErrorMessage("acquire_success", lang))
}

// GetTradingPointDetail 获取点位详情
// @Summary 获取点位详情
// @Description 获取点位详情
// @Tags 点位
// @Param id path int true "点位ID"
// @Router /api/v1/app/trading-point/detail/{id} [get]
// @Security Bearer
func (e Points) GetTradingPointDetail(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 解析点位ID
	pointId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("point_id_invalid", lang))
		return
	}

	userId := int64(user.Id)

	s := webService.AppPoint{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetTradingPointDetail(pointId, int(userId), lang)
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// ShareTradingPoint 分享点位
// @Summary 分享点位
// @Description 分享点位
// @Tags 点位
// @Param id path int true "点位ID"
// @Router /api/v1/app/trading-point/share/{id} [post]
// @Security Bearer
func (e Points) ShareTradingPoint(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 解析点位ID
	pointId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("point_id_invalid", lang))
		return
	}

	s := webService.AppPoint{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.ShareTradingPoint(int64(user.Id), pointId, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("share_point_success", lang))
}

// FavoriteTradingPoint 收藏点位
// @Summary 收藏点位
// @Description 收藏点位
// @Tags 点位
// @Param id path int true "点位ID"
// @Router /api/v1/app/trading-point/favorite/{id} [post]
// @Security Bearer
func (e Points) FavoriteTradingPoint(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 解析点位ID
	pointId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("point_id_invalid", lang))
		return
	}

	userId := int64(user.Id)

	s := webService.AppPoint{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.FavoriteTradingPoint(userId, pointId, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("collect_success", lang))
}

// UnfavoriteTradingPoint 取消收藏点位
// @Summary 取消收藏点位
// @Description 取消收藏点位
// @Tags 点位
// @Param id path int true "点位ID"
// @Router /api/v1/app/trading-point/unfavorite/{id} [post]
// @Security Bearer
func (e Points) UnfavoriteTradingPoint(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	// 解析点位ID
	pointId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("point_id_invalid", lang))
		return
	}

	userId := int64(user.Id)

	s := webService.AppPoint{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.UnfavoriteTradingPoint(userId, pointId, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("cancel_collect_success", lang))
}

// UnlockAllTradingPoint 解锁全部点位
// @Summary 解锁全部点位
// @Description 解锁全部点位，查询用户是否有相关的NFT，如果有就修改用户表的pointNotificationSettings为2
// @Tags 点位
// @Router /api/v1/app/trading-point/unlock-all [post]
// @Security Bearer
func (e Points) UnlockAllTradingPoint(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	user, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}
	userId := int64(user.Id)

	s := webService.AppPoint{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.UnlockAllTradingPoint(userId, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("unlock_success", lang))
}
