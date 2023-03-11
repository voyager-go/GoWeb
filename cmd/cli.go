package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"github.com/voyager-go/GoWeb/internal/config"
	"github.com/voyager-go/GoWeb/pkg/logging"
	"path/filepath"
	"strconv"
)

var App = &cli.App{
	Name:     "main",
	Usage:    "start this project",
	Commands: []*cli.Command{},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config-file",
			Value:   "example.yaml",
			Usage:   "path to configuration file",
			EnvVars: []string{"CONFIG_FILE"},
		},
	},
	Before: func(c *cli.Context) error {
		// 初始化配置文件
		config.InitConfig(c.String("config-file"))
		// 初始化日志追踪
		logging.InitLogger(filepath.Join("storage", "logs"))
		// 初始化数据库连接
		//mysql.NewMysql()
		// 初始化Redis连接
		//redis.NewRedis()
		// 初始化验证器翻译
		//validator_trans.NewTrans()
		return nil
	},
	Action: func(*cli.Context) error {
		var (
			srv = gin.New()
		)
		// 404 处理
		srv.NoRoute(func(ctx *gin.Context) {
		})
		// 启动项目
		return srv.Run(":" + strconv.Itoa(config.App.Server.Port))
	},
}
