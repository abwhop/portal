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

func (srv *Service) LoadAppointments(limit int, page int, repo *repository.Repository) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	loadStart := time.Now()
	var respondModel *models.AppointmentsGQLRespond
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.AppointmentsQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	loadedItemCount := len(respondModel.Data.Appointments)
	fmt.Printf("Data loaded: count: %d time: %s\n", loadedItemCount, time.Since(loadStart))

	if loadedItemCount == 0 {
		return 0, nil
	}
	startSaveTime := time.Now()
	appointments, err := convertAppointments(respondModel.Data.Appointments)
	if err != nil {
		return 0, err
	}
	if err := repo.SetAppointment(appointments); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return loadedItemCount, nil
}

func convertAppointments(newsAPI []*models.AppointmentsAPI) ([]*models.AppointmentsDB, error) {
	var items []*models.AppointmentsDB
	for _, item := range newsAPI {
		itemDB, err := convertAppointment(item)
		if err != nil {
			continue
		}
		if len(item.Comments) > 0 {
			var freshestComment *models.CommentAPI
			for _, comment := range item.Comments {
				if freshestComment == nil || comment.DateCreate > freshestComment.DateCreate {
					freshestComment = comment
				}
			}
			itemDB.FirstComment, err = convertComment(freshestComment)
			if err != nil {
				continue
			}
		}
		items = append(items, itemDB)
	}
	return items, nil
}

func convertAppointment(newsAPI *models.AppointmentsAPI) (*models.AppointmentsDB, error) {

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

	var firstCommentDB *models.CommentDB
	if newsAPI.FirstComment != nil {
		firstCommentDB, err = ConvertComment(newsAPI.FirstComment)
		if err != nil {
			firstCommentDB = nil
		}
	}
	viewsDB, err := ConvertViews(newsAPI.Views)
	if err != nil {
		viewsDB = nil
	}
	_, commentsDB, err := ConvertComments(newsAPI.Comments)
	if err != nil {
		commentsDB = nil
	}
	_, convertedText, descriptionsDB, err := ConvertDescriptions(newsAPI.Text)
	if err != nil {
		descriptionsDB = nil
	}

	return &models.AppointmentsDB{
		Id:                 newsAPI.Id,
		Type:               "appointment",
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
		FirstComment: firstCommentDB,
		Author:       authorDB,
		Likes:        likesDB,
		Views:        viewsDB,
		Comments:     commentsDB,
		Files:        filesDB,
	}, nil
}
