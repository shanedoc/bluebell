package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

/*func CreatePost(postId, communityId int64) error {
	pipline := rdb.TxPipeline()
	//帖子时间
	pipline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	//帖子分数
	pipline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	//更新：把帖子id加到社区的set
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityId)))
	pipline.SAdd(ckey, postId)
	_, err := pipline.Exec()
	return err
}*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePreVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

/*
情况：
direction=1
	之前没投票，现在赞成	更新分数和投票记录	差值绝对值1 	+432
	之前取消，现在赞成	更新分数和投票记录	差值绝对值2	+432*2
direction=0
	之前赞成，现在取消	更新分数和投票记录	差值绝对值1	-432
	之前反对，现在取消	更新分数和投票记录	差值绝对值1	+432
direction=-1
	之前没投票，现在反对	更新分数和投票记录	差值绝对值1	-432
	之前取消，现在反对	更新分数和投票记录	差值绝对值2	-432*2
投票限制：
每个帖子发表之日起，一周内允许投票
到期后将redis中数据保存到mysql并删除key
*/

// VoteForPost 投票
func VoteForPost(userID, postID string, value float64) error {
	//1、判断投票限制
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//23需要放在一个pipline事务中操作
	//2、更新帖子分数
	//查询投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipline := rdb.TxPipeline()
	pipline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePreVote, postID)

	//3、记录用户投票数据
	if value == 0 {
		pipline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID).Result()
	} else {
		pipline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, //赞成or反对
			Member: userID,
		})
	}
	_, err := pipline.Exec()
	return err
}

// CreatePost 创建帖子
func CreatePost(postId, communityId int64) error {
	//事务
	pipline := rdb.TxPipeline()
	//帖子时间
	pipline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	//帖子分数
	pipline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	_, err := pipline.Exec()
	return err
}

// GetPostVoteData 根据id获取投票数据
/*func GetPostVoteData(ids []string) (vote *[]models.ParamsVoteData, err error) {

}*/
