package service

import (
	"errors"
	"sync"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"
)

var (
	sysConfigCache     = make(map[string]models.SysConfig)
	sysConfigCacheLock sync.RWMutex
)

type SysConfig struct {
	service.Service
}

// InitCache 初始化配置到内存缓存
func (e *SysConfig) InitCache() error {
	var list []models.SysConfig
	err := e.Orm.Find(&list).Error
	if err != nil {
		e.Log.Errorf("Service InitSysConfigCache error:%s", err)
		return err
	}

	sysConfigCacheLock.Lock()
	defer sysConfigCacheLock.Unlock()

	sysConfigCache = make(map[string]models.SysConfig, len(list))
	for _, data := range list {
		sysConfigCache[data.ConfigKey] = data
	}
	return nil
}

// cacheSet 更新缓存中的单条记录
func cacheSet(item models.SysConfig) {
	sysConfigCacheLock.Lock()
	defer sysConfigCacheLock.Unlock()
	sysConfigCache[item.ConfigKey] = item
}

// cacheDelete 从缓存中删除记录
func cacheDelete(configKey string) {
	sysConfigCacheLock.Lock()
	defer sysConfigCacheLock.Unlock()
	delete(sysConfigCache, configKey)
}

// cacheGet 从缓存中读取记录
func cacheGet(configKey string) (models.SysConfig, bool) {
	sysConfigCacheLock.RLock()
	defer sysConfigCacheLock.RUnlock()
	val, ok := sysConfigCache[configKey]
	return val, ok
}

// cacheGetAll 从缓存中读取所有记录
func cacheGetAll() []models.SysConfig {
	sysConfigCacheLock.RLock()
	defer sysConfigCacheLock.RUnlock()

	list := make([]models.SysConfig, 0, len(sysConfigCache))
	for _, v := range sysConfigCache {
		list = append(list, v)
	}
	return list
}

// GetPage 获取SysConfig列表
func (e *SysConfig) GetPage(c *dto.SysConfigGetPageReq, list *[]models.SysConfig, count *int64) error {
	err := e.Orm.
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	return nil
}

// Get 获取SysConfig对象
func (e *SysConfig) Get(d *dto.SysConfigGetReq, model *models.SysConfig) error {
	err := e.Orm.First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("Service GetSysConfig error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysConfig对象
func (e *SysConfig) Insert(c *dto.SysConfigControl) error {
	var err error
	var data models.SysConfig
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("Service InsertSysConfig error:%s", err)
		return err
	}
	cacheSet(data)
	return nil
}

// Update 修改SysConfig对象
func (e *SysConfig) Update(c *dto.SysConfigControl) error {
	var err error
	var model = models.SysConfig{}
	e.Orm.First(&model, c.GetId())
	c.Generate(&model)
	db := e.Orm.Save(&model)
	err = db.Error
	if err != nil {
		e.Log.Errorf("Service UpdateSysConfig error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")

	}
	cacheSet(model)
	return nil
}

// SetSysConfig 修改SysConfig对象
func (e *SysConfig) SetSysConfig(c *[]dto.GetSetSysConfigReq) error {
	var err error
	for _, req := range *c {
		var model = models.SysConfig{}
		e.Orm.Where("config_key = ?", req.ConfigKey).First(&model)
		if model.Id != 0 {
			req.Generate(&model)
			db := e.Orm.Save(&model)
			err = db.Error
			if err != nil {
				e.Log.Errorf("Service SetSysConfig error:%s", err)
				return err
			}
			if db.RowsAffected == 0 {
				return errors.New("无权更新该数据")
			}
			cacheSet(model)
		}
	}
	return nil
}

func (e *SysConfig) GetForSet(c *[]dto.GetSetSysConfigReq) error {
	var err error
	var data models.SysConfig

	err = e.Orm.Model(&data).
		Find(c).Error
	if err != nil {
		e.Log.Errorf("Service GetSysConfigPage error:%s", err)
		return err
	}

	/*for index, item := range *c {
		if item.ConfigKey == "sys_user_2fa" && item.ConfigValue != "" {
			bytes, err := base64.StdEncoding.DecodeString(item.ConfigValue)

			if err != nil {
				logger.Errorf("解析base64失败")
			}

			decBytes, err1 := rsa.DecryptWithPrivateKey(bytes, crypto.PrivateStore.Get())

			if err1 != nil {
				logger.Errorf("2FA密钥解密失败")
			}
			(*c)[index].ConfigValue = string(decBytes)
		}
	}*/
	return nil
}

func (e *SysConfig) UpdateForSet(c *[]dto.GetSetSysConfigReq) error {
	m := *c
	for _, req := range m {
		var data models.SysConfig
		if err := e.Orm.Where("config_key = ?", req.ConfigKey).
			First(&data).Error; err != nil {
			e.Log.Errorf("Service GetSysConfigPage error:%s", err)
			return err
		}

		/*if req.ConfigKey == "sys_user_2fa" {
			encryptedValue, err := rsa.EncryptWithPublicKey([]byte(req.ConfigValue), crypto.PublicStore.Get())

			if err != nil {
				e.Log.Errorf("Service SetSysConfig error:%s", err)
				continue
			}
			req.ConfigValue = base64.StdEncoding.EncodeToString(encryptedValue)
		}*/

		if data.ConfigValue != req.ConfigValue {
			data.ConfigValue = req.ConfigValue

			if err := e.Orm.Save(&data).Error; err != nil {
				e.Log.Errorf("Service GetSysConfigPage error:%s", err)
				return err
			}
			cacheSet(data)
		}
	}

	return nil
}

// Remove 删除SysConfig
func (e *SysConfig) Remove(d *dto.SysConfigDeleteReq) error {
	var err error
	var data models.SysConfig

	// 先查出要删除的记录，用于清理缓存
	var delList []models.SysConfig
	if err = e.Orm.Where("id IN ?", d.Ids).Find(&delList).Error; err != nil {
		e.Log.Errorf("Service RemoveSysConfig find error:%s", err)
		return err
	}

	db := e.Orm.Delete(&data, d.Ids)
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Service RemoveSysConfig error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}

	// 清理缓存
	for _, item := range delList {
		cacheDelete(item.ConfigKey)
	}
	return nil
}

