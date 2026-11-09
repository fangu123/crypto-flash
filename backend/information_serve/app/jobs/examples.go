package jobs

import (
	"encoding/json"
	"errors"
	"fmt"
	adminModels "go-admin/app/admin/models"
	"go-admin/app/other/service"
	"go-admin/common/redishelper"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

// InitJob
// 需要将定义的struct 添加到字典中；
// 字典 key 可以配置到 自动任务 调用目标 中；
func InitJob() {
	jobList = map[string]JobExec{
		"ExamplesOne":       ExamplesOne{},
		"MarketOverview":    MarketOverview{},    //5分钟一次 首页市场概览
		"LongShortRatio":    LongShortRatio{},    //5分钟一次 首页多空比
		"FinancialCalendar": FinancialCalendar{}, //5分钟一次 首页财经日历
		"NewsFlash":         NewsFlash{},         //5分钟一次 首页重要新闻快讯 存数据库  需要内容打标签
		"OdailyNews":        OdailyNews{},        //odaily资讯
		"TheblockbeatsNews": TheblockbeatsNews{}, //blockbeats资讯
		//定时查询用户地址是否有相应的NFT
		"QueryUserAddressNFT": QueryUserAddressNFT{},
		//定时查询点位是否过期 如果过期 就需要把标题（点位币种名称 类型做多还是多空） 内容（点位币种名称 类型做多还是多空 入场价格 止盈价格数组 止损价格 有效时长 到期时间） 需要写入到information_notification表标题 内容 type为4。 并且向information_notification_user表 插入信息 用户id是开启了点位状态设置且点位权限是所有的用户NotificationSettings有含3的用户且point_notification_settings为2
		"QueryPoint": QueryPoint{},
	}
}

// ExamplesOne
// 新添加的job 必须按照以下格式定义，并实现Exec函数
type ExamplesOne struct {
}

func (t ExamplesOne) Exec(arg interface{}) error {
	str := time.Now().Format(timeFormat) + " [INFO] JobCore ExamplesOne exec success"
	// TODO: 这里需要注意 Examples 传入参数是 string 所以 arg.(string)；请根据对应的类型进行转化；
	switch arg.(type) {

	case string:
		if arg.(string) != "" {
			fmt.Println("string", arg.(string))
			fmt.Println(str, arg.(string))
		} else {
			fmt.Println("arg is nil")
			fmt.Println(str, "arg is nil")
		}
		break
	}

	return nil
}

type MarketOverview struct {
}

func (m MarketOverview) Exec(arg interface{}) error {
	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return errors.New("database is nil")
	}

	cacheKey := "market_overview"
	apikey, _ := getApikey(1)
	var coins []adminModels.InformationCoin
	if err := db.Where("status = 1").Find(&coins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Coin not found")
		}
		return err
	}
	//拼接成逗号隔开的字符串
	var coinIds strings.Builder
	for i, coin := range coins {
		coinIds.WriteString(coin.Name)
		if i < len(coins)-1 {
			coinIds.WriteString(",")
		}
	}

	url := apikey.BaseUrl + "?vs_currency=usd&ids=" + coinIds.String() + "&sparkline=true&price_change_percentage=1h"
	//请求头需要加apikey
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("x-cg-demo-api-key", apikey.Apikey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result []service.MarketOverview
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	// 存入缓存，设置1分钟过期
	if data, err := json.Marshal(result); err == nil {
		redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
		for _, v := range result {
			// 将MarketOverview对象序列化为JSON字符串后再存入Redis
			if data, err := json.Marshal(v); err == nil {
				err := redishelper.DefaultRedis.Set(cacheKey+"_"+v.ID, string(data), -1)
				if err != nil {
					//控制台输出日志
					log.Error(err)
					return err
				}
			} else {
				log.Error(err)
				return err
			}
		}
	}
	return nil
}

type LongShortRatio struct {
}

