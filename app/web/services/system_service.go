package services

import (
	"blog/app/repositories"
	"errors"
)

type SystemService struct {
}

// 获取系统配置项
func (service SystemService) GetSystemConfigs() map[string]interface{} {

	// 从数据库取出配置并格式化
	repo := repositories.NewSystemConfigRepositories()
	list, err := repo.GetConfigKv()

	configs := make(map[string]interface{},len(list))
	if err != nil {
		return nil
	}
	for _, v := range list {
		configs[v.Key] = v.Value
	}
	return configs
}

// 获取系统配置
func (service SystemService) GetSystemConfig(key string) (interface{}, error) {
	configs := service.GetSystemConfigs()
	value, ok := configs[key]
	if ok == false {
		return nil, errors.New("配置[" + key + "]不存在")
	}
	return value, nil
}
