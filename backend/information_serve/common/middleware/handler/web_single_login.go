package handler

import (
	"context"
	"fmt"
	"go-admin/common/redishelper"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
)

// CheckWebSingleLogin 检查Web端单点登录的中间件
// 如果用户的token与Redis中存储的token不一致，则返回401错误
func CheckWebSingleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 提取JWT claims
		claims := jwt.ExtractClaims(c)
		if claims == nil {
			response.Error(c, 401, nil, "not log in")
			c.Abort()
			return
		}

		// 获取用户ID和token ID
		userId, ok := claims["identity"]
		if !ok {
			response.Error(c, 401, nil, "invalid token")
			c.Abort()
			return
		}

		tokenId, ok := claims["tokenId"]
		if !ok {
			// 如果没有tokenId，说明是旧版本的token，允许继续使用
			c.Next()
			return
		}

		// 从Redis中获取当前有效的token ID
		ctx := context.Background()
		tokenKey := fmt.Sprintf("web_token:%v", userId)
		redisClient := redishelper.DefaultRedis.GetClient()

		cachedTokenId, err := redisClient.Get(ctx, tokenKey).Result()
		if err != nil {
			response.Error(c, 401, nil, "Please log in again")
			// Redis连接错误，记录日志但不阻止请求
			//log := api.GetRequestLogger(c)
			//log.Errorf("get token from redis error, %s", err.Error())
			c.Next()
			return
		}

		// 检查token是否一致
		if cachedTokenId == "" {
			// Redis中没有token信息，可能是token已过期，允许继续使用
			response.Error(c, 401, nil, "Please log in again")
			c.Next()
			return
		}

		if cachedTokenId != fmt.Sprintf("%v", tokenId) {
			// token不一致，说明用户已在别处登录，挤退当前会话
			response.Error(c, 401, nil, "Your account has already been logged in elsewhere. Please log in again")
			c.Abort()
			return
		}

		// token一致，允许继续访问
		c.Next()
	}
}
