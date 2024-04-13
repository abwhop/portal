package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DescriptionDB struct {
	Type     string `json:"type,omitempty"`
	Html     string `json:"html,omitempty"`
	Data     string `json:"data,omitempty"`
	Src      string `json:"src,omitempty"`
	Title    string `json:"title,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
	TypeFile string `json:"typeFile,omitempty"`
	Id       int    `json:"id,omitempty"`
}

type ListOfDescriptionDB []*DescriptionDB

func (u ListOfDescriptionDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u ListOfDescriptionDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}
