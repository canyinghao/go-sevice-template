package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/pkg/core"
	"github.com/canyinghao/go-sevice-template/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {

	buf, _ := os.ReadFile("../config/development.json")

	// 解析配置文件
	var config pkg.Config
	json.Unmarshal(buf, &config)

	services.InitServices(&config)

	router := gin.Default()
	router.GET("/task", core.Handler(Task))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/task", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	result := pkg.ResponseResult{}
	json.Unmarshal([]byte(w.Body.String()), &result)

	fmt.Println(result)
	assert.Equal(t, 0, result.Status)

}
