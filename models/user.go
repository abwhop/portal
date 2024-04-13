package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserAPI struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	SecondName     string `json:"second_name"`
	Email          string `json:"email"`
	Position       string `json:"position"`
	PersonalNumber int    `json:"personal_number"`
	Active         string `json:"active"`
	Gender         string `json:"gender"`
	Photo          string `json:"photo"`
	LoginAd        string `json:"login_ad"`
}
type UserGQLRespond struct {
	Data struct {
		Users []*UserFullAPI `json:"users"`
	} `json:"data"`
}
type UserFullAPI struct {
	Id                        int            `json:"id"`
	Login                     string         `json:"login"`
	Active                    string         `json:"active"`
	Name                      string         `json:"name"`
	LastName                  string         `json:"last_name"`
	SecondName                string         `json:"second_name"`
	Email                     string         `json:"email"`
	PersonalMobile            string         `json:"personal_mobile"`
	PersonalPhone             string         `json:"personal_phone"`
	Gender                    string         `json:"gender"`
	Photo                     string         `json:"photo"`
	Company                   string         `json:"company"`
	Department                string         `json:"department"`
	Position                  string         `json:"position"`
	Birthday                  int            `json:"birthday"`
	CompanyId                 int            `json:"company_id"`
	PersonalNumber            int            `json:"personal_number"`
	FullPersonalNumber        string         `json:"full_personal_number"`
	CreateDate                int            `json:"create_date"`
	UpdateDate                int64          `json:"update_date"`
	StartWorkDate             int            `json:"start_work_date"`
	FactoryId                 int            `json:"factory_id"`
	Education                 string         `json:"education"`
	AboutMe                   string         `json:"about_me"`
	HiddenFields              []string       `json:"hidden_fields"`
	ManufactoryName           string         `json:"manufactory_name"`
	DepartmentName            string         `json:"department_name"`
	DepartmentNameSp          string         `json:"department_name_sp"`
	DepartmentAddress         string         `json:"department_address"`
	ChiefId                   int            `json:"chief_id"`
	LoginAd                   string         `json:"login_ad"`
	Favorites                 *UserFavorites `json:"favorites"`
	Rights                    []string       `json:"rights"`
	WorkProfile               string         `json:"work_profile"`
	LastActivityDate          int64          `json:"last_activity_date"`
	LastLogin                 int64          `json:"last_login"`
	Rubrics                   *UserRubrics   `json:"rubrics"`
	UfSiteId                  string         `json:"uf_site_id"`
	BxDepartmentId            int            `json:"bx_department_id"`
	CovidQrCode               string         `json:"covid_qr_code"`
	CovidQrCodeDecoded        string         `json:"covid_qr_code_decoded"`
	CovidQrCodeValidationData string         `json:"covid_qr_code_validation_data"`
	BlackListType             string         `json:"black_list_type"`
	BlackListMessage          string         `json:"black_list_message"`
}

type UserDB struct {
	PersonnelNumber int    `json:"personnelNumber"`
	LastName        string `json:"lastName"`
	FirstName       string `json:"firstName"`
	MiddleName      string `json:"middleName"`
	PositionName    string `json:"positionName"`
	LoginAd         string `json:"loginAd"`
	BitrixUserId    int    `json:"bitrixUserId"`
	Active          bool   `json:"active"`
	Gender          string `json:"gender"`
	Email           string `json:"email"`
	Photo           string `json:"photo"`
}

type ListOfUsersDB []*UserDB

