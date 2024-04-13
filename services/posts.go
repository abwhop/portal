package services

import (
	"context"
	"fmt"
	"gitlab.com/kirill_ussr/portal_sync/gql"
	"gitlab.com/kirill_ussr/portal_sync/models"
	"gitlab.com/kirill_ussr/portal_sync/query"
	"gitlab.com/kirill_ussr/portal_sync/repository"
	"time"
)

func (srv *Service) LoadPosts(limit int, page int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	loafStart := time.Now()
	var respondModel *models.PostGQLRespond
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.BlogPostQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Data loaded:", time.Since(loafStart))
	startSaveTime := time.Now()

	blogDB, err := ConvertBlogPosts(respondModel.Data.BlogPosts)
	if err != nil {
		return 0, err
	}

	if err := repository.NewRepository(srv.config.Database).SetBlogPosts(blogDB); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.BlogPosts), nil
}
