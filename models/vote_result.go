package models

type ViteResultGQLRespond struct {
	Data struct {
		VoteResults []*VoteResultAPI `json:"voteResults"`
	} `json:"data"`
}

type VoteResultAPI struct {
	Id     int      `json:"id"`
	Date   int      `json:"date"`
	User   *UserAPI `json:"user"`
	VoteId int      `json:"vote_id"`
}

type VoteResultDB struct {
	Id     int     `gorm:"column:id;primaryKey" json:"id"`
	Date   int     `gorm:"column:date" json:"date"`
	User   *UserDB `gorm:"column:user" json:"user"`
	VoteId int     `gorm:"column:vote_id" json:"vote_id"`
}

func (VoteResultDB) TableName() string {
	return "public.bitrix_vote_results"
}
