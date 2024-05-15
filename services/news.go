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

func (srv *Service) LoadNews(limit int, page int, repo *repository.Repository) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	loadStart := time.Now()
	var respondModel *models.NewsGQLRespond
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.NewsQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	loadedItemCount := len(respondModel.Data.News)
	fmt.Printf("Data loaded: count: %d time: %s\n", loadedItemCount, time.Since(loadStart))

	if loadedItemCount == 0 {
		return 0, nil
	}

	startSaveTime := time.Now()
	newsDB, err := ConvertNews(respondModel.Data.News)
	if err != nil {
		return 0, err
	}

	if err := repo.SetNews(newsDB); err != nil {
		return 0, err
	}
	fmt.Printf("Data saved: time: %s\n", time.Since(startSaveTime))
	return loadedItemCount, nil
}
func ConvertNews(newsAPI []*models.NewsAPI) ([]*models.NewsDB, error) {
	var newsDB []*models.NewsDB
	for _, news := range newsAPI {
		newsOneDB, err := ConvertOneNews(news)
		if err != nil {
			continue
		}
		newsDB = append(newsDB, newsOneDB)
	}
	return newsDB, nil
}

func ConvertOneNews(newsAPI *models.NewsAPI) (*models.NewsDB, error) {

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
		repostBlogPostId = newsAPI.RepostBlog.Id
	}

	tags, err := convertNewsTags(newsAPI.Tags)
	if err != nil {
		tags = nil
	}

	return &models.NewsDB{
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
		VoteIds:         changeType(newsAPI.VoteNum),
		FormId:          formIds,
		Tags:            tags,
	}, nil
}
func changeType(arr []int) []int64 {
	var arr2 []int64
	for i := 0; i < len(arr); i++ {
		arr2[i] = int64(arr[i])
	}
	return arr2
}
func convertNewsTags(tags []*models.Tag) ([]string, error) {
	var tagsDB []string
	for _, tag := range tags {
		tagsDB = append(tagsDB, tag.Name)
	}
	return tagsDB, nil
}
