package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 处理请求
var ErrorUserNotLogin = errors.New("用户未登录")

const CtxtUserIdKey = "user_id"

// getCurrentUser 获取当前登录用户id
func getCurrentUser(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(CtxtUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// 列表分页
func getPageInfo(c *gin.Context) (int64, int64) {
	//分页参数
	offsetstr := c.Query("offset")
	limitstr := c.Query("limit")
	var (
		offset int64
		limit  int64
		err    error
	)
	offset, err = strconv.ParseInt(offsetstr, 20, 64)
	if err != nil {
		offset = 1
	}
	limit, err = strconv.ParseInt(limitstr, 20, 64)
	if err != nil {
		limit = 10
	}
	return offset, limit
}
