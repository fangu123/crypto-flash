package service

import (
	"encoding/json"
	"errors"
	"fmt"
	fault "go-admin/common/errors"
	"io"
	"net/http"
	"time"

	adminModels "go-admin/app/admin/models"

	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type AppPoint struct {
	service.Service
}

// 点位列表项结构
type PointListItem struct {
	ID            int64     `json:"id"`
	CoinId        int64     `json:"coinId"`
	CoinName      string    `json:"coinName"`
	CoinLogo      string    `json:"coinLogo"`
	Type          int64     `json:"type"` // 1-做多，2-做空
	EntryPrice    string    `json:"entryPrice"`
	TakeProfits   string    `json:"takeProfits"`
	StopLoss      string    `json:"stopLoss"`
	ValidDuration int64     `json:"validDuration"`
	ValidUntil    time.Time `json:"validUntil"`
	Status        int64     `json:"status"` // 1-进行中，2-已失效，3-已撤单
	IsNotified    int64     `json:"isNotified"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	IsFavorite    bool      `json:"isFavorite"`
}

// 点位详情结构
type PointDetail struct {
	ID            int64     `json:"id"`
	CoinId        int64     `json:"coinId"`
	CoinName      string    `json:"coinName"`
	CoinLogo      string    `json:"coinLogo"`
	Type          int64     `json:"type"` // 1-做多，2-做空
	EntryPrice    string    `json:"entryPrice"`
	TakeProfits   string    `json:"takeProfits"`
	StopLoss      string    `json:"stopLoss"`
	ValidDuration int64     `json:"validDuration"`
	ValidUntil    time.Time `json:"validUntil"`
	Status        int64     `json:"status"` // 1-进行中，2-已失效，3-已撤单
	IsNotified    int64     `json:"isNotified"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	IsFavorite    bool      `json:"isFavorite"`
}

func (e *AppPoint) GetTradingPointList(pointType int, page int, size int, userId int) ([]PointListItem, int64, error) {
	var points []adminModels.InformationTradingPoint
	var total int64
	var favorites []adminModels.InformationTradingPointFavorite
	var coins []adminModels.InformationCoin

	query := e.Orm.Model(&adminModels.InformationTradingPoint{})

	// 类型筛选
	if pointType > 0 && pointType < 3 {
		query = query.Where("type = ?", pointType)
	}

	// 收藏筛选
	if pointType == 3 && userId > 0 {
		// 先查询用户收藏的点位ID
		var favoriteIds []int64
		if err := e.Orm.Model(&adminModels.InformationTradingPointFavorite{}).
			Where("user_id = ?", userId).Pluck("trading_point_id", &favoriteIds).Error; err != nil {
			return nil, 0, err
		}

		if len(favoriteIds) > 0 {
			query = query.Where("id IN ?", favoriteIds)
		} else {
			// 如果没有收藏，返回空结果
			return []PointListItem{}, 0, nil
		}
	}

	// 只查询进行中的点位
	query = query.Where("status = ?", 1).Order("created_at DESC")

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Find(&points).Error; err != nil {
		return nil, 0, err
	}

	// 获取所有币种
	coinIds := make([]int64, 0)
	for _, point := range points {
		coinIds = append(coinIds, point.CoinId)
	}
	if len(coinIds) > 0 {
		if err := e.Orm.Where("id IN ?", coinIds).Find(&coins).Error; err != nil {
			return nil, 0, err
		}
	}

	// 获取用户收藏记录
	if userId > 0 {
		var pointIds []int64
		for _, point := range points {
			pointIds = append(pointIds, int64(point.Id))
		}
		if len(pointIds) > 0 {
			if err := e.Orm.Where("user_id = ? AND trading_point_id IN ?", userId, pointIds).
				Find(&favorites).Error; err != nil {
				return nil, 0, err
			}
		}
	}

	// 构建返回结果
	var result []PointListItem
	coinMap := make(map[int64]adminModels.InformationCoin)

	// 创建币种映射
	for _, coin := range coins {
		coinMap[int64(coin.Id)] = coin
	}

	// 创建收藏映射
	favoriteMap := make(map[int64]bool)
	for _, favorite := range favorites {
		favoriteMap[favorite.TradingPointId] = true
	}

	// 构建返回数据
	for _, point := range points {
		item := PointListItem{
			ID:            int64(point.Id),
			CoinId:        point.CoinId,
			Type:          point.Type,
			EntryPrice:    point.EntryPrice,
			TakeProfits:   point.TakeProfits,
			StopLoss:      point.StopLoss,
			ValidDuration: point.ValidDuration,
			ValidUntil:    point.ValidUntil,
			Status:        point.Status,
			IsNotified:    point.IsNotified,
			Notes:         point.Notes,
			CreatedAt:     point.CreatedAt,
			UpdatedAt:     point.UpdatedAt,
			IsFavorite:    favoriteMap[int64(point.Id)],
		}

		// 设置币种名称
		if coin, exists := coinMap[point.CoinId]; exists {
			item.CoinName = coin.Symbol
			item.CoinLogo = coin.LogoUrl
		}

		result = append(result, item)
	}

	return result, total, nil
}

func (e *AppPoint) GetTradingPointDetail(pointId int64, userId int, lang string) (*PointDetail, error) {
	var point adminModels.InformationTradingPoint
	var coins []adminModels.InformationCoin
	var isFavorite bool

	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		return nil, err
	}

	// 查询点位详情
	if err := e.Orm.Where("status = 1").First(&point, pointId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("point_not_exist", lang))
		}
		return nil, err
	}

	if user.PointNotificationSettings == 1 {
		// 查询历史的三个点位ID
		var latestPoints []adminModels.InformationTradingPoint
		if err := e.Orm.Where("status = ?", 1).Order("created_at ASC").Limit(3).Pluck("id", &latestPoints).Error; err != nil {
			return nil, err
		}

		// 检查pointId是否在前三位
		isInLatest := false
		for _, id := range latestPoints {
			if int64(id.Id) == pointId {
				isInLatest = true
				break
			}
		}

		if !isInLatest {
			return nil, errors.New(fault.GetErrorMessage("point_not_unlocked", lang))
		}
	}

	// 查询相关币种
	if point.CoinId > 0 {
		if err := e.Orm.First(&coins, point.CoinId).Error; err == nil {
			coins = append(coins, coins[0])
		}
	}

	// 检查用户是否收藏
	if userId > 0 {
		var count int64
		if err := e.Orm.Model(&adminModels.InformationTradingPointFavorite{}).
			Where("user_id = ? AND trading_point_id = ?", userId, pointId).Count(&count).Error; err != nil {
			return nil, err
		}
		isFavorite = count > 0
	}

	// 构建返回结果
	result := &PointDetail{
		ID:            int64(point.Id),
		CoinId:        point.CoinId,
		Type:          point.Type,
		EntryPrice:    point.EntryPrice,
		TakeProfits:   point.TakeProfits,
		StopLoss:      point.StopLoss,
		ValidDuration: point.ValidDuration,
		ValidUntil:    point.ValidUntil,
		Status:        point.Status,
		IsNotified:    point.IsNotified,
		Notes:         point.Notes,
		CreatedAt:     point.CreatedAt,
		UpdatedAt:     point.UpdatedAt,
		IsFavorite:    isFavorite,
	}

	// 设置币种名称
	if point.CoinId > 0 && len(coins) > 0 {
		result.CoinName = coins[0].Symbol
		result.CoinLogo = coins[0].LogoUrl
	}

	return result, nil
}

