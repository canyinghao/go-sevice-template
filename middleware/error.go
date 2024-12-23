package middleware

import (
	"github.com/gin-gonic/gin"

	"runtime"

	"github.com/canyinghao/go-sevice-template/pkg"
)

func Error() gin.HandlerFunc {

	return func(c *gin.Context) {

		defer func() {
			if p := recover(); p != nil {
				buf := make([]byte, 2048)
				runtime.Stack(buf, true)

				c.JSON(500, pkg.ResponseResult{Status: 500, Message: "server error"})

			}
		}()

		c.Next()

	}
}
