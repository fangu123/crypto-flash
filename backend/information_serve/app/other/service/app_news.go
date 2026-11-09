package service

import (
	"encoding/json"
	"errors"
	adminModels "go-admin/app/admin/models"
	fault "go-admin/common/errors"
	"go-admin/common/redishelper"
	"strings"

	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type AppNews struct {
	service.Service
}

// 新闻列表项结构
type NewsListItem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	EnTitle     string `json:"enTitle"`
	Content     string `json:"content"`
	EnContent   string `json:"enContent"`
	Summary     string `json:"summary"`
	CoverImage  string `json:"coverImage"`
	ViewCount   int64  `json:"viewCount"`
	ShareCount  int64  `json:"shareCount"`
	LikeCount   int64  `json:"likeCount"`
	Level       int64  `json:"level"`
	Sentiment   int64  `json:"sentiment"`
	Source      string `json:"source"`
	Logo        string `json:"logo"`
	PublishedAt string `json:"publishedAt"`
	TagName     string `json:"tagName"`
	EnTagName   string `json:"enTagName"`
	CoinName    string `json:"coinName"`
	IsFavorite  bool   `json:"isFavorite"`
}

// 新闻详情结构
type NewsDetail struct {
	ID                           int64   `json:"id"`
	Title                        string  `json:"title"`
	EnTitle                      string  `json:"enTitle"`
	Content                      string  `json:"content"`
	EnContent                    string  `json:"enContent"`
	Summary                      string  `json:"summary"`
	CoverImage                   string  `json:"coverImage"`
	Source                       string  `json:"source"`
	SourceId                     int64   `json:"sourceId"`
	LogoUrl                      string  `json:"logoUrl"`
	RelatedCoinId                int64   `json:"relatedCoinId"`
	RelatedTradingPointId        int64   `json:"relatedTradingPointId"`
	CoinName                     string  `json:"coinName"`
	CoinLogo                     string  `json:"coinLogo"`
	CoinPrice                    float64 `json:"coinPrice"`
	CoinPriceChange              float64 `json:"coinPriceChange"`
	Level                        int64   `json:"level"`
	Sentiment                    int64   `json:"sentiment"`
	IsFeatured                   int64   `json:"isFeatured"`
	ViewCount                    int64   `json:"viewCount"`
	LikeCount                    int64   `json:"likeCount"`
	ShareCount                   int64   `json:"shareCount"`
	Status                       int64   `json:"status"`
	PublishedAt                  string  `json:"publishedAt"`
	CreatedAt                    string  `json:"createdAt"`
	UpdatedAt                    string  `json:"updatedAt"`
	IsFavorite                   bool    `json:"isFavorite"`
	TagName                      string  `json:"tagName"`
	EnTagName                    string  `json:"enTagName"`
	PointNotificationName        string  `json:"pointNotificationName"`
	PointNotificationType        int64   `json:"pointNotificationType"`
	PointNotificationEntryPrice  string  `json:"pointNotificationEntryPrice"`
	PointNotificationTakeProfits string  `json:"pointNotificationTakeProfits"`
}

