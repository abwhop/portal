package services

import (
	"context"
	"fmt"
	"github.com/abwhop/portal_models/models"
	"github.com/abwhop/portal_sync/gql"
	"github.com/abwhop/portal_sync/query"
	"github.com/abwhop/portal_sync/repository"
	"time"
)

func (srv *Service) LoadVoteResults(limit int, page int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	loafStart := time.Now()

	var respondModel models.ViteResultGQLRespond

	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.VotesResultsQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Data loaded:", time.Since(loafStart))
	startSaveTime := time.Now()

	itemsDB, err := convertVoteResults(respondModel.Data.VoteResults)
	if err != nil {
		return 0, err
	}

	if err := repository.NewRepository(srv.config.Database).SetVoteResults(itemsDB); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.VoteResults), nil
}

func convertVoteResults(itemsAPI []*models.VoteResultAPI) ([]*models.VoteResultDB, error) {
	var itemsDB []*models.VoteResultDB
	for _, item := range itemsAPI {
		itemDB, err := convertVoteResult(item)
		if err != nil {
			continue
		}
		itemsDB = append(itemsDB, itemDB)
	}
	return itemsDB, nil
}

func convertVoteResult(itemAPI *models.VoteResultAPI) (*models.VoteResultDB, error) {
	userDB, err := ConvertUser(itemAPI.User)
	if err != nil {
		userDB = nil
	}
	return &models.VoteResultDB{
		Id:     itemAPI.Id,
		Date:   itemAPI.Date,
		User:   userDB,
		VoteId: itemAPI.VoteId,
	}, nil
}
