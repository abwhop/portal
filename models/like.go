package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LikesAPI struct {
	Count int        `json:"count"`
	Users []*UserAPI `json:"users"`
}

type LikesDB struct {
	Count int            `json:"count"`
	Users *ListOfUsersDB `json:"users"`
}

func (u LikesDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	if item, err := json.Marshal(u); err == nil {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{string(item)},
		}
	} else {
		return clause.Expr{
			SQL:  "?",
			Vars: []interface{}{nil},
		}
	}
}

func (u LikesDB) Value() (driver.Value, error) {
	var item []byte
	var err error
	if item, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(item), nil
}
