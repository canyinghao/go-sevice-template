package controllers

import (
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/pkg/core"
)

func DoRpc(c *core.Context) {
	client, err := rpc.DialHTTP("tcp", "localhost:9909")
	if err != nil {
		c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 2, Message: "rpc链接失败"})
		return
	}
	defer client.Close()

	a, e1 := strconv.Atoi(c.Query("a"))
	b, e2 := strconv.Atoi(c.Query("b"))
	if e1 != nil || e2 != nil {
		c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 1, Message: "参数错误"})
		return
	}
	args := Args{a, b}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 0, Data: reply, Message: "ok"})

}

type Args struct {
	A, B int
}
