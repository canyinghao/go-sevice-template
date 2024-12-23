package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/canyinghao/go-sevice-template/pkg"
	"github.com/canyinghao/go-sevice-template/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {

	router := gin.Default()
	router.GET("/", core.Handler(Root))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	result := pkg.ResponseResult{}
	json.Unmarshal([]byte(w.Body.String()), &result)
	assert.Equal(t, 0, result.Status)
}
