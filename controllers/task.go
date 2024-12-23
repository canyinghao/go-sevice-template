package controllers

import (
	"net/http"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/pkg/core"
	"github.com/canyinghao/go-sevice-template/services"
)

func Task(c *core.Context) {

	task := services.GetTaskOne()

	c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 0, Data: task.Name, Message: "ok"})
}

// application/json时需要用json，其它的get或post都可以用form
type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `form:"password" json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	Email    string `form:"email" json:"email" binding:"email" msg:"邮箱地址格式不正确"`
}

func PostTask(c *core.Context) {

	objA := UserInfo{}
	err := c.ShouldBind(&objA)

	c.C_BODY(objA)

	if err != nil {
		c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 1, Message: GetValidMsg(err, &objA)})
		return
	}
	c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 0, Data: objA.Username, Message: "ok"})

}
