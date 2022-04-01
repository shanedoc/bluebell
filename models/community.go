package models

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64  `json:"id" db:"community_id"`
	Name         string `json:"name" db:"community_name"`
	Introduction string `json:"introduction,omitempty" db:"introduction"`
	CreateTime   string `json:"create_time,omitempty" db:"create_time"`
}
