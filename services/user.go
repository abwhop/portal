package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/abwhop/portal_models/models"
	"github.com/abwhop/portal_sync/gql"
	"github.com/abwhop/portal_sync/query"
	"github.com/abwhop/portal_sync/repository"
	"time"
)

func (srv *Service) LoadUsers(limit int, page int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	loadStart := time.Now()
	var respondModel *models.UserGQLRespond
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.UsersQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Data loaded:", time.Since(loadStart))
	if len(respondModel.Data.Users) == 0 {
		return 0, nil
	}
	startSaveTime := time.Now()

	items, err := ConvertFullUsers(respondModel.Data.Users)
	if err != nil {
		return 0, err
	}

	if err := repository.NewRepository(srv.config.Database).SetUsers(items); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.Users), nil
}
func ConvertFullUsers(usersAPI []*models.UserFullAPI) ([]*models.UserFullDB, error) {
	var usersDB []*models.UserFullDB
	for _, userAPI := range usersAPI {
		userDB, err := ConvertFullUser(userAPI)
		if err != nil {
			continue
		}
		usersDB = append(usersDB, userDB)
	}
	return usersDB, nil
}
func ConvertFullUser(userAPI *models.UserFullAPI) (*models.UserFullDB, error) {
	if userAPI == nil {
		return nil, nil
	}
	gender := "men"
	if userAPI.Gender == "F" {
		gender = "women"
	}
	return &models.UserFullDB{
		Id:                 userAPI.Id,
		Login:              userAPI.Login,
		Active:             userAPI.Active,
		Name:               userAPI.Name,
		LastName:           userAPI.LastName,
		SecondName:         userAPI.SecondName,
		Email:              userAPI.Email,
		PersonalMobile:     userAPI.PersonalMobile,
		PersonalPhone:      userAPI.PersonalPhone,
		Gender:             gender,
		Photo:              userAPI.Photo,
		Company:            userAPI.Company,
		Department:         userAPI.Department,
		Position:           userAPI.Position,
		Birthday:           userAPI.Birthday,
		CompanyId:          userAPI.CompanyId,
		PersonalNumber:     userAPI.PersonalNumber,
		FullPersonalNumber: userAPI.FullPersonalNumber,
		CreateDate:         userAPI.CreateDate,
		UpdateDate:         userAPI.UpdateDate,
		StartWorkDate:      userAPI.StartWorkDate,
		FactoryId:          userAPI.FactoryId,
		Education:          userAPI.Education,
		AboutMe:            userAPI.AboutMe,
		HiddenFields:       userAPI.HiddenFields,
		ManufactoryName:    userAPI.ManufactoryName,
		DepartmentName:     userAPI.DepartmentNameSp,
		DepartmentNameSp:   userAPI.DepartmentName,
		DepartmentAddress:  userAPI.DepartmentAddress,
		ChiefId:            userAPI.ChiefId,
		LoginAd:            userAPI.LoginAd,
		Favorites:          userAPI.Favorites,
		Rights:             userAPI.Rights,
		WorkProfile:        userAPI.WorkProfile,
		LastActivityDate:   userAPI.LastActivityDate,
		LastLogin:          userAPI.LastLogin,
		Rubrics:            userAPI.Rubrics,
		UfSiteId:           userAPI.UfSiteId,
		BxDepartmentId:     userAPI.BxDepartmentId,
		CovidQrCode:        userAPI.CovidQrCode,
		CovidQrCodeDecoded: userAPI.CovidQrCodeDecoded,
		//CreatedAt:                 time.Now(),
		//UpdatedAt:                 time.Now(),
		//CovidQrCodeValidationData: userAPI.CovidQrCodeValidationData,
		BlackListType:    userAPI.BlackListType,
		BlackListMessage: userAPI.BlackListMessage,
	}, nil
}

func ConvertUser(userAPI *models.UserAPI) (*models.UserDB, error) {
	if userAPI == nil {
		return nil, nil
	}
	gender := "men"
	if userAPI.Gender == "F" {
		gender = "women"
	}
	return &models.UserDB{
		BitrixUserId:    userAPI.Id,
		FirstName:       userAPI.Name,
		LastName:        userAPI.LastName,
		MiddleName:      userAPI.SecondName,
		Email:           userAPI.Email,
		PositionName:    userAPI.Position,
		PersonnelNumber: userAPI.PersonalNumber,
		Active:          userAPI.Active == "Y",
		Gender:          gender,
		Photo:           userAPI.Photo,
		LoginAd:         userAPI.LoginAd,
	}, nil
}

func ConvertUsers(usersAPI []*models.UserAPI) (*models.ListOfUsersDB, error) {
	var usersDB []*models.UserDB
	var list *models.ListOfUsersDB
	for _, userAPI := range usersAPI {
		userDB, err := ConvertUser(userAPI)
		if err != nil {
			continue
		}
		usersDB = append(usersDB, userDB)
	}
	marshal, err := json.Marshal(usersDB)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}
	return list, nil
}
