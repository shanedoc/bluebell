package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//校验参数
	p := new(models.ParamsSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("signup params error", zap.Error(err))
		//类型断言：是否validator错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator类型错误
			/*c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})*/
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
	}
	//业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseErrorWithMsg(c, CodeUserExist, err)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err)
	}
	//返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//校验参数
	l := new(models.ParamsLogin)
	if err := c.ShouldBindJSON(l); err != nil {
		zap.L().Error("login params error", zap.Error(err))
		//类型断言：是否validator错误
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//非validator类型错误
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务处理
	token, err := logic.Login(l)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", l.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//返回响应
	ResponseSuccess(c, token)
}
