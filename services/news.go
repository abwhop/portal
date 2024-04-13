package services

import (
	"context"
	"fmt"
	"portal_sync/gql"
	"portal_sync/models"
	"portal_sync/query"
	"portal_sync/repository"
	"time"
)

func (srv *Service) LoadNews(limit int, page int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	repo := repository.NewRepository(srv.config.Database)
	var err error
	loafStart := time.Now()
	var respondModel *models.NewsGQLRespond

	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.NewsQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Data loaded:", time.Since(loafStart))
	startSaveTime := time.Now()

	newsDB, err := ConvertNews(respondModel.Data.News)
	if err != nil {
		return 0, err
	}

	if err := repo.SetNews(newsDB); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.News), nil
}
func ConvertNews(newsAPI []*models.NewsAPI) ([]*models.NewsBreafe, error) {
	var newsDB []*models.NewsBreafe
	for _, news := range newsAPI {
		newsOneDB, err := ConvertOneNews(news)
		if err != nil {
			continue
		}
		newsDB = append(newsDB, newsOneDB)
	}
	return newsDB, nil
}

func ConvertOneNews(newsAPI *models.NewsAPI) (*models.NewsBreafe, error) {

	authorDB, err := ConvertUser(newsAPI.Author)
	if err != nil {
		authorDB = nil
	}
	filesDB, err := ConvertFiles(newsAPI.Files)
	if err != nil {
		filesDB = nil
	}
	likesDB, err := ConvertLikes(newsAPI.Likes)
	if err != nil {
		filesDB = nil
	}

	viewsDB, err := ConvertViews(newsAPI.Views)
	if err != nil {
		viewsDB = nil
	}

	_, commentsDB, err := ConvertComments(newsAPI.Comments)
	if err != nil {
		commentsDB = nil
	}
	formIds, convertedText, descriptionsDB, err := ConvertDescriptions(newsAPI.Text)
	if err != nil {
		descriptionsDB = nil
	}

	repostBlogPostId := 0
	if newsAPI.RepostBlog != nil {
		repostBlogPostId = newsAPI.RepostBlog.BlogId
	}

	/*repostNewsId := 0
	if newsAPI.RepostNews != nil {
		repostNewsId = newsAPI.RepostNews.Id
	}*/

	return &models.NewsBreafe{
		Id:                 newsAPI.Id,
		Type:               "news",
		PublishDate:        newsAPI.PublishDate * 1000,
		CreateDate:         newsAPI.CreateDate * 1000,
		Title:              newsAPI.Name,
		LogId:              newsAPI.LogId,
		CanComment:         newsAPI.CanComment,
		Descriptions:       descriptionsDB,
		Description:        convertedText,
		Published:          newsAPI.Published,
		Rights:             newsAPI.Rights,
		ImageUrl:           newsAPI.Img,
		PreviewDescription: newsAPI.PreviewText,
		//XmlId:              newsAPI.XmlId,
		//SliderFile:   newsAPI.SliderFile,
		Rubric: &models.RubricDB{
			Name: newsAPI.Rubric.Name,
			Id:   newsAPI.Rubric.Id,
			Code: newsAPI.Rubric.Code,
		},
		Author:          authorDB,
		Likes:           likesDB,
		Views:           viewsDB,
		ReposBlogPostId: repostBlogPostId,
		Comments:        commentsDB,
		Files:           filesDB,
		VoteNum:         nil,
		FormId:          formIds,
	}, nil
}
