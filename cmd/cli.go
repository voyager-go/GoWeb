package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"github.com/voyager-go/GoWeb/internal/config"
	"github.com/voyager-go/GoWeb/internal/router"
	"github.com/voyager-go/GoWeb/pkg/logging"
	"github.com/voyager-go/GoWeb/pkg/mysql"
	"github.com/voyager-go/GoWeb/pkg/orm"
	"github.com/voyager-go/GoWeb/pkg/redis"
	"net/http"
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
		// 初始化Redis
		redis.InitPool(config.App.GetRedisURL(), "")
		// 初始化Gorm
		orm.InitPool(config.App.GetMysqlDSN())
		return nil
	},
	Action: func(*cli.Context) error {
		var (
			srv = gin.New()
		)
		// 404 处理
		srv.NoRoute(func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "The page you requested is not found",
			})
		})
		// 初始化路由
		router.InitRouter(srv)
		// 启动项目
		return srv.Run(":" + strconv.Itoa(config.App.Server.Port))
	},
	After: func(*cli.Context) error {
		defer redis.Conn.Close()
		defer mysql.Conn.Close()
		return nil
	},
}
