package web

import (
	fault "go-admin/common/errors"
	"strconv"

	"github.com/gin-gonic/gin"

	webService "go-admin/app/other/service"
)

type News struct {
	AppApi
}

// GetNewsList 获取新闻列表
// @Summary 获取新闻列表
// @Description 获取新闻列表，支持分类查询、标签筛选和关键词搜索
// @Tags 新闻
// @Param tagId query int false "标签ID" default(0)
// @Param keyword query string false "关键词" default("")
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Router /api/v1/app/news/list [get]
// @Security Bearer
func (e News) GetNewsList(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	//_, err = e.currentUser(c)
	//if err != nil {
	//	e.Error(401, err, fault.GetErrorMessage("not_login", lang))
	//	return
	//}

	tagId, _ := strconv.ParseInt(c.Query("tagId"), 10, 64)
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	s := webService.AppNews{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, total, err := s.GetNewsList(tagId, keyword, page, size)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_news_list_fail", lang))
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

// GetNewsDetail 获取新闻详情
// @Summary 获取新闻详情
// @Description 获取新闻详情，并增加浏览次数
// @Tags 新闻
// @Param id path int true "新闻ID"
// @Router /api/v1/app/news/detail/{id} [get]
// @Security Bearer
func (e News) GetNewsDetail(c *gin.Context) {
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

	newsId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("news_id_invalid", lang))
		return
	}

	s := webService.AppNews{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetNewsDetail(newsId, int64(user.Id), user.PointNotificationSettings, lang)
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// ShareNews 分享新闻
// @Summary 分享新闻
// @Description 分享新闻，增加分享次数
// @Tags 新闻
// @Param id path int true "新闻ID"
// @Router /api/v1/app/news/share/{id} [post]
// @Security Bearer
func (e News) ShareNews(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	_, err = e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}

	newsId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("news_id_invalid", lang))
		return
	}

	s := webService.AppNews{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.ShareNews(newsId, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("share_news_success", lang))
}

// FavoriteNews 收藏新闻
// @Summary 收藏新闻
// @Description 收藏新闻
// @Tags 新闻
// @Param id path int true "新闻ID"
// @Router /api/v1/app/news/favorite/{id} [post]
// @Security Bearer
func (e News) FavoriteNews(c *gin.Context) {
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

	newsId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("news_id_invalid", lang))
		return
	}

	s := webService.AppNews{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.FavoriteNews(int64(user.Id), newsId, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("collect_success", lang))
}

// UnfavoriteNews 取消收藏新闻
// @Summary 取消收藏新闻
// @Description 取消收藏新闻
// @Tags 新闻
// @Param id path int true "新闻ID"
// @Router /api/v1/app/news/unfavorite/{id} [post]
// @Security Bearer
func (e News) UnfavoriteNews(c *gin.Context) {
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

	newsId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("news_id_invalid", lang))
		return
	}

	s := webService.AppNews{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	if err := s.UnfavoriteNews(int64(user.Id), newsId, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(nil, fault.GetErrorMessage("cancel_collect_success", lang))
}

func (e News) GetTags(c *gin.Context) {
	err := e.MakeContext(c).MakeOrm().Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language") // 从请求头获取语言设置
	if lang == "" {
		lang = "cn" // 默认中文
	}
	//_, err = e.currentUser(c)
	//if err != nil {
	//	e.Error(401, err, fault.GetErrorMessage("not_login", lang))
	//	return
	//}

	s := webService.AppNews{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetTags()
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("acquire_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}
