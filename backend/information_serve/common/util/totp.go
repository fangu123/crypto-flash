package util

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"net/url"
	"strings"
	"time"
)

// TOTPConfig TOTP配置
type TOTPConfig struct {
	Secret     string
	Issuer     string
	Account    string
	Digits     int
	Period     int
	WindowSize int
}

// GenerateSecret 生成随机密钥
func GenerateSecret(length int) (string, error) {
	if length <= 0 {
		length = 32
	}

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base32.StdEncoding.EncodeToString(bytes), nil
}

// GenerateTOTP 生成TOTP验证码
func GenerateTOTP(secret string, digits int, period int) (string, error) {
	if digits <= 0 {
		digits = 6
	}
	if period <= 0 {
		period = 30
	}

	// 解码Base32密钥
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// 获取当前时间步长
	counter := uint64(time.Now().Unix()) / uint64(period)

	// 生成HMAC
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, counter)

	h := hmac.New(sha1.New, key)
	h.Write(buf)
	hash := h.Sum(nil)

	// 动态截取
	offset := hash[len(hash)-1] & 0x0F
	truncated := binary.BigEndian.Uint32(hash[offset:offset+4]) & 0x7FFFFFFF

	// 生成验证码
	code := truncated % uint32(math.Pow10(digits))

	return fmt.Sprintf("%0*d", digits, code), nil
}

// ValidateTOTP 验证TOTP验证码
func ValidateTOTP(secret string, code string, digits int, period int, windowSize int) bool {
	if digits <= 0 {
		digits = 6
	}
	if period <= 0 {
		period = 30
	}
	if windowSize <= 0 {
		windowSize = 1
	}

	// 解码Base32密钥
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return false
	}

	// 获取当前时间步长
	counter := uint64(time.Now().Unix()) / uint64(period)

	// 检查时间窗口内的所有可能值
	for i := -windowSize; i <= windowSize; i++ {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, counter+uint64(i))

		h := hmac.New(sha1.New, key)
		h.Write(buf)
		hash := h.Sum(nil)

		offset := hash[len(hash)-1] & 0x0F
		truncated := binary.BigEndian.Uint32(hash[offset:offset+4]) & 0x7FFFFFFF

		expectedCode := fmt.Sprintf("%0*d", digits, truncated%uint32(math.Pow10(digits)))
		if expectedCode == code {
			return true
		}
	}

	return false
}

// GenerateQRCodeURL 生成TOTP二维码URL
func GenerateQRCodeURL(config *TOTPConfig) string {
	if config == nil {
		return ""
	}

	if config.Digits <= 0 {
		config.Digits = 6
	}
	if config.Period <= 0 {
		config.Period = 30
	}

	issuer := url.QueryEscape(config.Issuer)
	account := url.QueryEscape(config.Account)
	secret := strings.ReplaceAll(config.Secret, " ", "")

	otpAuthURL := fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s&digits=%d&period=%d",
		issuer, account, secret, issuer, config.Digits, config.Period)

	// 返回一个使用Google Charts API生成的二维码图片URL
	qrURL := fmt.Sprintf("https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=%s", url.QueryEscape(otpAuthURL))

	return qrURL
}
