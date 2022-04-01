package redis

import "bluebell/models"

// GetPostIDsInOrder
func GetPostIDsInOrder(p *models.ParamsPostList) ([]string, error) {
	//reids获取id
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderByScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//确定查询索引起始点
	start := (p.Offset - 1) * p.Limit
	end := start + p.Limit - 1
	//zrevrange查询:分数倒序查询
	return rdb.ZRevRange(key, start, end).Result()
}
