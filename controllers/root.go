package controllers

import (
	"net/http"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/pkg/core"
)

// Root
// @Summary 测试
// @Schemes
// @Description 测试接口是否正常
// @Tags 测试
// @Accept json
// @Produce json
// @Success 200 {object} pkg.ResponseResult{data=string} "status为0表示成功，其它失败"
// @Router / [get]
func Root(c *core.Context) {

	c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 0, Data: "hello world", Message: "ok"})
}
