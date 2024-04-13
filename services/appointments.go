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

func (srv *Service) LoadAppointments(limit int, page int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var err error
	loafStart := time.Now()
	var respondModel *models.AppointmentsGQLRespond

	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.AppointmentsQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Data loaded:", time.Since(loafStart))
	startSaveTime := time.Now()

	appointments, err := convertAppointments(respondModel.Data.Appointments)
	if err != nil {
		return 0, err
	}

	if err := repository.NewRepository(srv.config.Database).SetAppointment(appointments); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.Appointments), nil
}

func convertAppointments(newsAPI []*models.AppointmentsAPI) ([]*models.AppointmentsDB, error) {
	var items []*models.AppointmentsDB
	for _, item := range newsAPI {
		itemDB, err := convertAppointment(item)
		if err != nil {
			continue
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
		Author:   authorDB,
		Likes:    likesDB,
		Views:    viewsDB,
		Comments: commentsDB,
		Files:    filesDB,
	}, nil
}
