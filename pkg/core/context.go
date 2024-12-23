package core

import (
	"encoding/json"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Context struct {
	*gin.Context
	Logger *zap.Logger
}

type Method interface {
	C_BODY(obj any)
	C_JSON(code int, obj any)
}

func (c Context) C_BODY(obj any) {
	j, _ := json.Marshal(obj)
	c.Request.Header.Add("c_body", string(j))
}

func (c Context) C_JSON(code int, obj pkg.ResponseResult) {

	loggerResult := pkg.LoggerResult{Status: obj.Status, Message: obj.Message}
	j, _ := json.Marshal(loggerResult)
	c.Request.Header.Add("c_result", string(j))

	c.JSON(code, obj)
}

type HandlerFunc func(context *Context)

func Handler(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Context = c
		context.Logger = zap.L()

		handler(context)
	}
}
