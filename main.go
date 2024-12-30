package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	static "github.com/canyinghao/gin-static"
	"github.com/canyinghao/go-sevice-template/cron"
	"github.com/canyinghao/go-sevice-template/middleware"
	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/routers"
	"github.com/canyinghao/go-sevice-template/rpc"
	"github.com/canyinghao/go-sevice-template/services"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/canyinghao/go-sevice-template/docs"
)

var APPENV = "development"
var PORT = ":8080"
var RPC_PORT = ""
var isCron = false

var (
	g errgroup.Group
)

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
	// rpc的端口，如果为空，将不会启动rpc
	RPC_PORT = config.Rpc

	isCron = config.Cron
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

	if APPENV != "production" {
		// api文档
		docs.SwaggerInfo.Title = "Swagger Example API"
		docs.SwaggerInfo.Description = "This is a sample server Petstore server."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost" + PORT
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	// 日志中间件
	r.Use(middleware.GinLogger(zap.L()), middleware.GinRecovery(zap.L(), true))

	// 支持gzip
	r.Use(static.GzipMiddleware())

	// 配置路由
	routers.InitRouters(r)

	// 配置是否开启cron
	if isCron {
		cronManager := cron.NewCronJobManager()
		cronManager.Start()
		defer cronManager.Stop()
	}

	go func() {
		time.Sleep(time.Second)
		fmt.Println("访问地址：http://localhost" + PORT)
	}()

	s := &http.Server{
		Addr:           PORT,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if len(RPC_PORT) > 0 {
		rpc.Register()

		g.Go(func() error {
			return s.ListenAndServe()
		})

		g.Go(func() error {
			return http.ListenAndServe(RPC_PORT, nil)
		})

		if err := g.Wait(); err != nil {
			log.Fatal(err)
		}

	} else {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}

}