func (l LongShortRatio) Exec(arg interface{}) error {

	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return errors.New("database is nil")
	}
	var coins []adminModels.InformationCoin
	if err := db.Where("status = 1").Find(&coins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Coin not found")
		}
		return err
	}
	apikey, _ := getApikey(2)
	for _, coin := range coins {
		for _, period := range []string{"5min", "15min", "30min", "1h", "4h", "1d"} {
			cacheKey := fmt.Sprintf("long_short_ratio_%s", coin.Symbol+period)

			url := fmt.Sprintf(apikey.BaseUrl+"?category=linear&symbol=%s&period=%s&limit=1", coin.Symbol+"USDT", period)
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			var result service.LongShortRatio
			if err := json.Unmarshal(body, &result); err != nil {
				return err
			}

			// 存入缓存
			if data, err := json.Marshal(result); err == nil {
				err := redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

type FinancialCalendar struct {
}

func (f FinancialCalendar) Exec(arg interface{}) error {

	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return errors.New("database is nil")
	}

	// 获取所有启用的标签
	var tags []adminModels.InformationContentTag
	if err := db.Where("status = ?", 1).Find(&tags).Error; err != nil {
		return err
	}
	lang := []string{"zh-cn", "en"}

	for _, l := range lang {
		cacheKey := "financial_calendar" + "_" + l

		apikey, err := getApikey(3)
		if err != nil {
			return fmt.Errorf("获取API key失败: %v", err)
		}

		if apikey.BaseUrl == "" {
			return errors.New("API key的BaseUrl为空，无法构建请求URL")
		}

		url := apikey.BaseUrl + "/newsflash/24h-hot?page=1&size=10&lang=" + l

		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("HTTP请求失败: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var result service.FinancialCalendar
		if err := json.Unmarshal(body, &result); err != nil {
			return err
		}
		// 添加解码处理
		result.DecodeHTMLContent()
		// 存入缓存，设置1小时过期
		if data, err := json.Marshal(result); err == nil {
			err := redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
			if err != nil {
				return err
			}
		}

		if err := json.Unmarshal(body, &result); err != nil {
			return err
		}
		// 添加解码处理
		result.DecodeHTMLContent()
		if err := processFinancialCalendarData(db, result, tags, apikey.Passphrase, "economic", l, 2); err != nil {
			return err
		}
	}

	return nil
}

type NewsFlash struct {
}

func (n NewsFlash) Exec(arg interface{}) error {
	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return errors.New("database is nil")
	}
	lang := []string{"cn", "en"}

	for _, l := range lang {
		cacheKey := "news_flash" + "_" + l
		apikey, _ := getApikey(4)

		url := apikey.BaseUrl + "/newsflash/important?lang=" + l
		//添加请求头
		//请求头需要加apikey
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Set("api-key", apikey.Apikey)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var result service.NewsFlash
		if err := json.Unmarshal(body, &result); err != nil {
			return err
		}
		// 添加解码处理
		result.DecodeHTMLContent()
		// 存入缓存，设置1小时过期
		if data, err := json.Marshal(result); err == nil {
			if result.Message != "" {
				return errors.New(result.Message)
			}
			err := redishelper.DefaultRedis.Set(cacheKey, string(data), -1)
			if err != nil {
				return err
			}
		}

		// 获取所有启用的标签
		var tags []adminModels.InformationContentTag
		if err := db.Where("status = ?", 1).Find(&tags).Error; err != nil {
			return err
		}

		// 处理每条新闻
		for _, news := range result.Data.Data {
			var count int64
			db.Model(&adminModels.InformationNews{}).Where("source_id = ? AND source = ?", news.ID, "theblockbeats").Count(&count)

			if count > 0 {
				var existRecord adminModels.InformationNews
				if err := db.Where("source_id = ? AND source = ?", news.ID, "theblockbeats").First(&existRecord).Error; err != nil {
					return err
				}

				if l == "en" && existRecord.EnTitle == "" { //英文 中文标题
					//修改新闻里面的英文内容
					existRecord.EnTitle = news.Title
					existRecord.EnContent = news.Content
					if err := db.Save(&existRecord).Error; err != nil {
						return err
					}
				}
				//中文 英文标题
				if l == "cn" && existRecord.Title == "" {
					//修改新闻里面的英文内容
					existRecord.Title = news.Title
					existRecord.Content = news.Content
					if err := db.Save(&existRecord).Error; err != nil {
						return err
					}
				}
				continue
			}
			// 加载上海时区
			shanghaiLoc, err := time.LoadLocation("Asia/Shanghai")
			if err != nil {
				return fmt.Errorf("failed to load Shanghai location: %v", err)
			}

			// 解析时间字符串并转换为上海时区
			createTime, err := time.ParseInLocation("2006-01-02 15:04:05", news.CreateTime, shanghaiLoc)
			if err != nil {
				return fmt.Errorf("failed to parse time: %v", err)
			}

			// 创建新闻记录
			newsRecord := adminModels.InformationNews{
				Summary:     news.Link, //站内链接
				CoverImage:  news.Url,  //外链接
				PublishedAt: createTime,
				SourceId:    int64(news.ID),
				Source:      "theblockbeats",
				LogoUrl:     apikey.Passphrase, //第三方logo
				Status:      2,
			}
			if l == "en" {
				newsRecord.EnTitle = news.Title
				newsRecord.EnContent = news.Content
			} else {
				newsRecord.Title = news.Title
				newsRecord.Content = news.Content
			}

			// 保存新闻记录
			if err := db.Create(&newsRecord).Error; err != nil {
				return err
			}

			// 检查每个标签是否匹配新闻标题或内容
			for _, tag := range tags {
				// 检查标签名称是否出现在标题或内容中
				if strings.Contains(strings.ToLower(news.Title), strings.ToLower(tag.Name)) ||
					strings.Contains(strings.ToLower(news.Content), strings.ToLower(tag.Name)) ||
					strings.Contains(strings.ToLower(news.Title), strings.ToLower(tag.Code)) ||
					strings.Contains(strings.ToLower(news.Content), strings.ToLower(tag.Code)) {

					// 创建标签关联
					tagRelation := adminModels.InformationNewsTagRelation{
						NewsId: int64(news.ID),
						TagId:  int64(tag.Id),
					}

					// 保存标签关联
					if err := db.Create(&tagRelation).Error; err != nil {
						return err
					}
				}
			}

		}
	}

	return nil
}

type OdailyNews struct {
}

func (o OdailyNews) Exec(arg interface{}) error {
	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return errors.New("database is nil")
	}

	// 获取所有启用的标签
	var tags []adminModels.InformationContentTag
	if err := db.Where("status = ?", 1).Find(&tags).Error; err != nil {
		return err
	}
	// 获取API密钥
	apikey, _ := getApikey(3)
	lang := []string{"zh-cn", "en"}

	for _, l := range lang {
		// API配置
		apis := []struct {
			name string
			url  string
		}{
			{"ai", apikey.BaseUrl + "/newsflash/ai?lang=" + l},
			{"prediction-market", apikey.BaseUrl + "/newsflash/prediction-market?lang=" + l},
			{"on-chain-data", apikey.BaseUrl + "/newsflash/on-chain-data?lang=" + l},
			{"market-report", apikey.BaseUrl + "/newsflash/market-report?lang=" + l},
			{"exchange-announcement", apikey.BaseUrl + "/newsflash/exchange-announcement?lang=" + l},
			{"project-trend", apikey.BaseUrl + "/newsflash/project-trend?lang=" + l},
			{"celebrity-opinion", apikey.BaseUrl + "/newsflash/celebrity-opinion?lang=" + l},
			{"crypto-stock", apikey.BaseUrl + "/newsflash/crypto-stock?lang=" + l},
			{"funding", apikey.BaseUrl + "/newsflash/funding?lang=" + l},
			{"macro-policy", apikey.BaseUrl + "/newsflash/macro-policy?lang=" + l},
		}

		// 处理每个API
		for _, api := range apis {

			// 创建HTTP请求
			req, err := http.NewRequest("GET", api.url, nil)
			if err != nil {
				return err
			}
			// 发送请求
			//client := &http.Client{Timeout: 30 * time.Second}
			//resp, err := client.Do(req)
			//if err != nil {
			//	return err
			//}
			//defer resp.Body.Close()
			// 发送请求，添加重试机制
			maxRetries := 3
			retryDelay := time.Second * 2
			var resp *http.Response
			for i := 0; i < maxRetries; i++ {
				client := &http.Client{
					Timeout: time.Minute, // 增加超时时间到1分钟
				}
				resp, err = client.Do(req)
				if err == nil {
					break
				}
				if i < maxRetries-1 {
					time.Sleep(retryDelay)
				}
			}
			if err != nil {
				return fmt.Errorf("request failed after %d retries: %v", maxRetries, err)
			}
			defer resp.Body.Close()

			// 读取响应
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			var result service.FinancialCalendar
			if err := json.Unmarshal(body, &result); err != nil {
				return err
			}
			// 添加解码处理
			result.DecodeHTMLContent()
			if err := processFinancialCalendarData(db, result, tags, apikey.Passphrase, api.name, l, 1); err != nil {
				return err
			}
		}
	}

	return nil
}

type TheblockbeatsNews struct {
}

func (t TheblockbeatsNews) Exec(arg interface{}) error {
	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return errors.New("database is nil")
	}

	// 获取所有启用的标签
	var tags []adminModels.InformationContentTag
	if err := db.Where("status = ?", 1).Find(&tags).Error; err != nil {
		return err
	}
	// 获取API密钥
	apikey, _ := getApikey(4)

	lang := []string{"cn", "en"}

	for _, l := range lang {
		// API配置
		apis := []struct {
			name string
			url  string
		}{
			{"on-chain-data", apikey.BaseUrl + "/newsflash/onchain?lang=" + l},
			{"funding", apikey.BaseUrl + "/newsflash/financing?lang=" + l},
			{"prediction-market", apikey.BaseUrl + "/newsflash/prediction?lang=" + l},
			{"ai", apikey.BaseUrl + "/newsflash/ai?lang=" + l},
		}

		// 处理每个API
		for _, api := range apis {

			// 创建HTTP请求
			req, err := http.NewRequest("GET", api.url, nil)
			if err != nil {
				return err
			}
			req.Header.Set("api-key", apikey.Apikey)

			// 发送请求
			client := &http.Client{Timeout: 30 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// 读取响应
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// 根据API类型解析数据
			var result service.NewsFlash
			if err := json.Unmarshal(body, &result); err != nil {
				return err
			}
			// 添加解码处理
			result.DecodeHTMLContent()
			// 处理NewsFlash类型数据
			if err := processNewsFlashData(db, result, tags, apikey.Passphrase, api.name, l); err != nil {
				return err
			}

		}
	}

	return nil
}

// 处理NewsFlash类型数据
func processNewsFlashData(db *gorm.DB, newsFlash service.NewsFlash, tags []adminModels.InformationContentTag, logoUrl, name, lang string) error {
	// 处理每条新闻
	for _, news := range newsFlash.Data.Data {
		var count int64
		db.Model(&adminModels.InformationNews{}).Where("source_id = ? AND source = ?", news.ID, "theblockbeats").Count(&count)

		if count > 0 {
			var existRecord adminModels.InformationNews
			if err := db.Where("source_id = ? AND source = ?", news.ID, "theblockbeats").First(&existRecord).Error; err != nil {
				return err
			}

			if lang == "en" && existRecord.EnTitle == "" { //英文 中文标题
				//修改新闻里面的英文内容
				existRecord.EnTitle = news.Title
				existRecord.EnContent = news.Content
				if err := db.Save(&existRecord).Error; err != nil {
					return err
				}
			}
			//中文 英文标题
			if lang == "cn" && existRecord.Title == "" {
				//修改新闻里面的英文内容
				existRecord.Title = news.Title
				existRecord.Content = news.Content
				if err := db.Save(&existRecord).Error; err != nil {
					return err
				}
			}
			continue
		}
		// 加载上海时区
		shanghaiLoc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return fmt.Errorf("failed to load Shanghai location: %v", err)
		}

		// 解析时间字符串并转换为上海时区
		createTime, err := time.ParseInLocation("2006-01-02 15:04:05", news.CreateTime, shanghaiLoc)
		if err != nil {
			return fmt.Errorf("failed to parse time: %v", err)
		}
		// 创建新闻记录
		newsRecord := adminModels.InformationNews{
			Title:       news.Title,
			Content:     news.Content,
			Summary:     news.Link, //站内链接
			CoverImage:  news.Url,  //外链接
			PublishedAt: createTime,
			SourceId:    int64(news.ID),
			Source:      "theblockbeats",
			LogoUrl:     logoUrl, //第三方logo
			Status:      1,
		}

		// 保存新闻记录
		if err := db.Create(&newsRecord).Error; err != nil {
			return err
		}

		// 检查每个标签是否匹配新闻标题或内容
		for _, tag := range tags {
			// 检查标签名称是否出现在标题或内容中
			if strings.Contains(strings.ToLower(news.Title), strings.ToLower(tag.Name)) ||
				strings.Contains(strings.ToLower(news.Content), strings.ToLower(tag.Name)) ||
				strings.Contains(strings.ToLower(news.Title), strings.ToLower(tag.Code)) ||
				strings.Contains(strings.ToLower(news.Content), strings.ToLower(tag.Code)) ||
				strings.ToLower(name) == strings.ToLower(tag.Code) {
				// 创建标签关联
				tagRelation := adminModels.InformationNewsTagRelation{
					NewsId: int64(news.ID),
					TagId:  int64(tag.Id),
				}

				// 保存标签关联
				if err := db.Create(&tagRelation).Error; err != nil {
					return err
				}
			}
		}

	}
	return nil
}

