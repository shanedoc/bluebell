package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//使用结构体
type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

//使用全局变量使用配置参数，避免在各个文件中引入viper包
// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

//TODO:: viper读取配置文件的几种形式
func Init(config string) (err error) {
	//配合远程配置中心使用，告诉viper当前数据格式
	//viper.SetConfigType("yaml")
	//读取viper文件的形式1
	//viper.SetConfigName("config") //配置文件名称
	//viper.AddConfigPath(".")      // 还可以在工作目录中查找配置
	//读取配置文件形式2
	viper.SetConfigFile(config)

	//查找并读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("viper readingconfig err:%v\n", err)
		return err
	}
	//把读取到的配置信息反序列化到conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper Unmarshal failed err:%v\n", err)
	}
	//监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		//TODO::通知
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed err:%v\n", err)
		}
	})
	return
}
