package models

//定义请求参数结构体

const (
	DefaultListOffset = 1
	DefaultListLimit  = 10
	OrderByTime       = "time"
	OrderByScore      = "score"
)

//login-请求参数
type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Repassword string `json:"repassword" binding:"required,eqfield=Password"`
}

//login-请求参数
type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 2.0 获取帖子列表参数
type ParamsPostList struct {
	Offset int64  `json:"offset" form:"offset"`
	Limit  int64  `json:"limit" form:"limit"`
	Order  string `json:"order" form:"order"`
}

// vote-投票
type ParamsVoteData struct {
	PostId    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成1 反对-1 取消0
}
