package redishelper

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
)

var (
	// DefaultRedis 默认Redis帮助类实例
	DefaultRedis = &RedisHelper{}
)

// RedisConfig Redis配置结构
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// CacheConfig 缓存配置结构
type CacheConfig struct {
	Redis  RedisConfig `yaml:"redis"`
	Memory interface{} `yaml:"memory"`
}

// Settings 配置文件结构
type Settings struct {
	Cache CacheConfig `yaml:"cache"`
}

// RedisHelper Redis帮助类
type RedisHelper struct {
	client *redis.Client
}

// GetClient 获取Redis客户端
func (r *RedisHelper) GetClient() *redis.Client {
	if r.client == nil {
		if err := r.initClient(); err != nil {
			log.Fatalf("初始化Redis客户端失败: %v", err)
		}
	}
	return r.client
}

// initClient 初始化Redis客户端
func (r *RedisHelper) initClient() error {
	// 读取配置文件
	configPath := "config/settings.yml"
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 将配置转换为Settings结构
	settings := &Settings{}
	err = yaml.Unmarshal(configData, settings)
	if err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 从配置中获取Redis配置
	redisConfig := settings.Cache.Redis

	// 创建Redis客户端
	r.client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,     // Redis服务器地址
		Password: redisConfig.Password, // Redis密码
		DB:       redisConfig.DB,       // Redis数据库编号
		PoolSize: 10,                   // 连接池大小
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = r.client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis连接失败: %v", err)
	}

	log.Println("Redis连接成功")
	return nil
}

// Close 关闭Redis连接
func (r *RedisHelper) Close() error {
	if r.client != nil {
		return r.client.Close()
	}
	return nil
}

func (r *RedisHelper) DeleteString(key any) {
	if r.client == nil {
		return
	}
	ctx := context.Background()
	r.client.Del(ctx, fmt.Sprintf("%v", key))
}

// Set 设置字符串值
func (r *RedisHelper) Set(key string, value interface{}, expiration time.Duration) error {
	if r.client == nil {
		return nil
	}
	ctx := context.Background()
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get 获取字符串值
func (r *RedisHelper) Get(key string) (string, error) {
	if r.client == nil {
		return "", nil
	}
	ctx := context.Background()
	return r.client.Get(ctx, key).Result()
}

// Delete 删除键
func (r *RedisHelper) Delete(keys ...string) error {
	if r.client == nil {
		return nil
	}
	ctx := context.Background()
	return r.client.Del(ctx, keys...).Err()
}

// HSet 设置哈希字段
func (r *RedisHelper) HSet(key, field string, value interface{}) error {
	if r.client == nil {
		return nil
	}
	ctx := context.Background()
	return r.client.HSet(ctx, key, field, value).Err()
}

// HGet 获取哈希字段值
func (r *RedisHelper) HGet(key, field string) (string, error) {
	if r.client == nil {
		return "", nil
	}
	ctx := context.Background()
	return r.client.HGet(ctx, key, field).Result()
}

// HGetAll 获取哈希所有字段
func (r *RedisHelper) HGetAll(key string) (map[string]string, error) {
	if r.client == nil {
		return make(map[string]string), nil
	}
	ctx := context.Background()
	return r.client.HGetAll(ctx, key).Result()
}

// HDel 删除哈希字段
func (r *RedisHelper) HDel(key string, fields ...string) error {
	if r.client == nil {
		return nil
	}
	ctx := context.Background()
	return r.client.HDel(ctx, key, fields...).Err()
}

// SAdd 添加到集合
func (r *RedisHelper) SAdd(key string, members ...interface{}) error {
	if r.client == nil {
		return nil
	}
	ctx := context.Background()
	return r.client.SAdd(ctx, key, members...).Err()
}

// SRem 从集合删除
func (r *RedisHelper) SRem(key string, members ...interface{}) error {
	if r.client == nil {
		return nil
	}
	ctx := context.Background()
	return r.client.SRem(ctx, key, members...).Err()
}

// SIsMember 检查是否在集合中
func (r *RedisHelper) SIsMember(key string, member interface{}) (bool, error) {
	if r.client == nil {
		return false, nil
	}
	ctx := context.Background()
	return r.client.SIsMember(ctx, key, member).Result()
}

// SMembers 获取集合所有成员
func (r *RedisHelper) SMembers(key string) ([]string, error) {
	if r.client == nil {
		return []string{}, nil
	}
	ctx := context.Background()
	return r.client.SMembers(ctx, key).Result()
}

// RandomKey 随机获取一个key
func (r *RedisHelper) RandomKey(pattern string) (string, error) {
	if r.client == nil {
		return "", nil
	}
	ctx := context.Background()
	// 获取匹配pattern的所有key
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return "", err
	}
	if len(keys) == 0 {
		return "", nil
	}
	// 随机选择一个key
	randomIndex := time.Now().UnixNano() % int64(len(keys))
	return keys[randomIndex], nil
}
