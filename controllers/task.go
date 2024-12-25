package controllers

import (
	"net/http"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/pkg/core"
	"github.com/canyinghao/go-sevice-template/services"
)

// Task
// @Summary 查询一条数据
// @Schemes
// @Description 查询一条数据
// @Tags Task
// @Accept json
// @Produce json
// @Param        id   query      int  true  "Account ID"
// @Success 200 {object} pkg.ResponseResult{data=model.Task} "status为0表示成功，其它失败"
// @Router /task [get]
func Task(c *core.Context) {

	id := c.Query("id")

	task, err := services.GetTaskOne(id)
	if err != nil {
		c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 1, Message: err.Error()})
		return
	}

	c.C_JSON(http.StatusOK, pkg.ResponseResult{Status: 0, Data: &task, Message: "ok"})
}

// application/json时需要用json，其它的get或post都可以用form
type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=6" msg:"密码长度不能小于3大于6"`
	Email    string `form:"email" json:"email" binding:"omitempty,email" msg:"邮箱地址格式不正确"`
}

// PostTask
// @Summary 提交数据
// @Schemes
// @Description 提交数据
// @Tags Task
// @Accept json
// @Produce json
// @Param        json   body      UserInfo  true  "用户名"
// @Success 200 {object} pkg.ResponseResult{data=string} "status为0表示成功，其它失败"
// @Router /postTask [post]
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
