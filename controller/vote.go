package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type VoteData struct {
	PostId    int64 `json:"post_id,string"`
	Direction int   `json:"direction,string"` //赞成1 反对-1
}

// PostVoteHandler 投票
func PostVoteHandler(c *gin.Context) {
	//校验参数
	p := new(models.ParamsVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		//翻译并去掉结构体的标示
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	uid, err := getCurrentUser(c)
	if err != nil {
		zap.L().Error("logic getCurrentUser() error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if err := logic.PostVote(uid, p); err != nil {
		zap.L().Error("logic PostVote() error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
