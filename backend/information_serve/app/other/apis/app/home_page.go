package web

import (
	webService "go-admin/app/other/service"
	fault "go-admin/common/errors"

	"github.com/gin-gonic/gin"
)

type HomePage struct {
	AppApi
}

// GetMarketOverview 首页市场概览
// @Summary 首页市场概览
// @Description 获取比特币、以太坊和Solana等的市场数据
// @Tags 首页
// @Router /api/v1/app/home/market-overview [get]
// @Security Bearer
func (e HomePage) GetMarketOverview(c *gin.Context) {
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

	s := webService.AppHomePage{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetMarketOverview(lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_market_overview_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// GetLongShortRatio 首页多空比
// @Summary 首页多空比
// @Description 获取Bybit的多空比数据
// @Tags 首页
// @Param period query string true "时间周期: "5min", "15min", "30min", "1h", "4h", "1d"" default(1h)
// @Router /api/v1/app/home/long-short-ratio [get]
// @Security Bearer
func (e HomePage) GetLongShortRatio(c *gin.Context) {
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

	period := c.Query("period")
	if period == "" {
		period = "1h"
	}
	symbol := c.Query("symbol")
	if symbol == "" {
		symbol = "BTC"
	}

	s := webService.AppHomePage{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetLongShortRatio(period, symbol, lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_market_overview_ratio_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// GetFinancialCalendar 首页财经日历
// @Summary 首页财经日历
// @Description 获取24小时热门财经新闻
// @Tags 首页
// @Router /api/v1/app/home/financial-calendar [get]
// @Security Bearer
func (e HomePage) GetFinancialCalendar(c *gin.Context) {
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

	s := webService.AppHomePage{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetFinancialCalendar(lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_finance_calendar_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// GetNewsFlash 首页新闻快讯
// @Summary 首页新闻快讯
// @Description 获取重要新闻快讯
// @Tags 首页
// @Router /api/v1/app/home/news-flash [get]
// @Security Bearer
func (e HomePage) GetNewsFlash(c *gin.Context) {
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

	s := webService.AppHomePage{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetNewsFlash(lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("get_news_flash_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// GetLatestTradingPoint 首页最新交易点位
// @Summary 首页最新交易点位
// @Description 获取最新的交易点位信息
// @Tags 首页
// @Router /api/v1/app/home/latest-trading-point [get]
// @Security Bearer
func (e HomePage) GetLatestTradingPoint(c *gin.Context) {
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

	s := webService.AppHomePage{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.GetLatestTradingPoint(lang)
	if err != nil {
		e.Error(400, err, err.Error())
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}

// Search 首页搜索功能
// @Summary 首页搜索功能
// @Description 搜索市场概览、多空比数据、点位、财经和快讯
// @Tags 首页
// @Param keyword query string true "搜索关键词"
// @Param type query int false "搜索类型：1-市场概览，2-财经，3-快讯，4-多空比，5-点位，默认为0表示全部" default(0)
// @Router /api/v1/app/home/search [get]
// @Security Bearer
func (e HomePage) Search(c *gin.Context) {
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

	keyword := c.Query("keyword")
	if keyword == "" {
		e.Error(400, nil, fault.GetErrorMessage("keyword_required", lang))
		return
	}

	searchType := c.DefaultQuery("type", "0")

	s := webService.AppHomePage{}
	if err := e.MakeService(&s.Service).Errors; err != nil {
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.Search(keyword, searchType, lang)
	if err != nil {
		e.Error(400, err, fault.GetErrorMessage("search_fail", lang))
		return
	}

	e.OK(result, fault.GetErrorMessage("acquire_success", lang))
}
