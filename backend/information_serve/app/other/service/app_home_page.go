package service

import (
	"encoding/json"
	"errors"
	"fmt"
	adminModels "go-admin/app/admin/models"
	fault "go-admin/common/errors"
	"go-admin/common/redishelper"
	"html"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
)

type AppHomePage struct {
	service.Service
}

// 首页市场概览数据结构
type MarketOverview struct {
	ID                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	Image                        string    `json:"image"`
	CurrentPrice                 float64   `json:"current_price"`
	MarketCap                    float64   `json:"market_cap"`
	MarketCapRank                int       `json:"market_cap_rank"`
	FullyDilutedValuation        float64   `json:"fully_diluted_valuation"`
	TotalVolume                  float64   `json:"total_volume"`
	High24h                      float64   `json:"high_24h"`
	Low24h                       float64   `json:"low_24h"`
	PriceChange24h               float64   `json:"price_change_24h"`
	PriceChangePercentage24h     float64   `json:"price_change_percentage_24h"`
	MarketCapChange24h           float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64   `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64   `json:"circulating_supply"`
	TotalSupply                  float64   `json:"total_supply"`
	MaxSupply                    float64   `json:"max_supply"`
	ATH                          float64   `json:"ath"`
	ATHChangePercentage          float64   `json:"ath_change_percentage"`
	ATHDate                      time.Time `json:"ath_date"`
	ATL                          float64   `json:"atl"`
	ATLChangePercentage          float64   `json:"atl_change_percentage"`
	ATLDate                      time.Time `json:"atl_date"`
	ROI                          struct {
		Times      float64 `json:"times"`
		Currency   string  `json:"currency"`
		Percentage float64 `json:"percentage"`
	} `json:"roi"`
	LastUpdated   time.Time `json:"last_updated"`
	SparklineIn7d struct {
		Price []float64 `json:"price"`
	} `json:"sparkline_in_7d"`
	PriceChangePercentage1hInCurrency float64 `json:"price_change_percentage_1h_in_currency"`
}

// 首页多空比数据结构
type LongShortRatio struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     HomePageLongShortRatio `json:"result"`
	RetExtInfo map[string]interface{} `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type HomePageLongShortRatio struct {
	List           []LongShortRatioItem `json:"list"`
	NextPageCursor string               `json:"nextPageCursor"`
}
type LongShortRatioItem struct {
	Symbol    string `json:"symbol"`
	BuyRatio  string `json:"buyRatio"`
	SellRatio string `json:"sellRatio"`
	Timestamp string `json:"timestamp"`
}

// 首页财经日历数据结构
type FinancialCalendar struct {
	Code            int                   `json:"code"`
	Message         string                `json:"msg"`
	Success         bool                  `json:"success"`
	ServerTimestamp int64                 `json:"serverTimestamp"`
	Data            FinancialCalendarData `json:"data"`
}

type FinancialCalendarData struct {
	Total     int            `json:"total"`
	PageSize  int            `json:"pageSize"`
	PageNum   int            `json:"pageNum"`
	List      []CalendarNews `json:"list"`
	Empty     bool           `json:"empty"`
	TotalPage int            `json:"totalPage"`
	HasMore   bool           `json:"hasMore"`
}

type CalendarNews struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	Content          string   `json:"content"`
	Images           []string `json:"images"`
	IsImportant      bool     `json:"isImportant"`
	PublishTimestamp int64    `json:"publishTimestamp"`
	PublishDate      string   `json:"publishDate"`
	SourceUrl        string   `json:"sourceUrl"`
	Link             string   `json:"link"`
}

