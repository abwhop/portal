package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ViteGQLRespond struct {
	Data struct {
		Votes []*VoteAPI `json:"vote"`
	} `json:"data"`
}

type VoteAPI struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Author      *UserAPI       `json:"author"`
	Active      bool           `json:"active"`
	DateFrom    int            `json:"date_from"`
	DateTo      int            `json:"date_to"`
	Questions   []*QuestionAPI `json:"questions"`
	Img         *FileAPI       `json:"img"`
	DateChange  int            `json:"date_change"`
	Url         string         `json:"url"`
	VoteGroup   *VoteGroupAPI  `json:"vote_group"`
	Views       int            `json:"views"`
	Counter     int            `json:"counter"`
}

type QuestionAPI struct {
	Id           int          `json:"id"`
	Sort         int          `json:"sort"`
	Question     string       `json:"question"`
	DateChange   int          `json:"date_change"`
	Active       bool         `json:"active"`
	Counter      int          `json:"counter"`
	Diagram      bool         `json:"diagram"`
	Required     bool         `json:"required"`
	DiagramType  string       `json:"diagram_type"`
	QuestionType string       `json:"question_type"`
	Answers      []*AnswerAPI `json:"answers"`
}

type AnswerAPI struct {
	Id         int    `json:"id"`
	Sort       int    `json:"sort"`
	Message    string `json:"message"`
	FieldType  string `json:"field_type"`
	DateChange int    `json:"date_change"`
	Active     bool   `json:"active"`
	Counter    *int   `json:"counter"`
}

type VoteGroupAPI struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Sort       int      `json:"sort"`
	Active     bool     `json:"active"`
	Hidden     bool     `json:"hidden"`
	DateChange int      `json:"date_change"`
	Title      string   `json:"title"`
	VoteSingle bool     `json:"vote_single"`
	UseCaptcha bool     `json:"use_captcha"`
	SiteId     []string `json:"site_id"`
}

type VoteDB struct {
	Id          int               `gorm:"column:id;primaryKey" json:"id"`
	Title       string            `gorm:"column:title" json:"title"`
	Description string            `gorm:"column:description" json:"description"`
	Author      *UserDB           `gorm:"column:author" json:"author"`
	Active      bool              `gorm:"column:active" json:"active"`
	DateFrom    int               `gorm:"column:date_from" json:"date_from"`
	DateTo      int               `gorm:"column:date_to" json:"date_to"`
	Questions   *ListOfQuestionDB `gorm:"column:questions" json:"questions"`
	Img         *FileDB           `gorm:"column:img" json:"img"`
	DateChange  int               `gorm:"column:date_change" json:"date_change"`
	Url         string            `gorm:"column:url" json:"url"`
	VoteGroup   *VoteGroupDB      `gorm:"column:vote_group" json:"vote_group"`
	Views       int               `gorm:"column:views" json:"views"`
	Counter     int               `gorm:"column:counter" json:"counter"`
}

func (VoteDB) TableName() string {
	return "public.bitrix_votes"
}

type QuestionDB struct {
	Id           int             `json:"id"`
	Sort         int             `json:"sort"`
	Question     string          `json:"question"`
	DateChange   int             `json:"date_change"`
	Active       bool            `json:"active"`
	Counter      int             `json:"counter"`
	Diagram      bool            `json:"diagram"`
	Required     bool            `json:"required"`
	DiagramType  string          `json:"diagram_type"`
	QuestionType string          `json:"question_type"`
	Answers      *ListOfAnswerDB `json:"answers"`
}

type ListOfQuestionDB []*QuestionDB

func (u ListOfQuestionDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u ListOfQuestionDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

type AnswerDB struct {
	Id         int    `json:"id"`
	Sort       int    `json:"sort"`
	Message    string `json:"message"`
	FieldType  string `json:"field_type"`
	DateChange int    `json:"date_change"`
	Active     bool   `json:"active"`
	Counter    *int   `json:"counter"`
}

type ListOfAnswerDB []*AnswerDB

func (u ListOfAnswerDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u ListOfAnswerDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

type VoteGroupDB struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Sort       int      `json:"sort"`
	Active     bool     `json:"active"`
	Hidden     bool     `json:"hidden"`
	DateChange int      `json:"date_change"`
	Title      string   `json:"title"`
	VoteSingle bool     `json:"vote_single"`
	UseCaptcha bool     `json:"use_captcha"`
	SiteId     []string `json:"site_id"`
}

func (u VoteGroupDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u VoteGroupDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}
