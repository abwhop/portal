package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FileAPI struct {
	Id           int     `json:"id"`
	Link         string  `json:"link"`
	FileName     string  `json:"file_name"`
	OriginalName string  `json:"original_name"`
	ContentType  string  `json:"content_type"`
	Size         int     `json:"size"`
	Height       float32 `json:"height"`
	Width        float32 `json:"width"`
}

type FileDB struct {
	Id           int     `json:"id"`
	Link         string  `json:"link"`
	FileName     string  `json:"fileName"`
	OriginalName string  `json:"originalName"`
	ContentType  string  `json:"contentType"`
	Size         int     `json:"size"`
	Height       float32 `json:"height"`
	Width        float32 `json:"width"`
}

type ListOfFileDB []*FileDB

func (u ListOfFileDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u ListOfFileDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

func (u FileDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u FileDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}