// GetWithKey 根据Key获取SysConfig
func (e *SysConfig) GetWithKey(c *dto.SysConfigByKeyReq, resp *dto.GetSysConfigByKEYForServiceResp) error {
	// 优先从缓存读取
	if item, ok := cacheGet(c.ConfigKey); ok {
		resp.ConfigKey = item.ConfigKey
		resp.ConfigValue = item.ConfigValue
		return nil
	}

	// 缓存未命中，从数据库读取
	var data models.SysConfig
	err := e.Orm.Table(data.TableName()).Where("config_key = ?", c.ConfigKey).First(resp).Error
	if err != nil {
		e.Log.Errorf("At Service GetSysConfigByKEY Error:%s", err)
		return err
	}

	return nil
}

func (e *SysConfig) GetWithKeyList(c *dto.SysConfigGetToSysAppReq, list *[]models.SysConfig) error {
	// 优先从缓存读取
	allCache := cacheGetAll()
	if len(allCache) > 0 {
		// 按搜索条件过滤
		for _, item := range allCache {
			if c.IsFrontend != "" && item.IsFrontend != c.IsFrontend {
				continue
			}
			*list = append(*list, item)
		}
	} else {
		// 缓存为空，从数据库读取
		var err error
		err = e.Orm.
			Scopes(
				cDto.MakeCondition(c.GetNeedSearch()),
			).
			Find(list).Error
		if err != nil {
			e.Log.Errorf("Service GetSysConfigByKey error:%s", err)
			return err
		}
	}

	/*for index, item := range *list {
		if item.ConfigKey == "sys_user_2fa" && item.ConfigValue != "" {
			bytes, err := base64.StdEncoding.DecodeString(item.ConfigValue)

			if err != nil {
				logger.Errorf("解析base64失败")
			}

			decBytes, err1 := rsa.DecryptWithPrivateKey(bytes, crypto.PrivateStore.Get())

			if err1 != nil {
				logger.Errorf("2FA密钥解密失败")
			}
			(*list)[index].ConfigValue = string(decBytes)
		}
	}*/
	return nil
}

func (e *SysConfig) GetSecret() (string, error) {
	// 优先从缓存读取
	/*if item, ok := cacheGet("sys_user_2fa"); ok && item.ConfigValue != "" {
		bytes, err := base64.StdEncoding.DecodeString(item.ConfigValue)
		if err != nil {
			logger.Errorf("解析base64失败")
		}

		decBytes, err1 := rsa.DecryptWithPrivateKey(bytes, crypto.PrivateStore.Get())
		if err1 != nil {
			logger.Errorf("2FA密钥解密失败")
		}
		return string(decBytes), nil
	}*/

	// 缓存未命中，走原有逻辑
	data := dto.GetSysConfigByKEYForServiceResp{}
	err := e.GetWithKey(&dto.SysConfigByKeyReq{ConfigKey: "sys_user_2fa"}, &data)

	if err != nil {
		return "", err
	}

	/*if data.ConfigValue != "" {
		bytes, err := base64.StdEncoding.DecodeString(data.ConfigValue)

		if err != nil {
			logger.Errorf("解析base64失败")
		}

		decBytes, err1 := rsa.DecryptWithPrivateKey(bytes, crypto.PrivateStore.Get())

		if err1 != nil {
			logger.Errorf("2FA密钥解密失败")
		}
		return string(decBytes), nil
	}*/

	return "", nil
}
