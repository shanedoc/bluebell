package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (list []*models.Community, err error) {
	sqlstr := `select  community_id,community_name from community`
	if err := db.Select(&list, sqlstr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("GetCommunityList() no rows")
			err = nil
		}
	}
	return
}

func GetCommunityDetailById(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlstr := `select community_id,community_name,introduction,create_time from community where community_id=?`
	if err := db.Get(community, sqlstr, id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("GetCommunityList() no rows")
			err = ErrorInvalidID
		}
	}
	return community, err
}