// 处理FinancialCalendar类型数据
func processFinancialCalendarData(db *gorm.DB, financialCalendar service.FinancialCalendar, tags []adminModels.InformationContentTag, logoUrl, name, lang string, status int64) error {
	// 处理每条新闻
	for _, news := range financialCalendar.Data.List {

		var count int64
		db.Model(&adminModels.InformationNews{}).Where("source_id = ? AND source = ?", news.ID, "odaily").Count(&count)

		if count > 0 {
			var existRecord adminModels.InformationNews
			if err := db.Where("source_id = ? AND source = ?", news.ID, "odaily").First(&existRecord).Error; err != nil {
				return err
			}
			if status == 2 && existRecord.Status != 2 {
				existRecord.Status = status
				if err := db.Save(&existRecord).Error; err != nil {
					return err
				}
			}

			if lang == "en" && existRecord.EnTitle == "" { //英文 中文标题
				//修改新闻里面的英文内容
				existRecord.EnTitle = news.Title
				existRecord.EnContent = news.Content
				if err := db.Save(&existRecord).Error; err != nil {
					return err
				}
			}
			//中文 英文标题
			if lang == "cn" && existRecord.Title == "" {
				//修改新闻里面的英文内容
				existRecord.Title = news.Title
				existRecord.Content = news.Content
				if err := db.Save(&existRecord).Error; err != nil {
					return err
				}
			}

			continue
		}

		level := 1
		if news.IsImportant == true {
			level = 2
		}

		// 创建新闻记录
		newsRecord := adminModels.InformationNews{
			Title:       news.Title,
			Content:     news.Content,
			Summary:     news.Link,      //站内链接
			CoverImage:  news.SourceUrl, //外链接
			PublishedAt: time.Unix(news.PublishTimestamp/1000, 0),
			SourceId:    int64(news.ID),
			Source:      "odaily",
			LogoUrl:     logoUrl, //第三方logo
			Status:      status,
			Level:       int64(level),
		}

		// 保存新闻记录
		if err := db.Create(&newsRecord).Error; err != nil {
			return err
		}

		// 检查每个标签是否匹配新闻标题或内容
		for _, tag := range tags {
			// 检查标签名称是否出现在标题或内容中
			if strings.Contains(strings.ToLower(news.Title), strings.ToLower(tag.Name)) ||
				strings.Contains(strings.ToLower(news.Content), strings.ToLower(tag.Name)) ||
				strings.Contains(strings.ToLower(news.Title), strings.ToLower(tag.Code)) ||
				strings.Contains(strings.ToLower(news.Content), strings.ToLower(tag.Code)) ||
				strings.ToLower(name) == strings.ToLower(tag.Code) {

				// 创建标签关联
				tagRelation := adminModels.InformationNewsTagRelation{
					NewsId: int64(news.ID),
					TagId:  int64(tag.Id),
				}

				// 保存标签关联
				if err := db.Create(&tagRelation).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func getApikey(t int) (adminModels.InformationThirdpartyApikey, error) {
	var apikey adminModels.InformationThirdpartyApikey
	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return apikey, errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return apikey, errors.New("database is nil")
	}

	if err := db.Where("type = ?", t).First(&apikey).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apikey, errors.New("API key not found")
		}
		return apikey, err
	}
	return apikey, nil
}

type QueryUserAddressNFT struct {
}

func (q QueryUserAddressNFT) Exec(arg interface{}) error {
	//查询information_user  wallet_address不为空 且point_notification_settings为2的用户列表
	var users []adminModels.InformationUser

	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return errors.New("database is nil")
	}
	apikey, _ := getApikey(5)

	db.Where("wallet_address IS NOT NULL AND point_notification_settings = ?", 2).Find(&users)
	for _, user := range users {
		//通过地址查询链上的nft 如果没有就修改point_notification_settings为1
		// 构建请求URL
		url := fmt.Sprintf("%s/nft/v3/%s/getNFTsForOwner?owner=%s&contractAddresses[]=%s",
			apikey.BaseUrl, apikey.Apikey, user.WalletAddress, apikey.Passphrase)

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
		if ownedNfts, ok := result["ownedNfts"].([]interface{}); ok && len(ownedNfts) == 0 {
			// 没有NFT数据，修改point_notification_settings为1
			if err := db.Model(&user).Update("point_notification_settings", 1).Error; err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

type QueryPoint struct {
}

func (q QueryPoint) Exec(arg interface{}) error {
	//定时查询点位是否过期 如果过期 就需要把标题（点位币种名称 类型做多还是多空） 内容（点位币种名称 类型做多还是多空 入场价格 止盈价格数组 止损价格 有效时长 到期时间） 需要写入到information_notification表标题 内容 type为4。 并且向information_notification_user表 插入信息 用户id是开启了点位状态设置且点位权限是所有的用户NotificationSettings有含3的用户且point_notification_settings为2
	//查询information_trading_point 进行中status为1  已通知isNotified为2
	var points []adminModels.InformationTradingPoint

	// 获取数据库连接
	dbMap := sdk.Runtime.GetDb()
	if len(dbMap) == 0 {
		return errors.New("no database configured for jobs")
	}
	var db *gorm.DB
	for _, v := range dbMap {
		db = v
		break
	}

	if db == nil {
		return errors.New("database is nil")
	}

	if err := db.Where("status = 1", 1, 2).Find(&points).Error; err != nil {
		return err
	}
	for _, point := range points {
		if point.ValidUntil.Before(time.Now()) {
			coin := adminModels.InformationCoin{}
			db.First(&coin, point.CoinId)
			var enPointName string
			if point.Type == 1 {
				point.Name = coin.Symbol + "_做多" + ":" + "失效"
				enPointName = coin.Symbol + "_long" + ":" + "expired"
			} else {
				point.Name = coin.Symbol + "_做空" + ":" + "失效"
				enPointName = coin.Symbol + "_short" + ":" + "expired"
			}
			//title 币名称+做多还是做空
			//content 币种名称 类型做多还是多空 入场价格 止盈价格数组 止损价格 有效时长 到期时间
			notification := adminModels.InformationNotification{
				Title:     point.Name,
				EnTitle:   enPointName,
				Content:   point.Name + "，入场价格是" + point.EntryPrice + "，止盈价格是" + point.TakeProfits + "，止损价格是" + point.StopLoss + "，有效时长是" + strconv.FormatInt(point.ValidDuration, 10) + "小时，截止时间是" + point.ValidUntil.Format("2006-01-02 15:04:05"),
				EnContent: enPointName + "，entry price is " + point.EntryPrice + "，take profits is " + point.TakeProfits + "，stop loss is " + point.StopLoss + "，valid duration is " + strconv.FormatInt(point.ValidDuration, 10) + " hours，deadline is " + point.ValidUntil.Format("2006-01-02 15:04:05"),
				Type:      4, // 点位通知类型
			}
			err := db.Create(&notification).Error
			if err != nil {
				return err
			}
			// 查询开启了点位通知的用户
			var Users []adminModels.InformationUser
			err2 := db.Where("JSON_CONTAINS(notification_settings, JSON_QUOTE(?))  AND point_notification_settings = 2", "3").Find(&Users).Error
			if err2 != nil {
				return err2
			}

			if len(Users) == 0 {
				// 没有用户需要通知，直接更新点位状态
				if err := db.Model(&point).Update("status", 2).Error; err != nil {
					return err
				}
				continue
			}

			// 批量创建通知关联
			var notificationUsers []adminModels.InformationNotificationUser
			for _, user := range Users {
				notificationUsers = append(notificationUsers, adminModels.InformationNotificationUser{
					UserId:         int64(user.Id),
					NotificationId: int64(notification.Id),
					IsRead:         2,
				})
			}

			// 批量插入通知关联
			err3 := db.CreateInBatches(notificationUsers, 100).Error
			if err3 != nil {
				return err3
			}
			//修改状态
			if err := db.Model(&point).Update("status", 2).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
