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

func (srv *Service) LoadAllForms() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	loadStart := time.Now()

	var respondModel *models.FormGQLRespond

	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.FormQuery), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}

	loadedItemCount := len(respondModel.Data.Forms)
	fmt.Printf("Data loaded: count: %d time: %s\n", loadedItemCount, time.Since(loadStart))

	if loadedItemCount == 0 {
		return 0, nil
	}

	startSaveTime := time.Now()

	itemsDB, err := convertForms(respondModel.Data.Forms)
	if err != nil {
		return 0, err
	}

	if err := repository.NewRepository(srv.config.Database).SetForms(itemsDB); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.Forms), nil
}

func convertForms(itemsAPI []*models.FormAPI) ([]*models.FormDB, error) {
	var formsDB []*models.FormDB
	for _, item := range itemsAPI {
		itemDB, err := convertForm(item)
		if err != nil {
			continue
		}
		formsDB = append(formsDB, itemDB)
	}
	return formsDB, nil
}
func convertForm(itemAPI *models.FormAPI) (*models.FormDB, error) {
	listFields, err := convertFields(itemAPI.ListFields)
	if err != nil {
		listFields = nil
	}

	listProperties, err := convertProperties(itemAPI.Properties)
	if err != nil {
		listFields = nil
	}
	return &models.FormDB{
		Id:             itemAPI.Id,
		FormCode:       itemAPI.FormCode,
		FormType:       itemAPI.FormType,
		Sort:           itemAPI.Sort,
		Name:           itemAPI.Name,
		Active:         itemAPI.Active,
		Properties:     listProperties,
		LastUpdateDate: time.Now().UnixMilli(),
		ListFields:     listFields,
	}, nil
}

func convertFields(fieldsAPI []*models.FormFieldAPI) (*models.ListOfFormFieldDB, error) {
	var list *models.ListOfFormFieldDB
	marshal, err := json.Marshal(fieldsAPI)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}

	return list, err
}

func convertProperties(fieldsAPI []*models.FormPropertyAPI) (*models.ListOfFormPropertyDB, error) {
	var list *models.ListOfFormPropertyDB
	marshal, err := json.Marshal(fieldsAPI)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}

	return list, err
}
