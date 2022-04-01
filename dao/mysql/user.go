package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
)

//执行mysql的CURD

const sercet = "huanghe"

// InsertUser 插入用户数据
func InsertUser(user *models.User) (err error) {
	//密码加密
	user.Password = encrptPassword(user.Password)
	//入库
	sqlstr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlstr, user.UserId, user.Username, user.Password)
	return
}

func encrptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(sercet))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// CheckUserExist 校验重名用户
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlstr := `select user_id,username, password from user where username=?`
	err = db.Get(user, sqlstr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	//校验密码
	password := encrptPassword(oPassword)
	fmt.Println(password, user.Password)
	if password != user.Password {
		return ErrorUserPassword
	}
	return
}

// GetUserById 根据id查询用户信息
func GetUserById(id int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id=?`
	err = db.Get(user, sqlStr, id)
	return
}
