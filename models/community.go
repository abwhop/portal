package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommunityGQLRespond struct {
	Data struct {
		Workgroups []*CommunityAPI `json:"workgroups"`
	} `json:"data"`
}

type CommunityAPI struct {
	Id               int                  `json:"id"`
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	Active           bool                 `json:"active"`
	DateCreate       int64                `json:"date_create"`
	Img              string               `json:"img"`
	Closed           bool                 `json:"closed"`
	Visible          bool                 `json:"visible"`
	Opened           bool                 `json:"opened"`
	Project          bool                 `json:"project"`
	UserIsMember     bool                 `json:"user_is_member"`
	GroupIsFavorite  bool                 `json:"group_is_favorite"`
	ProjectDateStart int64                `json:"project_date_start"`
	ProjectDateEnd   int64                `json:"project_date_end"`
	Subject          *CommunitySubjectAPI `json:"subject"`
	Type             *CommunityTypeAPI    `json:"type"`
	Author           *UserAPI             `json:"author"`
	Files            []*FileAPI           `json:"files"`
	Members          []*UserAPI           `json:"members"`
	Moderators       []*UserAPI           `json:"moderators"`
	Favorites        []*UserAPI           `json:"favorites"`
	Features         struct {
		Blog struct {
			WritePost string `json:"write_post"`
		} `json:"blog"`
	} `json:"features"`
}

type CommunitySubjectAPI struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
}

type CommunityTypeAPI struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CommunitySubjectDB struct {
	Id   int    `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Sort int    `gorm:"column:sort" json:"sort"`
}

func (CommunitySubjectDB) TableName() string {
	return "public.bitrix_communities_subjects"
}
func (u CommunitySubjectDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if user, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(user)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u CommunitySubjectDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

type CommunityTypeDB struct {
	Code        string `gorm:"column:code;primaryKey" json:"code"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

func (CommunityTypeDB) TableName() string {
	return "public.bitrix_communities_types"
}

func (u CommunityTypeDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if user, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(user)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u CommunityTypeDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

type CommunityDB struct {
	Id          int                 `gorm:"column:id;primaryKey" json:"id"`
	Name        string              `gorm:"column:name" json:"name"`
	Description string              `gorm:"column:description" json:"description"`
	Active      bool                `gorm:"column:active" json:"active"`
	DateCreated int64               `gorm:"column:date_created" json:"dateCreated"`
	Img         string              `gorm:"column:img" json:"img"`
	Closed      bool                `gorm:"column:closed" json:"closed"`
	Visible     bool                `gorm:"column:visible" json:"visible"`
	Opened      bool                `gorm:"column:opened" json:"opened"`
	Project     bool                `gorm:"column:project" json:"project"`
	Author      *UserDB             `gorm:"column:author" json:"author"`
	Type        *CommunityTypeDB    `gorm:"column:type" json:"type"`
	Subject     *CommunitySubjectDB `gorm:"column:subject" json:"subject"`
	Files       *ListOfFileDB       `gorm:"column:files" json:"files"`
	//CountMembers    int                 `gorm:"column:name" json:"countMembers"`
	Members         *ListOfUsersDB `gorm:"column:members" json:"members"`
	Moderators      *ListOfUsersDB `gorm:"column:moderators" json:"moderators"`
	Favorites       *ListOfUsersDB `gorm:"column:favorites" json:"favorites"`
	UserIsMember    bool           `gorm:"column:user_is_member" json:"userIsMember"`
	GroupIsFavorite bool           `gorm:"column:group_is_favorite" json:"groupIsFavorite"`
}

func (CommunityDB) TableName() string {
	return "public.bitrix_communities"
}