func (u ListOfUsersDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u ListOfUsersDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

func (u UserDB) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
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

func (u UserDB) Value() (driver.Value, error) {
	var user []byte
	var err error
	if user, err = json.Marshal(u); err != nil {
		return nil, err
	}
	return string(user), nil
}

type UserShort struct {
	PersonnelNumber int    `json:"personnel_number"`
	LastName        string `json:"last_name"`
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	CompanyName     string `json:"company_name"`
	DepartmentName  string `json:"department_name"`
	PositionName    string `json:"position_name"`
	LoginAd         string `json:"login_ad"`
	BitrixUserId    int    `json:"bitrix_user_id"`
	CompanyCode     int    `json:"company_code"`
	Photo           string `json:"photo"`
	Birthday        int    `json:"birthday"`
	Active          bool   `json:"active"`
	Gender          string `json:"gender"`
}

type UserFullDB struct {
	Id                 int            `gorm:"column:id;primaryKey" json:"id"`
	Login              string         `gorm:"column:login" json:"login"`
	Active             string         `gorm:"column:active" json:"active"`
	Name               string         `gorm:"column:name" json:"name"`
	LastName           string         `gorm:"column:last_name" json:"last_name"`
	SecondName         string         `gorm:"column:second_name" json:"second_name"`
	Email              string         `gorm:"column:email" json:"email"`
	PersonalMobile     string         `gorm:"column:personal_mobile" json:"personal_mobile"`
	PersonalPhone      string         `gorm:"column:personal_phone" json:"personal_phone"`
	Gender             string         `gorm:"column:gender" json:"gender"`
	Photo              string         `gorm:"column:photo" json:"photo"`
	Company            string         `gorm:"column:company" json:"company"`
	Department         string         `gorm:"column:department" json:"department"`
	Position           string         `gorm:"column:position" json:"position"`
	Birthday           int            `gorm:"column:birthday" json:"birthday"`
	CompanyId          int            `gorm:"column:company_id" json:"company_id"`
	PersonalNumber     int            `gorm:"column:personal_number" json:"personal_number"`
	FullPersonalNumber string         `gorm:"column:full_personal_number" json:"full_personal_number"`
	CreateDate         int            `gorm:"column:create_date" json:"create_date"`
	UpdateDate         int64          `gorm:"column:update_date" json:"update_date"`
	StartWorkDate      int            `gorm:"column:start_work_date" json:"start_work_date"`
	FactoryId          int            `gorm:"column:factory_id" json:"factory_id"`
	Education          string         `gorm:"column:education" json:"education"`
	AboutMe            string         `gorm:"column:about_me" json:"about_me"`
	HiddenFields       pq.StringArray `gorm:"column:hidden_fields;type:varchar[]" json:"hidden_fields"`
	ManufactoryName    string         `gorm:"column:manufactory_name" json:"manufactory_name"`
	DepartmentName     string         `gorm:"column:department_name" json:"department_name"`
	DepartmentNameSp   string         `gorm:"column:department_name_sp" json:"department_name_sp"`
	DepartmentAddress  string         `gorm:"column:department_address" json:"department_address"`
	ChiefId            int            `gorm:"column:chief_id" json:"chief_id"`
	LoginAd            string         `gorm:"column:login_ad" json:"login_ad"`
	Favorites          *UserFavorites `gorm:"column:favorites" json:"favorites"`
	Rights             pq.StringArray `gorm:"column:rights;type:varchar[]" json:"rights"`
	WorkProfile        string         `gorm:"column:work_profile" json:"work_profile"`
	LastActivityDate   int64          `gorm:"column:last_activity_date" json:"last_activity_date"`
	LastLogin          int64          `gorm:"column:last_login" json:"last_login"`
	Rubrics            *UserRubrics   `gorm:"column:id" json:"rubrics"`
	UfSiteId           string         `gorm:"column:uf_site_id" json:"uf_site_id"`
	BxDepartmentId     int            `gorm:"column:bx_department_id" json:"bx_department_id"`
	CovidQrCode        string         `gorm:"column:covid_qr_code" json:"covid_qr_code"`
	CovidQrCodeDecoded string         `gorm:"column:covid_qr_code_decoded" json:"covid_qr_code_decoded"`
	//CreatedAt                 time.Time `json:"createdAt"`
	//UpdatedAt                 time.Time `json:"updatedAt"`
	//CovidQrCodeValidationData string `json:"covid_qr_code_validation_data"`
	BlackListType    string `gorm:"column:black_list_type" json:"black_list_type"`
	BlackListMessage string `gorm:"column:black_list_message" json:"black_list_message"`
}

func (UserFullDB) TableName() string {
	return "public.bitrix_users"
}

/*
type User struct {
	Id                        int         `json:"id"`
	Login                     string      `json:"login"`
	Active                    string      `json:"active"`
	Name                      string      `json:"name"`
	LastName                  string      `json:"last_name"`
	SecondName                string      `json:"second_name"`
	Email                     string      `json:"email"`
	PersonalMobile            string      `json:"personal_mobile"`
	PersonalPhone             interface{} `json:"personal_phone"`
	Gender                    string      `json:"gender"`
	Photo                     interface{} `json:"photo"`
	Company                   string      `json:"company"`
	Department                interface{} `json:"department"`
	Position                  string      `json:"position"`
	Birthday                  int         `json:"birthday"`
	CompanyId                 int         `json:"company_id"`
	PersonalNumber            int         `json:"personal_number"`
	FullPersonalNumber        string      `json:"full_personal_number"`
	CreateDate                int         `json:"create_date"`
	UpdateDate                interface{} `json:"update_date"`
	StartWorkDate             int         `json:"start_work_date"`
	FactoryId                 int         `json:"factory_id"`
	Education                 interface{} `json:"education"`
	AboutMe                   interface{} `json:"about_me"`
	HiddenFields              []string    `json:"hidden_fields"`
	ManufactoryName           string      `json:"manufactory_name"`
	DepartmentName            string      `json:"department_name"`
	DepartmentNameSp          string      `json:"department_name_sp"`
	DepartmentAddress         string      `json:"department_address"`
	ChiefId                   int         `json:"chief_id"`
	LoginAd                   string      `json:"login_ad"`
	UserFavorites                 interface{} `json:"favorites"`
	Rights                    interface{} `json:"rights"`
	WorkProfile               string      `json:"work_profile"`
	LastActivityDate          interface{} `json:"last_activity_date"`
	LastLogin                 interface{} `json:"last_login"`
	UserRubrics                   interface{} `json:"rubrics"`
	UfSiteId                  interface{} `json:"uf_site_id"`
	BxDepartmentId            int         `json:"bx_department_id"`
	CovidQrCode               interface{} `json:"covid_qr_code"`
	CovidQrCodeDecoded        interface{} `json:"covid_qr_code_decoded"`
	CovidQrCodeValidationData interface{} `json:"covid_qr_code_validation_data"`
	BlackListType             interface{} `json:"black_list_type"`
	BlackListMessage          interface{} `json:"black_list_message"`
}*/
