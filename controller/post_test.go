package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

//test-post

func TestCreatePostHandler(t *testing.T) {
	r := gin.Default()
	url := "/api/v1/post"
	//调用创建帖子接口
	r.POST(url, CreatePostHandler)
	body := `{
		"community_id":1,
		"title":"test",
		"content":"just a test"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	//判断响应内容是否按照预期返回需要登录的错误
	//1：响应字段中是否包含需要登录
	//assert.Contains(t, w.Body.String(), "需要登录")
	//2：将响应的内容反序列化到ResponseData 判断字段与预期是否一致
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