func (e *AppPoint) ShareTradingPoint(userId, pointId int64, lang string) error {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		return err
	}

	if user.PointNotificationSettings == 1 {
		// 查询最新的三个点位ID
		var latestPoints []adminModels.InformationTradingPoint
		if err := e.Orm.Order("created_at DESC").Limit(3).Pluck("id", &latestPoints).Error; err != nil {
			return err
		}

		// 检查pointId是否在前三位
		isInLatest := false
		for _, id := range latestPoints {
			if int64(id.Id) == pointId {
				isInLatest = true
				break
			}
		}

		if !isInLatest {
			return errors.New(fault.GetErrorMessage("point_not_exist", lang))
		}
	}
	var point adminModels.InformationTradingPoint
	if err := e.Orm.Where("status = 1").First(&point, pointId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("point_not_exist", lang))
		}
		return err
	}

	// 更新分享次数
	//if err := e.Orm.Model(&point).Update("share_count", point.ShareCount+1).Error; err != nil {
	//	return err
	//}

	return nil
}

func (e *AppPoint) FavoriteTradingPoint(userId int64, pointId int64, lang string) error {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		return err
	}

	if user.PointNotificationSettings == 1 {
		// 查询最新的三个点位ID
		var latestPoints []adminModels.InformationTradingPoint
		if err := e.Orm.Order("created_at DESC").Limit(3).Pluck("id", &latestPoints).Error; err != nil {
			return err
		}

		// 检查pointId是否在前三位
		isInLatest := false
		for _, id := range latestPoints {
			if int64(id.Id) == pointId {
				isInLatest = true
				break
			}
		}

		if !isInLatest {
			return errors.New(fault.GetErrorMessage("point_not_exist", lang))
		}
	}

	var point adminModels.InformationTradingPoint
	if err := e.Orm.Where("status = 1").First(&point, pointId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("point_not_exist", lang))
		}
		return err
	}

	// 检查是否已经收藏
	var count int64
	if err := e.Orm.Model(&adminModels.InformationTradingPointFavorite{}).
		Where("user_id = ? AND trading_point_id = ?", userId, pointId).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New(fault.GetErrorMessage("point_already_collected", lang))
	}

	// 创建收藏记录
	favorite := adminModels.InformationTradingPointFavorite{
		UserId:         userId,
		TradingPointId: pointId,
	}

	if err := e.Orm.Create(&favorite).Error; err != nil {
		return err
	}

	return nil
}

