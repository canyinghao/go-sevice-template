package main

import (
	"embed"
	"encoding/json"
	"log"
	"os"
	"time"

	static "github.com/canyinghao/gin-static"
	"github.com/canyinghao/go-sevice-template/middleware"
	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/routers"
	"github.com/canyinghao/go-sevice-template/services"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

var APPENV = "development"
var PORT = ":8080"

func init() {

	env := os.Getenv("APP_ENV")

	if env != "" {
		APPENV = env
	}

	// 将config打包到可执行文件里
	staticFiles, _ := static.EmbedFolder(configFS, "config")
	// 读取配置文件
	fl, _ := staticFiles.Open(APPENV + ".json")
	buf, _ := static.ReadFile(fl)
	fl.Close()
	// 解析配置文件
	var config pkg.Config
	json.Unmarshal(buf, &config)

	PORT = config.Port
	// 日志初始化
	middleware.InitLog(&config)
	// 初始化数据库
	services.InitServices(&config)

}

//go:embed config
var configFS embed.FS

func main() {

	if APPENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	// 日志中间件
	r.Use(middleware.GinLogger(zap.L()), middleware.GinRecovery(zap.L(), true))

	// 支持gzip
	r.Use(static.GzipMiddleware())

	// 配置路由
	routers.InitRouters(r)

	go func() {
		time.Sleep(time.Second)
		log.Println("访问地址：http://localhost" + PORT)
	}()

	if err := r.Run(PORT); err != nil {
		log.Fatal(err)
	}
}
