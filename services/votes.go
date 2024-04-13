package services

import (
	"context"
	"encoding/json"
	"fmt"
	"portal_sync/gql"
	"portal_sync/models"
	"portal_sync/query"
	"portal_sync/repository"
	"time"
)

func (srv *Service) LoadVotes(limit int, page int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	loafStart := time.Now()
	var respondModel models.ViteGQLRespond
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.VotesQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Data loaded:", time.Since(loafStart))
	startSaveTime := time.Now()

	itemsDB, err := convertVotes(respondModel.Data.Votes)
	if err != nil {
		return 0, err
	}

	if err := repository.NewRepository(srv.config.Database).SetVotes(itemsDB); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.Votes), nil
}

func convertVotes(itemsAPI []*models.VoteAPI) ([]*models.VoteDB, error) {
	var itemsDB []*models.VoteDB
	for _, item := range itemsAPI {
		itemDB, err := convertVote(item)
		if err != nil {
			continue
		}
		itemsDB = append(itemsDB, itemDB)
	}
	return itemsDB, nil
}
func convertVote(itemAPI *models.VoteAPI) (*models.VoteDB, error) {
	authorDB, err := ConvertUser(itemAPI.Author)
	if err != nil {
		authorDB = nil
	}
	questionsDB, err := convertQuestions(itemAPI.Questions)
	if err != nil {
		questionsDB = nil
	}
	voteGroup, err := convertVoteGroup(itemAPI.VoteGroup)
	if err != nil {
		voteGroup = nil
	}
	fileDB, err := ConvertFile(itemAPI.Img)
	if err != nil {
		fileDB = nil
	}
	return &models.VoteDB{
		Id:          itemAPI.Id,
		Title:       itemAPI.Title,
		Description: itemAPI.Description,
		Author:      authorDB,
		Active:      itemAPI.Active,
		DateFrom:    itemAPI.DateFrom,
		DateTo:      itemAPI.DateTo,
		Questions:   questionsDB,
		Img:         fileDB,
		DateChange:  itemAPI.DateChange,
		Url:         itemAPI.Url,
		VoteGroup:   voteGroup,
		Views:       itemAPI.Views,
		Counter:     itemAPI.Counter,
	}, nil
}

func convertQuestions(itemAPI []*models.QuestionAPI) (*models.ListOfQuestionDB, error) {
	var list *models.ListOfQuestionDB
	marshal, err := json.Marshal(itemAPI)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}

	return list, err
}

func convertVoteGroup(itemAPI *models.VoteGroupAPI) (*models.VoteGroupDB, error) {
	var list *models.VoteGroupDB
	marshal, err := json.Marshal(itemAPI)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}

	return list, err
}
