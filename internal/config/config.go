package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var App = new(AppConfig)

// AppConfig 包含应用程序的所有配置项，包括 MySQL 和 Redis 的配置项
type AppConfig struct {
	Server   ServerConfig `yaml:"server"`
	Database MysqlConfig  `yaml:"database"`
	Redis    RedisConfig  `yaml:"redis"`
	Jwt      JwtConfig    `yaml:"jwt"`
}

// InitConfig 从指定路径读取配置文件，并解析为 AppConfig 结构体
func InitConfig(configFile string) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, App)
	if err != nil {
		panic(err)
	}
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
	url := fmt.Sprintf("%s:%d",
		c.Redis.Host,
		c.Redis.Port,
	)

	return url
}