func (e *AppPoint) UnfavoriteTradingPoint(userId int64, pointId int64, lang string) error {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		return err
	}

	if user.PointNotificationSettings == 1 {
		// 查询最新的三个点位ID
		var latestPoints []adminModels.InformationTradingPoint
		if err := e.Orm.Order("created_at DESC").Limit(3).Pluck("id", &latestPoints).Error; err != nil {
			return err
		}

		// 检查pointId是否在前三位
		isInLatest := false
		for _, id := range latestPoints {
			if int64(id.Id) == pointId {
				isInLatest = true
				break
			}
		}

		if !isInLatest {
			return errors.New(fault.GetErrorMessage("point_not_exist", lang))
		}
	}
	// 删除收藏记录
	if err := e.Orm.Where("user_id = ? AND trading_point_id = ?", userId, pointId).
		Delete(&adminModels.InformationTradingPointFavorite{}).Error; err != nil {
		return err
	}

	return nil
}

// UnlockAllTradingPoint 解锁全部点位
// @Summary 解锁全部点位
// @Description 解锁全部点位，查询用户是否有相关的NFT，如果有就修改用户表的pointNotificationSettings为2
// @Tags 点位
// @Router /api/v1/app/trading-point/unlock-all [post]
// @Security Bearer
func (e *AppPoint) UnlockAllTradingPoint(userId int64, lang string) error {
	// 查询用户信息
	var user adminModels.InformationUser
	if err := e.Orm.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("user_not_exist", lang))
		}
		return err
	}
	if user.WalletAddress == "" {
		return errors.New(fault.GetErrorMessage("wallet_address_cannot_be_empty", lang))
	}

	// 如果已经是全部解锁状态，直接返回
	if user.PointNotificationSettings == 2 {
		return errors.New(fault.GetErrorMessage("unlock_point", lang))
	}

	// 查询NFT API配置
	var apiConfig adminModels.InformationThirdpartyApikey
	if err := e.Orm.Where("type = ? AND status = ?", 5, 1).First(&apiConfig).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("nft_api_config_not_found", lang))
		}
		return err
	}

	// 这里应该是调用NFT链上查询API，检查用户是否有相关NFT
	// 由于示例中没有具体的NFT查询逻辑，这里简化处理
	// 实际项目中，这里应该调用第三方API查询用户NFT持有情况
	// 链上数据
	// 构建请求URL
	url := fmt.Sprintf("%s/nft/v3/%s/getNFTsForOwner?owner=%s&contractAddresses[]=%s",
		apiConfig.BaseUrl, apiConfig.Apikey, user.WalletAddress, apiConfig.Passphrase)

	// 创建HTTP请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// 发送请求
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 解析响应数据
	var result map[string]interface{}
	if err := json.Unmarshal(all, &result); err != nil {
		return err
	}

	// 检查是否有NFT数据
	if ownedNfts, ok := result["ownedNfts"].([]interface{}); ok && len(ownedNfts) > 0 {
		// 假设查询结果为有NFT，更新用户点位通知设置为全部解锁
		if err := e.Orm.Model(&user).Update("point_notification_settings", 2).Error; err != nil {
			return err
		}
	}

	return nil
}
