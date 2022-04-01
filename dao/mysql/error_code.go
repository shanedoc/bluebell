package mysql

import "errors"

//定义常量
var (
	ErrorUserExist    = errors.New("用户已存在")
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorUserPassword = errors.New("用户名或密码错误")
	ErrorInvalidID    = errors.New("无效的id")
)
