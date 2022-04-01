package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	_ "bluebell/pkg/jwt"
	"bluebell/settings"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strings"
)

//路由注册

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	//绑定自定义日志
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	v1 := r.Group("/api/v1")
	// 注册
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.GET("/posts", controller.GetPostListHandler)
		//根据时间或分数对帖子列表进行排序
		v1.GET("/posts2", controller.GetPostListHandler2)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDeatilHandler)

		v1.POST("/vote", controller.PostVoteHandler)
	}

	//未知路由-404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