func (f *FinancialCalendar) DecodeHTMLContent() {
	for i := range f.Data.List {
		// 首先解码HTML实体
		decodedContent := html.UnescapeString(f.Data.List[i].Content)
		// 然后移除HTML标签
		f.Data.List[i].Content = strings.ReplaceAll(decodedContent, "<p>", "")
		f.Data.List[i].Content = strings.ReplaceAll(f.Data.List[i].Content, "</p>", "")
		f.Data.List[i].Content = strings.ReplaceAll(f.Data.List[i].Content, "<br>", "")
		f.Data.List[i].Content = strings.ReplaceAll(f.Data.List[i].Content, "</br>", "")
		f.Data.List[i].Content = strings.ReplaceAll(f.Data.List[i].Content, `\`, "")
	}
}

// 首页新闻快讯数据结构
type NewsFlash struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    NewsItem `json:"data"`
}

type NewsItem struct {
	Page int    `json:"page"`
	Data []Data `json:"data"`
}

type Data struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Pic        string `json:"pic"`
	Link       string `json:"link"`
	Url        string `json:"url"`
	CreateTime string `json:"create_time"`
}

func (f *NewsFlash) DecodeHTMLContent() {
	for i := range f.Data.Data {
		f.Data.Data[i].Content = html.UnescapeString(f.Data.Data[i].Content)
		// 然后移除HTML标签
		f.Data.Data[i].Content = strings.ReplaceAll(f.Data.Data[i].Content, "\u003cp", "")
		f.Data.Data[i].Content = strings.ReplaceAll(f.Data.Data[i].Content, "\u003e", "")
		f.Data.Data[i].Content = strings.ReplaceAll(f.Data.Data[i].Content, "/p", "")
		f.Data.Data[i].Content = strings.ReplaceAll(f.Data.Data[i].Content, "\u003c", "")
		f.Data.Data[i].Content = strings.ReplaceAll(f.Data.Data[i].Content, "\u003cbr", "")
		f.Data.Data[i].Content = strings.ReplaceAll(f.Data.Data[i].Content, `\`, "")
	}
}

// 获取首页市场概览数据
func (e *AppHomePage) GetMarketOverview(lang string) ([]MarketOverview, error) {
	// 尝试从缓存获取
	cacheKey := "market_overview"
	cachedData, err := redishelper.DefaultRedis.Get(cacheKey)
	if err == nil && cachedData != "" {
		var result []MarketOverview
		if err := json.Unmarshal([]byte(cachedData), &result); err == nil {
			return result, nil
		}
	}
	apikey, _ := e.getApikey(1, lang)
	var coins []adminModels.InformationCoin
	if err := e.Orm.Where("status = 1").Find(&coins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("coin_not_found", lang))
		}
		return nil, err
	}
	//拼接成逗号隔开的字符串
	var coinIds string
	for _, coin := range coins {
		coinIds += coin.Name + ","
	}

	// 缓存未命中，从API获取
	url := apikey.BaseUrl + "?vs_currency=usd&ids=" + coinIds + "&sparkline=true&price_change_percentage=1h"
	//请求头需要加apikey
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-cg-demo-api-key", apikey.Apikey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []MarketOverview
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// 存入缓存，设置1分钟过期
	if data, err := json.Marshal(result); err == nil {
		redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
	}

	return result, nil
}

// 获取首页多空比数据
func (e *AppHomePage) GetLongShortRatio(period, symbol, lang string) (interface{}, error) {
	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("long_short_ratio_%s", symbol+period)
	cachedData, err := redishelper.DefaultRedis.Get(cacheKey)
	if err == nil && cachedData != "" {
		var result LongShortRatio
		if err := json.Unmarshal([]byte(cachedData), &result); err == nil {
			return &result.Result.List[0], nil
		}
	}

	// 缓存未命中，从API获取
	apikey, _ := e.getApikey(2, lang)

	url := fmt.Sprintf(apikey.BaseUrl+"?category=linear&symbol=%s&period=%s&limit=1", symbol+"USDT", period)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result LongShortRatio
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// 存入缓存
	if data, err := json.Marshal(result); err == nil {
		err := redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
		if err != nil {
			return nil, err
		}
	}

	return &result.Result.List[0], nil
}

// 获取首页财经日历数据
func (e *AppHomePage) GetFinancialCalendar(lang string) (interface{}, error) {
	var apiLang string
	if lang == "cn" {
		apiLang = "zh-cn"
	} else {
		apiLang = lang
	}
	// 尝试从缓存获取
	cacheKey := "financial_calendar" + "_" + apiLang
	cachedData, err := redishelper.DefaultRedis.Get(cacheKey)
	if err == nil && cachedData != "" {
		var result FinancialCalendar
		if err := json.Unmarshal([]byte(cachedData), &result); err == nil {
			return &result.Data.List, nil
		}
	}
	apikey, _ := e.getApikey(3, lang)

	// 缓存未命中，从API获取
	url := apikey.BaseUrl + "/newsflash/24h-hot?page=1&size=10&lang=" + apiLang
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result FinancialCalendar
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	// 添加解码处理
	result.DecodeHTMLContent()
	// 存入缓存，设置1小时过期
	if data, err := json.Marshal(result); err == nil {
		err := redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
		if err != nil {
			return nil, err
		}
	}

	return &result.Data.List, nil
}

// 获取首页新闻快讯数据
func (e *AppHomePage) GetNewsFlash(lang string) (interface{}, error) {
	// 尝试从缓存获取
	cacheKey := "news_flash" + "_" + lang
	cachedData, err := redishelper.DefaultRedis.Get(cacheKey)
	if err == nil && cachedData != "" {
		var result NewsFlash
		if err := json.Unmarshal([]byte(cachedData), &result); err == nil {
			return &result.Data.Data, nil
		}
	}
	apikey, _ := e.getApikey(4, lang)

	// 缓存未命中，从API获取
	url := apikey.BaseUrl + "/newsflash/important?lang=" + lang
	//添加请求头
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("api-key", apikey.Apikey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result NewsFlash
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	// 添加解码处理
	result.DecodeHTMLContent()
	// 存入缓存，设置1小时过期
	if data, err := json.Marshal(result); err == nil {
		err := redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
		if err != nil {
			return nil, err
		}
	}

	return &result.Data.Data, nil
}

// 获取首页最新交易点位数据
func (e *AppHomePage) GetLatestTradingPoint(lang string) (*adminModels.InformationTradingPoint, error) {
	var result adminModels.InformationTradingPoint

	// 从数据库查询最新的一条数据
	err := e.Orm.Where("status = 1").Order("created_at desc").First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("point_not_found", lang))
		}
		return nil, err
	}
	var coin adminModels.InformationCoin
	if err := e.Orm.Where("id = ?", result.CoinId).First(&coin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("coin_not_found", lang))
		}
		return nil, err
	}
	result.Name = coin.Symbol

	return &result, nil
}

func (e *AppHomePage) getApikey(t int, lang string) (adminModels.InformationThirdpartyApikey, error) {
	var apikey adminModels.InformationThirdpartyApikey
	if err := e.Orm.Where("type = ?", t).First(&apikey).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apikey, errors.New(fault.GetErrorMessage("apikey_not_found", lang))
		}
		return apikey, err
	}
	return apikey, nil
}

// Search 搜索功能实现
func (e *AppHomePage) Search(keyword, searchType, lang string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// 根据搜索类型执行不同的查询
	switch searchType {
	case "1": // 只搜索市场概览
		marketOverview, err := e.searchMarketOverview(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["marketOverview"] = marketOverview

	case "2": // 只搜索财经
		financialNews, err := e.searchFinancialNews(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["financialNews"] = financialNews

	case "3": // 只搜索快讯
		newsFlash, err := e.searchNewsFlash(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["newsFlash"] = newsFlash

	case "4": // 只搜索多空比
		longShortRatio, err := e.searchLongShortRatio(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["longShortRatio"] = longShortRatio

	case "5": // 只搜索点位
		tradingPoints, err := e.searchTradingPoints(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["tradingPoints"] = tradingPoints

	default: // 搜索全部
		// 搜索市场概览
		marketOverview, err := e.searchMarketOverview(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["marketOverview"] = marketOverview

		// 搜索财经
		financialNews, err := e.searchFinancialNews(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["financialNews"] = financialNews

		// 搜索快讯
		newsFlash, err := e.searchNewsFlash(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["newsFlash"] = newsFlash

		// 搜索多空比
		longShortRatio, err := e.searchLongShortRatio(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["longShortRatio"] = longShortRatio

		// 搜索点位
		tradingPoints, err := e.searchTradingPoints(keyword, lang)
		if err != nil {
			return nil, err
		}
		result["tradingPoints"] = tradingPoints
	}

	return result, nil
}

// searchMarketOverview 搜索市场概览数据
func (e *AppHomePage) searchMarketOverview(keyword, lang string) ([]MarketOverview, error) {
	// 尝试从缓存获取
	cacheKey := "market_overview"
	cachedData, err := redishelper.DefaultRedis.Get(cacheKey)
	if err == nil && cachedData != "" {
		var result []MarketOverview
		if err := json.Unmarshal([]byte(cachedData), &result); err == nil {
			// 过滤包含关键词的数据
			var filteredResult []MarketOverview
			for _, item := range result {
				if strings.Contains(strings.ToLower(item.Name), strings.ToLower(keyword)) ||
					strings.Contains(strings.ToLower(item.Symbol), strings.ToLower(keyword)) {
					filteredResult = append(filteredResult, item)
				}
			}
			return filteredResult, nil
		}
	}

	// 如果缓存中没有，直接调用GetMarketOverview方法获取数据
	return e.GetMarketOverview(lang)
}

// searchLongShortRatio 搜索多空比数据
func (e *AppHomePage) searchLongShortRatio(keyword, lang string) ([]LongShortRatioItem, error) {
	// 定义常见的时间周期
	periods := []string{"1h", "4h", "1d", "7d"}
	var allResults []LongShortRatioItem

	// 尝试搜索不同时间周期的数据
	for _, period := range periods {
		// 尝试从缓存获取
		cacheKey := fmt.Sprintf("long_short_ratio_%s_%s", keyword, period)
		cachedData, err := redishelper.DefaultRedis.Get(cacheKey)

		if err == nil && cachedData != "" {
			var result LongShortRatio
			if err := json.Unmarshal([]byte(cachedData), &result); err == nil {
				// 过滤包含关键词的数据
				for _, item := range result.Result.List {
					if strings.Contains(strings.ToUpper(item.Symbol), strings.ToUpper(keyword)) {
						// 添加周期信息
						itemWithPeriod := item
						allResults = append(allResults, itemWithPeriod)
					}
				}
			}
		} else {
			// 缓存未命中，从API获取
			apiKey, err := e.getApikey(2, lang)
			if err != nil {
				continue // 如果获取API密钥失败，跳过这个周期
			}

			// 构建URL，使用搜索关键词作为symbol
			symbol := keyword + "USDT"
			url := fmt.Sprintf(apiKey.BaseUrl+"?category=linear&symbol=%s&period=%s&limit=1", symbol, period)

			resp, err := http.Get(url)
			if err != nil {
				continue // 请求失败，跳过这个周期
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue // 读取响应失败，跳过这个周期
			}

			var result LongShortRatio
			if err := json.Unmarshal(body, &result); err == nil && result.RetCode == 0 {
				// 过滤包含关键词的数据
				for _, item := range result.Result.List {
					if strings.Contains(strings.ToUpper(item.Symbol), strings.ToUpper(keyword)) {
						// 添加周期信息
						itemWithPeriod := item
						allResults = append(allResults, itemWithPeriod)
					}
				}

				// 存入缓存
				if data, err := json.Marshal(result); err == nil {
					redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
				}
			}
		}
	}

	// 如果没有找到任何数据，尝试获取默认数据
	if len(allResults) == 0 {
		// 尝试获取BTC的数据作为默认
		return e.searchLongShortRatio("BTC", lang)
	}

	return allResults, nil
}

// searchFinancialNews 搜索财经新闻
func (e *AppHomePage) searchFinancialNews(keyword, lang string) ([]adminModels.InformationNews, error) {
	var news []adminModels.InformationNews
	// 使用模糊搜索，搜索标题和内容
	err := e.Orm.Where("status = 2 AND (title LIKE ? OR content LIKE ?)", "%"+keyword+"%", "%"+keyword+"%").Find(&news).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("financial_news_not_found", lang))
		}
		return nil, err
	}

	// 如果没有找到，尝试获取所有财经新闻
	if len(news) == 0 {
		err = e.Orm.Where("status = 2 AND level = 1").Find(&news).Error
		if err != nil {
			return nil, err
		}
	}

	return news, nil
}

// searchNewsFlash 搜索新闻快讯
func (e *AppHomePage) searchNewsFlash(keyword, lang string) ([]adminModels.InformationNews, error) {
	var news []adminModels.InformationNews
	// 使用模糊搜索，搜索标题和内容
	err := e.Orm.Where("status = 2 AND level >= 2 AND (title LIKE ? OR content LIKE ?)", "%"+keyword+"%", "%"+keyword+"%").Find(&news).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("news_flash_not_found", lang))
		}
		return nil, err
	}

	// 如果没有找到，尝试获取所有新闻快讯
	if len(news) == 0 {
		err = e.Orm.Where("status = 2 AND level >= 2").Find(&news).Error
		if err != nil {
			return nil, err
		}
	}

	return news, nil
}

// searchTradingPoints 搜索交易点位
func (e *AppHomePage) searchTradingPoints(keyword, lang string) ([]adminModels.InformationTradingPoint, error) {
	var points []adminModels.InformationTradingPoint
	// 查询关联币种名称包含关键词的点位
	query := e.Orm.Table("information_trading_point").Select("information_trading_point.*, information_coin.symbol as name").
		Joins("LEFT JOIN information_coin ON information_trading_point.coin_id = information_coin.id").
		Where("information_trading_point.status = 1 AND information_coin.symbol LIKE ?", "%"+keyword+"%")
	err := query.Find(&points).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(fault.GetErrorMessage("trading_points_not_found", lang))
		}
		return nil, err
	}

	// 如果没有找到，尝试获取所有点位
	if len(points) == 0 {
		err = e.Orm.Where("status = 1").Find(&points).Error
		if err != nil {
			return nil, err
		}
		// 为点位添加币种名称
		for i := range points {
			var coin adminModels.InformationCoin
			if err := e.Orm.Where("id = ?", points[i].CoinId).First(&coin).Error; err == nil {
				points[i].Name = coin.Symbol
			}
		}
	}

	return points, nil
}
