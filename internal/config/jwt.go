package config

// JwtConfig 包含 Redis 相关的配置项
type JwtConfig struct {
	Secret  string `yaml:"secret"`
	Expired int64  `yaml:"expired"`
}
