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

func (srv *Service) LoadCommunities(limit int, page int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	repo := repository.NewRepository(srv.config.Database)
	var err error
	loafStart := time.Now()
	var respondModel *models.CommunityGQLRespond
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.CommunityQuery, limit, page), &respondModel); err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Data loaded:", time.Since(loafStart))
	startSaveTime := time.Now()
	if len(respondModel.Data.Workgroups) == 0 {
		return 0, nil
	}
	items, err := convertCommunities(respondModel.Data.Workgroups)
	if err != nil {
		return 0, err
	}

	var types []*models.CommunityTypeDB
	var subjects []*models.CommunitySubjectDB
	uniqTypesList := make(map[string]*models.CommunityTypeDB)
	uniqSubjectsList := make(map[int]*models.CommunitySubjectDB)
	for _, item := range items {
		if item.Type != nil {
			uniqTypesList[item.Type.Code] = item.Type
		}
		if item.Subject != nil {
			uniqSubjectsList[item.Subject.Id] = item.Subject
		}

	}
	for _, value := range uniqTypesList {
		types = append(types, value)
	}
	for _, value := range uniqSubjectsList {
		subjects = append(subjects, value)
	}
	if err := repo.SetCommunityTypes(types); err != nil {
		return 0, err
	}
	if err := repo.SetCommunitySubjects(subjects); err != nil {
		return 0, err
	}
	if err := repo.SetCommunity(items); err != nil {
		return 0, err
	}
	fmt.Println("Data saved:", time.Since(startSaveTime))
	return len(respondModel.Data.Workgroups), nil
}

func convertCommunities(newsAPI []*models.CommunityAPI) ([]*models.CommunityDB, error) {
	var items []*models.CommunityDB
	for _, item := range newsAPI {
		itemDB, err := convertCommunity(item)
		if err != nil {
			continue
		}
		items = append(items, itemDB)
	}
	return items, nil
}

func convertCommunity(communityAPI *models.CommunityAPI) (*models.CommunityDB, error) {

	authorDB, err := ConvertUser(communityAPI.Author)
	if err != nil {
		authorDB = nil
	}
	filesDB, err := ConvertFiles(communityAPI.Files)
	if err != nil {
		filesDB = nil
	}
	communityType, err := ConvertCommunityType(communityAPI.Type)
	if err != nil {
		communityType = nil
	}
	communitySubject, err := ConvertCommunitySubject(communityAPI.Subject)
	if err != nil {
		communitySubject = nil
	}
	members, err := ConvertUsers(communityAPI.Members)
	if err != nil {
		members = nil
	}
	moderators, err := ConvertUsers(communityAPI.Moderators)
	if err != nil {
		moderators = nil
	}
	favorites, err := ConvertUsers(communityAPI.Favorites)
	if err != nil {
		favorites = nil
	}

	return &models.CommunityDB{
		Id:          communityAPI.Id,
		Name:        communityAPI.Name,
		Description: communityAPI.Description,
		Active:      communityAPI.Active,
		DateCreated: communityAPI.DateCreate,
		Img:         communityAPI.Img,
		Closed:      communityAPI.Closed,
		Visible:     communityAPI.Visible,
		Opened:      communityAPI.Opened,
		Project:     communityAPI.Project,
		Author:      authorDB,
		Type:        communityType,
		Subject:     communitySubject,
		Files:       filesDB,
		//CountMembers:    len(communityAPI.Members),
		Members:         members,
		Moderators:      moderators,
		Favorites:       favorites,
		UserIsMember:    false,
		GroupIsFavorite: false,
	}, nil
}

func ConvertCommunityType(communityTypeAPI *models.CommunityTypeAPI) (*models.CommunityTypeDB, error) {
	if communityTypeAPI == nil {
		return nil, nil
	}
	return &models.CommunityTypeDB{
		Code:        communityTypeAPI.Code,
		Name:        communityTypeAPI.Name,
		Description: communityTypeAPI.Description,
	}, nil
}

func ConvertCommunitySubject(communitySubjectAPI *models.CommunitySubjectAPI) (*models.CommunitySubjectDB, error) {
	if communitySubjectAPI == nil {
		return nil, nil
	}
	return &models.CommunitySubjectDB{
		Id:   communitySubjectAPI.Id,
		Name: communitySubjectAPI.Name,
		Sort: communitySubjectAPI.Sort,
	}, nil
}
