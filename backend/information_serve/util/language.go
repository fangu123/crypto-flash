package util

import (
	"github.com/gin-gonic/gin"
)

type contextKey string

const UserLanguageKey contextKey = "userLanguage"

func GetUserLanguage(c *gin.Context) string {
	if lang, exists := c.Get("userLanguage"); exists {
		return lang.(string)
	}
	return "zh-CN" // 默认中文
}

func SetUserLanguage(c *gin.Context, language string) {
	c.Set("userLanguage", language)
}
