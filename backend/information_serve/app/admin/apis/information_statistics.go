package apis

import (
	"fmt"
	"go-admin/app/admin/service"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	coreService "github.com/go-admin-team/go-admin-core/sdk/service"
)

type InformationStatistics struct {
	api.Api
}

// GetStatistics 获取统计数据
// @Summary 获取统计数据
// @Description 获取统计数据
// @Tags 统计
// @Success 200 {object} response.Response{data=StatisticsResp} "{"code": 200, "data": {...}}"
// @Router /api/v1/information/statistics [get]
// @Security Bearer
func (e InformationStatistics) GetStatistics(c *gin.Context) {
	s := service.InformationStatistics{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService((*coreService.Service)(&s.Service)).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	result, err := s.GetStatistics()
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取统计数据失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(result, "获取统计数据成功")
}
