package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FormGQLRespond struct {
	Data struct {
		Forms []*FormAPI `json:"iblock"`
	} `json:"data"`
}

type FormAPI struct {
	Id             int                `json:"iblock_id"`
	FormCode       string             `json:"iblock_code"`
	FormType       string             `json:"iblock_type"`
	Sort           int                `json:"sort"`
	Name           string             `json:"name"`
	Active         bool               `json:"active"`
	Properties     []*FormPropertyAPI `json:"properties"`
	LastUpdateDate int64              `json:"last_update_date"`
	ListFields     []*FormFieldAPI    `json:"list_fields"`
}

type FormPropertyAPI struct {
	Id           int                     `json:"id"`
	Name         string                  `json:"name"`
	Code         *string                 `json:"code"`
	Type         string                  `json:"type"`
	UserType     string                  `json:"user_type"`
	IblockId     int                     `json:"iblock_id"`
	IsRequired   bool                    `json:"is_required"`
	DefaultValue string                  `json:"default_value"`
	Sort         int                     `json:"sort"`
	Active       bool                    `json:"active"`
	Multiple     bool                    `json:"multiple"`
	XmlId        string                  `json:"xml_id"`
	Values       []*FormPropertyValueAPI `json:"values"`
}

/*func (u FormProperty) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u FormProperty) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}*/

type FormPropertyValueAPI struct {
	Id         int    `json:"id"`
	Sort       int    `json:"sort"`
	Name       string `json:"name"`
	XmlId      string `json:"xml_id"`
	IsDefault  bool   `json:"is_default"`
	PropertyId int    `json:"property_id"`
}

type FormFieldAPI struct {
	IblockId int    `json:"iblock_id"`
	FieldId  string `json:"field_id"`
	Sort     int    `json:"sort"`
	Name     string `json:"name"`
	Settings string `json:"settings"`
}

/*func (u FormField) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u FormField) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}*/

type FormDB struct {
	Id             int                   `gorm:"column:id;primaryKey" json:"iblock_id"`
	FormCode       string                `gorm:"column:form_code" json:"iblock_code"`
	FormType       string                `gorm:"column:form_type" json:"iblock_type"`
	Sort           int                   `gorm:"column:sort" json:"sort"`
	Name           string                `gorm:"column:name" json:"name"`
	Active         bool                  `gorm:"column:active" json:"active"`
	Properties     *ListOfFormPropertyDB `gorm:"column:properties" json:"properties"`
	LastUpdateDate int64                 `gorm:"column:last_update_date" json:"last_update_date"`
	ListFields     *ListOfFormFieldDB    `gorm:"column:list_fields" json:"list_fields"`
}

func (FormDB) TableName() string {
	return "public.bitrix_forms"
}

type ListOfFormFieldDB []*FormFieldDB

func (u ListOfFormFieldDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u ListOfFormFieldDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

type FormFieldDB struct {
	IblockId int    `json:"iblock_id"`
	FieldId  string `json:"field_id"`
	Sort     int    `json:"sort"`
	Name     string `json:"name"`
	Settings string `json:"settings"`
}

type ListOfFormPropertyDB []*FormPropertyDB

type FormPropertyDB struct {
	Id           int                        `json:"id"`
	Name         string                     `json:"name"`
	Code         *string                    `json:"code"`
	Type         string                     `json:"type"`
	UserType     string                     `json:"user_type"`
	IblockId     int                        `json:"iblock_id"`
	IsRequired   bool                       `json:"is_required"`
	DefaultValue string                     `json:"default_value"`
	Sort         int                        `json:"sort"`
	Active       bool                       `json:"active"`
	Multiple     bool                       `json:"multiple"`
	XmlId        string                     `json:"xml_id"`
	Values       *ListOfFormPropertyValueDB `json:"values"`
}

func (u ListOfFormPropertyDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u ListOfFormPropertyDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

type ListOfFormPropertyValueDB []*FormPropertyValueDB

type FormPropertyValueDB struct {
	Id         int    `json:"id"`
	Sort       int    `json:"sort"`
	Name       string `json:"name"`
	XmlId      string `json:"xml_id"`
	IsDefault  bool   `json:"is_default"`
	PropertyId int    `json:"property_id"`
}

func (u ListOfFormPropertyValueDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u ListOfFormPropertyValueDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}
