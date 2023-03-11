package config

// ServerConfig 包含 系统 相关的配置项
type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
