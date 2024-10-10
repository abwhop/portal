package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/abwhop/portal_models/models"
	"github.com/abwhop/portal_sync/gql"
	"github.com/abwhop/portal_sync/query"
	"github.com/abwhop/portal_sync/repository"
	"strconv"
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
		if len(news.Comments) > 0 {
			var freshestComment *models.CommentAPI
			for _, comment := range news.Comments {
				if freshestComment == nil || comment.DateCreate > freshestComment.DateCreate {
					freshestComment = comment
				}
			}
			newsOneDB.FirstComment, err = convertComment(freshestComment)
			if err != nil {
				continue
			}
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
		likesDB = nil
	}
	firstCommentDB, err := ConvertComment(newsAPI.FirstComment)
	if err != nil {
		firstCommentDB = nil
	}

	viewsDB, err := ConvertViews(newsAPI.Views)
	if err != nil {
		viewsDB = nil
	}

	_, commentsDB, err := ConvertComments(newsAPI.Comments)
	if err != nil {
		commentsDB = nil
	}
	formIds, _, descriptionsDB, err := ConvertDescriptions(newsAPI.Text)
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

	calendarEvents, err := convertCalendarEvents(newsAPI.CalendarEvents)
	if err != nil {
		calendarEvents = nil
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
		Description:        newsAPI.Text,
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
		FirstComment:    firstCommentDB,
		Author:          authorDB,
		Likes:           likesDB,
		Views:           viewsDB,
		ReposBlogPostId: repostBlogPostId,
		Comments:        commentsDB,
		Files:           filesDB,
		VoteIds:         changeType(newsAPI.VoteNum),
		FormId:          formIds,
		Tags:            tags,
		CalendarEvents:  calendarEvents,
	}, nil
}
func changeType(arr []int) []int64 {
	arr2 := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		arr2[i] = int64(arr[i])
	}
	return arr2
}
func convertNewsTags(tags []*models.TagAPI) ([]string, error) {
	var tagsDB []string
	for _, tag := range tags {
		tagsDB = append(tagsDB, tag.Name)
	}
	return tagsDB, nil
}

func convertCalendarEvents(itemsAPI []*models.CalendarEventAPI) (*models.ListOfCalendarEventDB, error) {
	var listOfCalendarEventDB *models.ListOfCalendarEventDB
	var calendarEventsDB []*models.CalendarEventDB
	for _, eventAPI := range itemsAPI {
		eventDB, err := convertCalendarEvent(eventAPI)
		if err != nil {
			continue
		}
		calendarEventsDB = append(calendarEventsDB, eventDB)
	}
	marshal, err := json.Marshal(calendarEventsDB)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &listOfCalendarEventDB); err != nil {
		return nil, err
	}
	return listOfCalendarEventDB, nil
}
func convertCalendarEvent(itemAPI *models.CalendarEventAPI) (*models.CalendarEventDB, error) {

	dateStart, err := strconv.Atoi(itemAPI.DateStart)
	if err != nil {
		dateStart = 0
	}
	dateEnd, err := strconv.Atoi(itemAPI.DateEnd)
	if err != nil {
		dateEnd = 0
	}

	dateCreate, err := strconv.Atoi(itemAPI.DateCreate)
	if err != nil {
		dateCreate = 0
	}
	dateUpdate, err := strconv.Atoi(itemAPI.DateUpdate)
	if err != nil {
		dateUpdate = 0
	}

	return &models.CalendarEventDB{
		Id:          itemAPI.Id,
		Title:       itemAPI.Title,
		DateStart:   int64(dateStart),
		DateEnd:     int64(dateEnd),
		Description: itemAPI.Description,
		Location:    itemAPI.Location,
		SourceId:    itemAPI.SourceId,
		EntityType:  itemAPI.EntityType,
		CreatedBy:   itemAPI.CreatedBy,
		ModifiedBy:  itemAPI.ModifiedBy,
		DateCreate:  int64(dateCreate),
		DateUpdate:  int64(dateUpdate),
	}, nil
}
