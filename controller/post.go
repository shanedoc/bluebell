package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
	//获取分页信息
	offset, limit := getPageInfo(c)
	//获取数据
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("Logic.GetPostList() error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// V2.0 GetPostListHandler 根据不同的排序规则展示列表数据
// 1、获取参数
// 2、redis查询id列表
// 3、根据id查询mysql详细数据
func GetPostListHandler2(c *gin.Context) {
	//获取分页信息
	//get请求使用bindquery方法
	//初始化结构体时默认初始化参数
	p := &models.ParamsPostList{
		Offset: models.DefaultListOffset,
		Limit:  models.DefaultListLimit,
		Order:  models.OrderByTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("Logic.GetPostList() error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// CreatePostHandler
func CreatePostHandler(c *gin.Context) {
	//获取参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("c.ShouldBindJSON(p) invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取当前用户id
	userId, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userId
	//创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, CodeSuccess)
}

// GetPostDeatilHandler 帖子详情
func GetPostDeatilHandler(c *gin.Context) {
	//获取id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDeatilHandler invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取帖子详情
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("GetPostDetailById(pid) error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}
