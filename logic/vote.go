package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 投票功能：简易版投票算法
// 用户投一票+432分 86400/200 200张可以让帖子在首页呆一天
// 赞成1 反对-1 取消0

// PostVote 投票功能
func PostVote(userID int64, p *models.ParamsVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostId),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostId, float64(p.Direction))
}
