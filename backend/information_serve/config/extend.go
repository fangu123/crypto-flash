package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	AMap  AMap  // 这里配置对应配置文件的结构即可
	Email Email // 邮件配置
	OAuth OAuth // OAuth配置
}

type AMap struct {
	Key string
}

// Email 邮件配置
type Email struct {
	Host     string `yaml:"host"`     // SMTP服务器地址
	Port     int    `yaml:"port"`     // SMTP端口
	User     string `yaml:"user"`     // 邮箱账号
	Password string `yaml:"password"` // 邮箱密码或授权码
	From     string `yaml:"from"`     // 发件人地址
}

// OAuth OAuth相关配置
type OAuth struct {
	Google    OAuthProvider `yaml:"google"`    // Google OAuth配置
	Microsoft OAuthProvider `yaml:"microsoft"` // Microsoft OAuth配置
}

// OAuthProvider OAuth提供商配置
type OAuthProvider struct {
	ClientID     string `yaml:"clientID"`     // 客户端ID
	ClientSecret string `yaml:"clientSecret"` // 客户端密钥
	RedirectURL  string `yaml:"redirectURL"`  // 重定向URL
}
