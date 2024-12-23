package controllers

import (
	"net/http"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/pkg/core"
)

func Root(c *core.Context) {

	c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 0, Data: "hello world", Message: "ok"})
}
