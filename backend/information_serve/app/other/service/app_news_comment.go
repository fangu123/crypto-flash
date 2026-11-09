package service

import (
	"errors"
	fault "go-admin/common/errors"
	"strings"
	"time"

	adminModels "go-admin/app/admin/models"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
)

const (
	commentStatusNormal  int64 = 1
	commentStatusDeleted int64 = 3
	commentRootParentID  int64 = 1
)

type AppInformationComment struct {
	service.Service
}

// GetCommentList 获取资讯评论分页列表。parentId 默认为 1（顶级评论）。
func (e *AppInformationComment) GetCommentList(newsID, parentID, page, size int64, lang string) ([]adminModels.InformationComment, int64, error) {
	if parentID <= 0 {
		parentID = commentRootParentID
	}
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	if size > 200 {
		size = 200
	}

	var news adminModels.InformationNews
	if err := e.Orm.Select("id").Where("status = ?", 2).First(&news, newsID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.New(fault.GetErrorMessage("news_does_not_exist_or_not_published", lang))
		}
		return nil, 0, err
	}

	query := e.Orm.Model(&adminModels.InformationComment{}).
		Where("news_id = ? AND parent_id = ? AND status = ?", newsID, parentID, commentStatusNormal)
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	list := make([]adminModels.InformationComment, 0)
	offset := (page - 1) * size
	if err := query.Order("is_top ASC, created_at DESC, id DESC").Offset(int(offset)).Limit(int(size)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	if len(list) > 0 {
		for i, comment := range list {
			var replyToUser adminModels.InformationUser
			if comment.ReplyToUserId > 0 {
				if err := e.Orm.Select("id, username, head").Where("id = ?", comment.ReplyToUserId).First(&replyToUser).Error; err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, 0, err
					}
				}
				list[i].ReplyToUserName = maskUsername(replyToUser.Username)
			}
			if comment.ParentId > 1 {
				if err := e.Orm.Select("id, username, head").Where("id = ?", comment.ParentId).First(&replyToUser).Error; err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, 0, err
					}
				}
				list[i].ParentName = maskUsername(replyToUser.Username)
			}
			if comment.UserId > 0 {
				if err := e.Orm.Select("id, username, head").Where("id = ?", comment.UserId).First(&replyToUser).Error; err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						return nil, 0, err
					}
				}
				list[i].UserName = maskUsername(replyToUser.Username)
			}
		}
	}
	return list, total, nil
}

// PublishComment 发布评论。用户ID始终取自已认证的会话，不接受客户端指定。
func (e *AppInformationComment) PublishComment(userID, newsID, parentID int64, content, lang string) (int64, error) {
	content = strings.TrimSpace(content)
	if newsID <= 0 {
		return 0, errors.New(fault.GetErrorMessage("invalid_comment_id", lang))
	}
	if content == "" {
		return 0, errors.New(fault.GetErrorMessage("comment_content_cannot_be_empty", lang))
	}
	if len([]rune(content)) > 2000 {
		return 0, errors.New(fault.GetErrorMessage("comment_content_cannot_be_more_than_2000_characters", lang))
	}
	var news adminModels.InformationNews
	if err := e.Orm.Select("id").Where("status = ?", 2).First(&news, newsID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New(fault.GetErrorMessage("news_does_not_exist_or_not_published", lang))
		}
		return 0, err
	}

	if parentID <= 0 {
		parentID = commentRootParentID
	}
	var replyToUserID int64
	if parentID != commentRootParentID {
		var parent adminModels.InformationComment
		if err := e.Orm.Where("id = ? AND news_id = ? AND status = ?", parentID, newsID, commentStatusNormal).First(&parent).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return 0, errors.New(fault.GetErrorMessage("reply_to_comment_does_not_exist", lang))
			}
			return 0, err
		}
		replyToUserID = parent.UserId
	}

	comment := adminModels.InformationComment{
		NewsId: newsID, UserId: userID, ParentId: parentID,
		ReplyToUserId: replyToUserID, Content: content,
		LikeCount: 0, IsTop: 2, Status: commentStatusNormal,
	}
	if err := e.Orm.Create(&comment).Error; err != nil {
		return 0, err
	}
	return int64(comment.Id), nil
}

// DeleteComment 仅允许作者删除自己的评论，以状态 3 软删除，保留历史数据。
func (e *AppInformationComment) DeleteComment(userID, commentID int64, lang string) error {
	result := e.Orm.Model(&adminModels.InformationComment{}).
		Where("id = ? AND user_id = ? AND status = ?", commentID, userID, commentStatusNormal).
		Updates(map[string]interface{}{"status": commentStatusDeleted, "update_time": time.Now()})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New(fault.GetErrorMessage("comment_does_not_exist_or_no_permission_to_delete", lang))
	}
	return nil
}

// ChangeCommentLike 更新评论点赞数。此项目尚无点赞关系表，因此此处只维护计数。
func (e *AppInformationComment) ChangeCommentLike(commentID int64, increase bool, lang string) error {

	query := e.Orm.Model(&adminModels.InformationComment{}).
		Where("id = ? AND status = ?", commentID, commentStatusNormal)
	if increase {
		query = query.UpdateColumn("like_count", gorm.Expr("like_count + 1"))
	} else {
		query = query.Where("like_count > 0").UpdateColumn("like_count", gorm.Expr("like_count - 1"))
	}
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New(fault.GetErrorMessage("comment_does_not_exist_or_like_status_invalid", lang))
	}
	return nil
}

func maskUsername(name string) string {
	runes := []rune(name)
	length := len(runes)
	if length == 0 {
		return ""
	}
	if length <= 2 {
		return string(runes[:1]) + "**"
	}
	return string(runes[:2]) + "**" + string(runes[length-2:])
}
