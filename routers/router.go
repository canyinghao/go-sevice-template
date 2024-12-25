package routers

import (
	"github.com/canyinghao/go-sevice-template/controllers"
	"github.com/canyinghao/go-sevice-template/middleware"
	"github.com/canyinghao/go-sevice-template/pkg/core"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {

	r.Use(middleware.Error())

	r.NoRoute(func(c *gin.Context) {

		c.String(404, "404 not found")

	})

	r.GET("/", core.Handler(controllers.Root))

	r.GET("/task", core.Handler(controllers.Task))

	r.Any("/postTask", core.Handler(controllers.PostTask))

	// rpc测试路由
	r.GET("/doRpc", core.Handler(controllers.DoRpc))

}
