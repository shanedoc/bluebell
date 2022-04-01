package main

import (
	"bluebell/controller"
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/settings"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	//go-web框架搭建
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}
	//TODO:: 1、加载配置文件
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Printf("init settings err:%v\n", err)
		return
	}
	fmt.Println(settings.Conf)
	fmt.Println(settings.Conf.LogConfig == nil)
	//TODO:: 2、初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//TODO:: 3、初始化mysql,将全局变量当成参数传入init函数
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql err:%v\n", err)
		return
	}
	defer mysql.Close()
	//TODO:: 4、、初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis err:%v\n", err)
		return
	}
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	//TODO:: validator汉化包初始化
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init trans err:%v\n", err)
		return
	}
	//TODO:: 5、、注册路由
	r := router.SetUp(settings.Conf.Mode)
	if err := r.Run(); err != nil {
		fmt.Printf("run server err:%v\n", err)
		return
	}
}
