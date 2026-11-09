package web

import (
	"strconv"

	webService "go-admin/app/other/service"
	fault "go-admin/common/errors"

	"github.com/gin-gonic/gin"
)

type Comment struct {
	AppApi
}

// GetCommentList 获取资讯评论列表
// @Summary 获取资讯评论列表
// @Description 获取资讯评论列表，parentId 默认 1 表示顶级评论
// @Tags 评论
// @Param newsId query int true "资讯ID"
// @Param parentId query int false "父评论ID，1表示顶级评论"
// @Param page query int false "页码，默认1"
// @Param size query int false "每页数量，默认10，最大200"
// @Router /api/v1/app/comment/list [get]
// @Security Bearer
func (e Comment) GetCommentList(c *gin.Context) {
	if err := e.MakeContext(c).MakeOrm().Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "cn"
	}
	_, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}
	newsID, err := strconv.ParseInt(c.Query("newsId"), 10, 64)
	if err != nil || newsID <= 0 {
		e.Error(400, err, fault.GetErrorMessage("invalid_news_id", lang))
		return
	}
	parentID := int64(1)
	if value := c.Query("parentId"); value != "" {
		parentID, err = strconv.ParseInt(value, 10, 64)
		if err != nil || parentID <= 0 {
			e.Error(400, err, fault.GetErrorMessage("invalid_parent_comment_id", lang))
			return
		}
	}
	page, size := parsePage(c)
	s := webService.AppInformationComment{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	list, total, err := s.GetCommentList(newsID, parentID, int64(page), int64(size), lang)
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}
	e.OK(gin.H{
		"list":       list,
		"pagination": gin.H{"page": page, "size": size, "total": total},
	}, fault.GetErrorMessage("acquire_success", lang))
}

type publishCommentReq struct {
	NewsID   int64  `json:"newsId" binding:"required,gt=0"`
	ParentID int64  `json:"parentId"`
	Content  string `json:"content" binding:"required"`
}

// PublishComment 发布评论或回复
// @Summary 发布评论
// @Description 发布评论或回复，评论用户由登录态确定
// @Tags 评论
// @Accept application/json
// @Param data body publishCommentReq true "评论信息"
// @Router /api/v1/app/comment [post]
// @Security Bearer
func (e Comment) PublishComment(c *gin.Context) {
	if err := e.MakeContext(c).MakeOrm().Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "cn"
	}
	current, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}
	var req publishCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		e.Error(400, err, fault.GetErrorMessage("parameter_error", lang))
		return
	}
	s := webService.AppInformationComment{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	id, err := s.PublishComment(int64(current.Id), req.NewsID, req.ParentID, req.Content, lang)
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}
	e.OK(gin.H{"id": id}, fault.GetErrorMessage("comment_published_successfully", lang))
}

// DeleteComment 删除本人发布的评论
// @Summary 删除评论
// @Tags 评论
// @Param id path int true "评论ID"
// @Router /api/v1/app/comment/{id} [delete]
// @Security Bearer
func (e Comment) DeleteComment(c *gin.Context) {
	if err := e.MakeContext(c).MakeOrm().Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "cn"
	}
	current, err := e.currentUser(c)
	if err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		e.Error(400, err, fault.GetErrorMessage("invalid_comment_id", lang))
		return
	}
	s := webService.AppInformationComment{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	if err := s.DeleteComment(int64(current.Id), id, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}
	e.OK(nil, fault.GetErrorMessage("comment_deleted_successfully", lang))
}

// LikeComment 点赞评论
// @Summary 点赞评论
// @Tags 评论
// @Param id path int true "评论ID"
// @Router /api/v1/app/comment/{id}/like [post]
// @Security Bearer
func (e Comment) LikeComment(c *gin.Context) {
	e.changeLike(c, true)
}

// UnlikeComment 取消点赞评论
// @Summary 取消点赞评论
// @Tags 评论
// @Param id path int true "评论ID"
// @Router /api/v1/app/comment/{id}/unlike [post]
// @Security Bearer
func (e Comment) UnlikeComment(c *gin.Context) {
	e.changeLike(c, false)
}

func (e Comment) changeLike(c *gin.Context, increase bool) {
	if err := e.MakeContext(c).MakeOrm().Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "cn"
	}
	if _, err := e.currentUser(c); err != nil {
		e.Error(401, err, fault.GetErrorMessage("not_login", lang))
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		e.Error(400, err, fault.GetErrorMessage("invalid_comment_id", lang))
		return
	}
	s := webService.AppInformationComment{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}
	if err := s.ChangeCommentLike(id, increase, lang); err != nil {
		e.Error(400, err, err.Error())
		return
	}
	if increase {
		e.OK(nil, fault.GetErrorMessage("like_success", lang))
	} else {
		e.OK(nil, fault.GetErrorMessage("cancel_like_success", lang))
	}
}
