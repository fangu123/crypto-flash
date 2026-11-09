package router

import (
	app "go-admin/app/other/apis/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerWebRouter)
}

func registerWebRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	// 添加跨域中间件
	v1.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization", "Accept-Language", "accept-Language"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	profile := app.Profile{}
	//当输入邮箱的时候就查询2FA是否为true 如果为true就显示2FA验证码 如果为false 就不显示2FA验证码
	v1.GET("/profile/query2fa", profile.Query2FA)
	//当输入邮箱的时候就查询2FA是否为true 如果为true就显示2FA验证码 如果为false 就不显示2FA验证码
	//发送邮箱验证码
	v1.POST("/profile/sendEmailCode", profile.SendEmailCode)
	//忘记密码
	v1.POST("/profile/forgetPassword", profile.ForgetPassword)

	// 首页相关接口
	home := app.HomePage{}
	//首页市场概览  查询api接口 https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=bitcoin,ethereum,solana&sparkline=true&price_change_percentage=1h  1分钟请求一次 查询出了数据就存缓存 如果没查出数据就从缓存取
	v1.GET("/app/home/market-overview", home.GetMarketOverview)
	//首页多空比 查询bybit的api接口 /v5/market/account-ratio?category=linear&symbol=BTCUSDT&period=1h&limit=2&startTime=1696089600000&endTime=1696262400000 分1h 4h 24h 7d查询条件 查询出了数据就存缓存 如果没查出数据就从缓存取
	v1.GET("/app/home/long-short-ratio", home.GetLongShortRatio)
	//首页重要财经日历 查询odaily的api接口  https://api.odaily.news/api/v1/newsflash/24h-hot?page=1&size=10   查询出了数据就存缓存 如果没查出数据就从缓存取
	v1.GET("/app/home/financial-calendar", home.GetFinancialCalendar)
	//首页新闻快讯 查询blockBeats的api接口 https://api-pro.theblockbeats.info/v1/newsflash/important   查询出了数据就存缓存 如果没查出数据就从缓存取
	v1.GET("/app/home/news-flash", home.GetNewsFlash)
	//首页最新交易点位 查询表information_trading_point的最新一条数据
	v1.GET("/app/home/latest-trading-point", home.GetLatestTradingPoint)
	//首页搜索功能 搜索市场概览用币种、多空比用币种、点位用币种、财经和快讯
	v1.GET("/app/home/search", home.Search)
	// 资讯相关接口
	news := app.News{}
	//资讯列表（通过搜索标题 内容直接查询information_news表和分类查询就是分类表information_content_tag关联information_news_tag_relation 查询information_news
	v1.GET("/app/news/list", news.GetNewsList)
	//获取标签
	v1.GET("/app/news/getTags", news.GetTags)
	points := app.Points{}
	//客户联系方式
	v1.GET("/app/profile/contact", profile.GetContact)
	//关于我们版本
	v1.GET("/app/profile/appVersion", profile.GetAppVersion)

	r := v1.Group("/app").Use(authMiddleware.MiddlewareFunc())
	{
		auth := app.Auth{}
		r.POST("/logout", auth.Logout)

		//资讯详情（分类 相关币种 相关点位） 增加一次浏览次数字段view_count+1
		r.GET("/news/detail/:id", news.GetNewsDetail)
		//资讯分享 增加一次分享次数字段share_count+1
		r.POST("/news/share/:id", news.ShareNews)
		//资讯收藏和取消收藏 操作information_news_favorite表
		r.POST("/news/favorite/:id", news.FavoriteNews)
		r.POST("/news/unfavorite/:id", news.UnfavoriteNews)

		//点位列表（通过分类展示）(全部 做多 做空 自选收藏的点位) 查询表information_trading_point  自选查询收藏表information_trading_point_favorite
		r.GET("/trading-point/list", points.GetTradingPointList)
		//点位详情
		r.GET("/trading-point/detail/:id", points.GetTradingPointDetail)
		//点位分享
		r.POST("/trading-point/share/:id", points.ShareTradingPoint)
		//点位收藏
		r.POST("/trading-point/favorite/:id", points.FavoriteTradingPoint)
		r.POST("/trading-point/unfavorite/:id", points.UnfavoriteTradingPoint)
		//解锁全部点位 查询当前地址是否有相关的NFT  如果有就修改表information_user的字段pointNotificationSettings为2
		r.POST("/trading-point/unlock-all", points.UnlockAllTradingPoint)

		//消息通知列表
		r.GET("/profile/messageList", profile.GetMessageList)
		//标记消息为已读
		r.POST("/profile/markMessageRead/:id", profile.MarkMessageRead)
		//我的收藏
		r.GET("/profile/myCollection", profile.GetMyCollection)

		//修改密码
		r.POST("/profile/password", profile.UpdatePassword)
		//消息通知设置列表（数组：1资讯 2新点位 3点位状态 4系统）有值就是true 没有就是false
		r.GET("/profile/messageSettingList", profile.GetMessageSetting) //返回information_user的字段NotificationSettings数组
		//消息通知设置（数组：1资讯 2新点位 3点位状态 4系统）根据type1开启 2关闭修改表information_user的字段NotificationSettings数组
		r.POST("/profile/messageSettingUpdate", profile.UpdateMessageSetting)
		//获取2FA
		r.GET("/profile/get2fa", profile.Get2FA)
		//开启关闭2FA 验证2FA 开启2FA时就生成url 和密钥  输入邮箱验证码+谷歌验证码 （两步登录就会开启  two_factor_enabled为true）   关闭2FA   输入邮箱验证码+谷歌验证码（清空hash 都清空）
		r.POST("/profile/update2fa", profile.Update2FA)
		//关闭两步登录保护 需要输入验证码  two_factor_enabled设置为false  开启两步登录保护 需要输入验证码 two_factor_enabled设置为true
		r.POST("/profile/enabled2fa", profile.Enabled2FA)
		//绑定钱包地址
		r.POST("/profile/bindWallet", profile.BindWallet)
		//刷新token
		r.POST("/profile/refreshToken", profile.RefreshToken)
		//获取个人资料
		r.GET("/profile/personalInfo", profile.GetPersonalInfo)
		//修改个人资料
		r.POST("/profile/updatePersonalInfo", profile.UpdatePersonalInfo)
		comment := app.Comment{}
		//资讯评论列表、发布、删除、点赞及取消点赞
		r.GET("/comment/list", comment.GetCommentList)
		r.POST("/comment/publish", comment.PublishComment)
		r.DELETE("/comment/:id", comment.DeleteComment)
		r.POST("/comment/:id/like", comment.LikeComment)
		r.POST("/comment/:id/unlike", comment.UnlikeComment)
	}

}
