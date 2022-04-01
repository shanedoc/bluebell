package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	_ "errors"
)

//业务逻辑

func SignUp(p *models.ParamsSignUp) (err error) {
	//判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//生成uid
	userId := snowflake.GenID()
	//构造user
	u := &models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	//写入数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamsLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传指针
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	//生成jwt-token
	return jwt.GenToken(user.UserId, user.Username)
}