// 新闻收藏结构
type NewsFavorite struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"userId"`
	NewsId    int64  `json:"newsId"`
	CreatedAt string `json:"createdAt"`
}

// 获取新闻列表
func (e *AppNews) GetNewsList(tagId int64, keyword string, page, size int) ([]NewsListItem, int64, error) {
	var news []adminModels.InformationNews
	var total int64
	var newsRelations []adminModels.InformationNewsTagRelation
	var tags []adminModels.InformationContentTag
	var coins []adminModels.InformationCoin

	query := e.Orm.Model(&adminModels.InformationNews{})

	// 标签筛选
	if tagId > 0 {
		query = query.Joins("LEFT JOIN information_news_tag_relation ON information_news.source_id = information_news_tag_relation.news_id").
			Where("information_news_tag_relation.tag_id = ?", tagId)
	}

	// 关键词搜索
	if keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 只查询已发布的新闻
	query = query.Where("status = ?", 2)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	offset := (page - 1) * size
	if err := query.Offset(offset).Limit(size).Order("published_at DESC").Find(&news).Error; err != nil {
		return nil, 0, err
	}

	// 获取标签关联信息
	newsIds := make([]int64, len(news))
	for i, n := range news {
		newsIds[i] = n.SourceId
	}
	if err := e.Orm.Where("news_id IN ?", newsIds).Find(&newsRelations).Error; err != nil {
		return nil, 0, err
	}

	// 获取所有标签
	tagIds := make([]int64, 0)
	for _, relation := range newsRelations {
		tagIds = append(tagIds, relation.TagId)
	}
	if len(tagIds) > 0 {
		if err := e.Orm.Where("id IN ?", tagIds).Find(&tags).Error; err != nil {
			return nil, 0, err
		}
	}

	// 获取所有相关币种
	coinIds := make([]int64, 0)
	for _, n := range news {
		if n.RelatedCoinId > 0 {
			coinIds = append(coinIds, n.RelatedCoinId)
		}
	}
	if len(coinIds) > 0 {
		if err := e.Orm.Where("id IN ?", coinIds).Find(&coins).Error; err != nil {
			return nil, 0, err
		}
	}

	// 构建返回结果
	var result []NewsListItem
	tagMap := make(map[int64]adminModels.InformationContentTag)
	coinMap := make(map[int64]adminModels.InformationCoin)

	// 创建标签映射
	for _, tag := range tags {
		tagMap[int64(tag.Id)] = tag
	}

	// 创建币种映射
	for _, coin := range coins {
		coinMap[int64(coin.Id)] = coin
	}

	// 构建返回数据
	for _, n := range news {
		item := NewsListItem{
			ID:      int64(n.Id),
			Title:   n.Title,
			EnTitle: n.EnTitle,
			//只展示内容的第一句
			Content:     n.Content,
			EnContent:   n.EnContent,
			Summary:     n.Summary,
			CoverImage:  n.CoverImage,
			Sentiment:   n.Sentiment,
			Source:      n.Source,
			Logo:        n.LogoUrl,
			ViewCount:   n.ViewCount,
			ShareCount:  n.ShareCount,
			LikeCount:   n.LikeCount,
			Level:       n.Level,
			PublishedAt: n.PublishedAt.Format("2006-01-02 15:04:05"),
		}

		// 设置标签名称
		for _, relation := range newsRelations {
			if relation.NewsId == n.SourceId {
				if tag, exists := tagMap[relation.TagId]; exists {
					item.TagName = tag.Name
					item.EnTagName = tag.EnName
					break
				}
			}
		}

		// 设置币种名称
		if n.RelatedCoinId > 0 {
			if coin, exists := coinMap[n.RelatedCoinId]; exists {
				item.CoinName = coin.Symbol
			}
		}

		result = append(result, item)
	}

	return result, total, nil
}

// 获取新闻详情
func (e *AppNews) GetNewsDetail(newsId, userId, PointNotificationSettings int64, lang string) (*NewsDetail, error) {
	var news adminModels.InformationNews
	var newsRelations []adminModels.InformationNewsTagRelation
	var tags []adminModels.InformationContentTag
	var coin adminModels.InformationCoin
	var tradingPoint adminModels.InformationTradingPoint
	var isFavorite bool

	// 查询新闻详情 查询status状态为2的已发布

	if err := e.Orm.Where("(id = ? OR source_id = ?) AND status = 2", newsId, newsId).First(&news).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("news_not_exist", lang))
		}
		return nil, err
	}

	// 增加浏览次数
	if err := e.Orm.Model(&news).Update("view_count", news.ViewCount+1).Error; err != nil {
		return nil, err
	}

	// 查询标签关联信息
	if err := e.Orm.Where("news_id = ?", news.SourceId).Find(&newsRelations).Error; err != nil {
		return nil, err
	}

	// 查询标签
	tagIds := make([]int64, 0)
	for _, relation := range newsRelations {
		tagIds = append(tagIds, relation.TagId)
	}
	if len(tagIds) > 0 {
		if err := e.Orm.Where("id IN ?", tagIds).Find(&tags).Error; err != nil {
			return nil, err
		}
	}

	// 查询相关币种
	if news.RelatedCoinId > 0 {
		if err := e.Orm.First(&coin, news.RelatedCoinId).Error; err != nil {
			return nil, err
		}
	}

	// 检查用户是否收藏
	if userId > 0 {
		var count int64
		if err := e.Orm.Model(&adminModels.InformationNewsFavorite{}).
			Where("user_id = ? AND news_id = ?", userId, news.Id).Count(&count).Error; err != nil {
			return nil, err
		}
		isFavorite = count > 0
	}

	// 构建返回结果
	result := &NewsDetail{
		ID:         int64(news.Id),
		Title:      news.Title,
		EnTitle:    news.EnTitle,
		Content:    news.Content,
		EnContent:  news.EnContent,
		Summary:    news.Summary,
		CoverImage: news.CoverImage,
		Source:     news.Source,
		//SourceId:              news.SourceId,
		LogoUrl:               news.LogoUrl,
		RelatedCoinId:         news.RelatedCoinId,
		RelatedTradingPointId: news.RelatedTradingPointId,
		Level:                 news.Level,
		Sentiment:             news.Sentiment,
		IsFeatured:            news.IsFeatured,
		ViewCount:             news.ViewCount,
		LikeCount:             news.LikeCount,
		ShareCount:            news.ShareCount,
		Status:                news.Status,
		PublishedAt:           news.PublishedAt.Format("2006-01-02 15:04:05"),
		CreatedAt:             news.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:             news.UpdatedAt.Format("2006-01-02 15:04:05"),
		IsFavorite:            isFavorite,
	}

	// 设置标签名称
	tagNames := make([]string, 0)
	enTagNames := make([]string, 0)
	for _, relation := range newsRelations {
		for _, tag := range tags {
			if int64(tag.Id) == relation.TagId {
				tagNames = append(tagNames, tag.Name)
				enTagNames = append(enTagNames, tag.EnName)
				break
			}
		}
	}
	result.TagName = strings.Join(tagNames, ",")
	result.EnTagName = strings.Join(enTagNames, ",")

	// 设置币种名称
	if news.RelatedCoinId > 0 && coin.Id > 0 {
		result.CoinName = coin.Symbol
		result.CoinLogo = coin.LogoUrl
		//缓存获取价格 涨跌幅
		cacheKey := "market_overview" + "_" + coin.Name
		cachedData, err := redishelper.DefaultRedis.Get(cacheKey)
		if err == nil && cachedData != "" {
			var result2 MarketOverview
			if err2 := json.Unmarshal([]byte(cachedData), &result2); err2 == nil {
				result.CoinPrice = result2.CurrentPrice
				result.CoinPriceChange = result2.PriceChangePercentage24h
			}
		}
	}
	if PointNotificationSettings == 2 { //才可以显示点位信息
		if news.RelatedTradingPointId > 0 {
			if err := e.Orm.Where("id = ?", news.RelatedTradingPointId).Find(&tradingPoint).Error; err != nil {
				return nil, err
			}
			var pointCoin adminModels.InformationCoin
			if err := e.Orm.Where("id = ?", tradingPoint.CoinId).Find(&pointCoin).Error; err != nil {
				return nil, err
			}
			result.PointNotificationEntryPrice = tradingPoint.EntryPrice
			result.PointNotificationTakeProfits = tradingPoint.TakeProfits
			result.PointNotificationType = tradingPoint.Type
			result.PointNotificationName = pointCoin.Symbol
		}
	}

	return result, nil
}

// 分享新闻
func (e *AppNews) ShareNews(newsId int64, lang string) error {
	var news adminModels.InformationNews
	if err := e.Orm.Where("status = 2").First(&news, newsId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("news_not_exist", lang))
		}
		return err
	}

	// 增加分享次数
	if err := e.Orm.Model(&news).Update("share_count", news.ShareCount+1).Error; err != nil {
		return err
	}

	return nil
}

// 收藏新闻
func (e *AppNews) FavoriteNews(userId, newsId int64, lang string) error {
	if userId <= 0 {
		return errors.New(fault.GetErrorMessage("user_id_invalid", lang))
	}

	var news adminModels.InformationNews
	if err := e.Orm.Where("status = 2").First(&news, newsId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("news_not_exist", lang))
		}
		return err
	}

	// 检查是否已经收藏
	var count int64
	if err := e.Orm.Model(&adminModels.InformationNewsFavorite{}).
		Where("user_id = ? AND news_id = ?", userId, newsId).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New(fault.GetErrorMessage("news_already_collected", lang))
	}

	// 创建收藏记录
	favorite := adminModels.InformationNewsFavorite{
		UserId: userId,
		NewsId: newsId,
	}

	if err := e.Orm.Create(&favorite).Error; err != nil {
		return err
	}

	return nil
}

// 取消收藏新闻
func (e *AppNews) UnfavoriteNews(userId, newsId int64, lang string) error {
	if userId <= 0 {
		return errors.New(fault.GetErrorMessage("user_id_invalid", lang))
	}
	var news adminModels.InformationNews
	if err := e.Orm.Where("status = 2").First(&news, newsId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(fault.GetErrorMessage("news_not_exist", lang))
		}
		return err
	}
	// 删除收藏记录
	if err := e.Orm.Where("user_id = ? AND news_id = ?", userId, newsId).
		Delete(&adminModels.InformationNewsFavorite{}).Error; err != nil {
		return err
	}

	return nil
}

func (e *AppNews) GetTags() ([]adminModels.InformationContentTag, error) {
	var tags []adminModels.InformationContentTag
	if err := e.Orm.Where("status = 1").Find(&tags).Error; err != nil {
		return tags, err
	}
	return tags, nil
}
