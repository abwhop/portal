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

func (srv *Service) LoadVotes(limit int, page int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	loadStart := time.Now()
	var respondModel models.ViteGQLRespond
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.VotesQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	loadedItemCount := len(respondModel.Data.Votes)
	fmt.Printf("Data loaded: count: %d time: %s\n", loadedItemCount, time.Since(loadStart))

	if loadedItemCount == 0 {
		return 0, nil
	}
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
		DateFrom:    itemAPI.DateFrom * 1000,
		DateTo:      itemAPI.DateTo * 1000,
		Questions:   questionsDB,
		Img:         fileDB,
		DateChange:  itemAPI.DateChange,
		Url:         itemAPI.Url,
		VoteGroup:   voteGroup,
		Views:       itemAPI.Views,
		Counter:     itemAPI.Counter,
	}, nil
}

func convertQuestions(itemsAPI []*models.QuestionAPI) (*models.ListOfQuestionDB, error) {
	var itemsDB []*models.QuestionDB
	var list *models.ListOfQuestionDB
	for _, item := range itemsAPI {
		itemDB, err := convertQuestion(item)
		if err != nil {
			continue
		}
		itemsDB = append(itemsDB, itemDB)
	}
	marshal, err := json.Marshal(itemsDB)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func convertQuestion(itemAPI *models.QuestionAPI) (*models.QuestionDB, error) {
	answerDB, err := convertAnswers(itemAPI.Answers)
	if err != nil {
		answerDB = nil
	}
	return &models.QuestionDB{
		Id:           itemAPI.Id,
		Sort:         itemAPI.Sort,
		Question:     itemAPI.Question,
		DateChange:   itemAPI.DateChange,
		Active:       itemAPI.Active,
		Counter:      itemAPI.Counter,
		Diagram:      itemAPI.Diagram,
		Required:     itemAPI.Required,
		DiagramType:  itemAPI.DiagramType,
		QuestionType: itemAPI.QuestionType,
		Answers:      answerDB,
	}, nil
}

func convertAnswers(itemsAPI []*models.AnswerAPI) (*models.ListOfAnswerDB, error) {
	var itemsDB []*models.AnswerDB
	var list *models.ListOfAnswerDB
	for _, item := range itemsAPI {
		itemDB, err := convertAnswer(item)
		if err != nil {
			continue
		}
		itemsDB = append(itemsDB, itemDB)
	}
	marshal, err := json.Marshal(itemsDB)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func convertAnswer(itemAPI *models.AnswerAPI) (*models.AnswerDB, error) {
	return &models.AnswerDB{
		Id:         itemAPI.Id,
		Sort:       itemAPI.Sort,
		Message:    itemAPI.Message,
		FieldType:  itemAPI.FieldType,
		DateChange: itemAPI.DateChange,
		Active:     itemAPI.Active,
		Counter:    itemAPI.Counter,
	}, nil
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
