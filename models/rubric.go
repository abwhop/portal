package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RubricAPI struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Code string `json:"code"`
}

type RubricDB struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Code string `json:"code"`
}

func (u RubricDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u RubricDB) Value() (driver.Value, error) {
	var item []byte
	var err error
	if item, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(item), nil
}

type UserRubrics struct {
	Appointment []*RubricAppointment `json:"appointment"`
	News        []*RubricNews        `json:"news"`
	Vacancy     []*RubricVacancy     `json:"vacancy"`
}

func (u UserRubrics) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u UserRubrics) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

type RubricNews struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RubricAppointment struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RubricVacancy struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
