package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// AppConfig 包含应用程序的所有配置项，包括 MySQL 和 Redis 的配置项
type AppConfig struct {
	Server   ServerConfig `yaml:"server"`
	Database MysqlConfig  `yaml:"database"`
	Redis    RedisConfig  `yaml:"redis"`
}

// LoadConfig 从指定路径读取配置文件，并解析为 AppConfig 结构体
func LoadConfig(configFile string) (*AppConfig, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	config := new(AppConfig)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetMysqlDSN 返回 MySQL 数据源名称
func (c *AppConfig) GetMysqlDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)

	return dsn
}

// GetRedisURL 返回 Redis 连接 URL
func (c *AppConfig) GetRedisURL() string {
	url := fmt.Sprintf("redis://%s:%s/%d",
		c.Redis.Host,
		c.Redis.Password,
		c.Redis.DB,
	)

	return url
}
