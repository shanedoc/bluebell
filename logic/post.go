package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

// GetPostById 获取帖子详情
// 帖子详情优化：字段展示
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	//查询数据组合页面所需的数据字段
	post, err := mysql.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailById error",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}
	//查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById error",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}
	//根据社区id查询社区详情
	community, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById error",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	//生成post_id
	p.ID = snowflake.GenID()
	//写入数据库
	if err := mysql.CreatePost(p); err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
}

// GetPostList 获取帖子列表
func GetPostList(offset, limit int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(offset, limit)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		//查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById error",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		//根据社区id查询社区详情
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById error",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		detail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, detail)
	}
	return
}

// GetPostList2 2.0获取帖子列表，根据不同排序规则展示数据
func GetPostList2(p *models.ParamsPostList) (data []*models.ApiPostDetail, err error) {
	//func GetPostList2(p *models.ParamsPostList) (err error) {
	//1、redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	fmt.Println(ids)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return row=0")
		return
	}
	zap.L().Debug("redis.GetPostIDsInOrder", zap.Any("ids", ids))
	//2、根据id去mysql获取详细数据
	posts, err := mysql.GetPostListByIds(ids)
	if err != nil {
		return
	}
	//zap.L().Debug("redis.GetPostList2", zap.Any("posts", posts))
	//查询每个帖子的投票数
	//redis.GetPostVoteData(ids)
	// 3、将帖子的作者及分区信息查询出来填充到帖子中
	for _, post := range posts {
		//查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById error",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		//根据社区id查询社区详情
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById error",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		detail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, detail)
	}
	return
}
