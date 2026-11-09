package web

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"

	adminModels "go-admin/app/admin/models"
)

type AppApi struct {
	api.Api
}

func (e *AppApi) currentUserID(c *gin.Context) int {
	return user.GetUserId(c)
}

func (e *AppApi) currentUser(c *gin.Context) (adminModels.InformationUser, error) {
	var u adminModels.InformationUser
	uid := e.currentUserID(c)
	if uid <= 0 {
		return u, errors.New("not login in")
	}
	if err := e.Orm.Model(&adminModels.InformationUser{}).Where("id = ?", uid).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func parsePage(c *gin.Context) (pageIndex int, pageSize int) {
	pageIndex = 1
	pageSize = 10
	if v := c.Query("pageIndex"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			pageIndex = n
		}
	}
	if v := c.Query("pageSize"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 200 {
			pageSize = n
		}
	}
	return
}

func dayKey(t time.Time) string {
	return t.Format("2006-01-02")
}
