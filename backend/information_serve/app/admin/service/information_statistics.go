package service

import (
	"database/sql"
	"errors"
	"sync"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/service"
)

type InformationStatistics struct {
	service.Service
	cacheLock   sync.RWMutex
	cacheData   interface{}
	cacheTime   time.Time
	cacheExpire time.Duration
}

// GetStatistics 获取统计数据
func (e *InformationStatistics) GetStatistics() (interface{}, error) {
	// 检查缓存是否有效
	if e.isCacheValid() {
		e.cacheLock.RLock()
		defer e.cacheLock.RUnlock()
		return e.cacheData, nil
	}

	// 获取今天开始的时间
	today := time.Now().Truncate(24 * time.Hour)

	// 获取最近7天的日期
	var dates []string
	for i := 6; i >= 0; i-- {
		date := today.AddDate(0, 0, -i).Format("2006-01-02")
		dates = append(dates, date)
	}

	// 设置缓存过期时间（例如5分钟）
	e.cacheExpire = 5 * time.Minute

	// 1. 统计快讯数据
	newsTotal := int64(0)
	newsToday := int64(0)
	newsTrend := make([]int64, 7)
	for i := range newsTrend {
		newsTrend[i] = 0
	}

	var newsCount []struct {
		Date  string `json:"date"`
		Count int64  `json:"count"`
	}

	if e.Orm == nil {
		return nil, errors.New("database connection not initialized")
	}

	// 总数量
	e.Orm.Model(&models.InformationNews{}).Count(&newsTotal)

	// 今日数量
	e.Orm.Model(&models.InformationNews{}).Where("DATE(created_at) = ?", today.Format("2006-01-02")).Count(&newsToday)

	// 曲线数据（最近7天）
	e.Orm.Model(&models.InformationNews{}).
		Select("DATE_FORMAT(created_at, '%Y-%m-%d') as date, COUNT(*) as count").
		Where("created_at >= ?", today.AddDate(0, 0, -6)).
		Group("DATE_FORMAT(created_at, '%Y-%m-%d')").
		Order("date ASC").
		Scan(&newsCount)

	// 填充趋势数据
	for i, date := range dates {
		for _, item := range newsCount {
			if item.Date == date {
				newsTrend[i] = item.Count
				break
			}
		}
	}

	// 2. 统计点位数据
	pointTotal := int64(0)
	pointToday := int64(0)
	pointTrend := make([]int64, 7)
	for i := range pointTrend {
		pointTrend[i] = 0
	}

	var pointCount []struct {
		Date  string `json:"date"`
		Count int64  `json:"count"`
	}

	// 总数量
	e.Orm.Model(&models.InformationTradingPoint{}).Count(&pointTotal)

	// 今日数量
	e.Orm.Model(&models.InformationTradingPoint{}).Where("DATE(created_at) = ?", today.Format("2006-01-02")).Count(&pointToday)

	// 曲线数据（最近7天）
	e.Orm.Model(&models.InformationTradingPoint{}).
		Select("DATE_FORMAT(created_at, '%Y-%m-%d') as date, COUNT(*) as count").
		Where("created_at >= ?", today.AddDate(0, 0, -6)).
		Group("DATE_FORMAT(created_at, '%Y-%m-%d')").
		Order("date ASC").
		Scan(&pointCount)

	// 填充趋势数据
	for i, date := range dates {
		for _, item := range pointCount {
			if item.Date == date {
				pointTrend[i] = item.Count
				break
			}
		}
	}

	// 3. 统计用户数据
	userTotal := int64(0)
	userToday := int64(0)

	e.Orm.Model(&models.InformationUser{}).Count(&userTotal)
	e.Orm.Model(&models.InformationUser{}).Where("DATE(created_at) = ?", today.Format("2006-01-02")).Count(&userToday)

	// 4. 统计阅读次数数据
	readTotal := int64(0)
	readToday := int64(0)

	var readTotalPtr sql.NullInt64
	e.Orm.Model(&models.InformationNews{}).Select("COALESCE(SUM(view_count), 0)").Scan(&readTotalPtr)
	readTotal = readTotalPtr.Int64

	var readTodayPtr sql.NullInt64
	e.Orm.Model(&models.InformationNews{}).Select("COALESCE(SUM(view_count), 0)").Where("DATE(created_at) = ?", today.Format("2006-01-02")).Scan(&readTodayPtr)
	readToday = readTodayPtr.Int64
	// 组装返回数据
	result := map[string]interface{}{
		"news": map[string]interface{}{
			"total": newsTotal,
			"today": newsToday,
			"trend": make([]map[string]interface{}, 7), // 初始化为7个元素
		},
		"tradingPoint": map[string]interface{}{
			"total": pointTotal,
			"today": pointToday,
			"trend": make([]map[string]interface{}, 7), // 初始化为7个元素
		},
		"user": map[string]interface{}{
			"total": userTotal,
			"today": userToday,
		},
		"readCount": map[string]interface{}{
			"total": readTotal,
			"today": readToday,
		},
	}

	// 填充快讯趋势数据
	for i, date := range dates {
		count := int64(0)
		for _, item := range newsCount {
			if item.Date == date {
				count = item.Count
				break
			}
		}
		result["news"].(map[string]interface{})["trend"].([]map[string]interface{})[i] = map[string]interface{}{
			"date":  date,
			"count": count,
		}
	}

	// 填充点位趋势数据
	for i, date := range dates {
		count := int64(0)
		for _, item := range pointCount {
			if item.Date == date {
				count = item.Count
				break
			}
		}
		result["tradingPoint"].(map[string]interface{})["trend"].([]map[string]interface{})[i] = map[string]interface{}{
			"date":  date,
			"count": count,
		}
	}

	// 更新缓存
	e.updateCache(result)

	return result, nil
}

// isCacheValid 检查缓存是否有效
func (e *InformationStatistics) isCacheValid() bool {
	e.cacheLock.RLock()
	defer e.cacheLock.RUnlock()

	// 如果缓存为空或已过期，返回false
	if e.cacheData == nil || time.Since(e.cacheTime) > e.cacheExpire {
		return false
	}
	return true
}

// updateCache 更新缓存
func (e *InformationStatistics) updateCache(data interface{}) {
	e.cacheLock.Lock()
	defer e.cacheLock.Unlock()

	e.cacheData = data
	e.cacheTime = time.Now()
}

// ClearCache 清除缓存
func (e *InformationStatistics) ClearCache() {
	e.cacheLock.Lock()
	defer e.cacheLock.Unlock()

	e.cacheData = nil
	e.cacheTime = time.Time{}
}
