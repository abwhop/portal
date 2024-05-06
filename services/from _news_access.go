package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/abwhop/portal_models/models"
	"github.com/abwhop/portal_sync/gql"
	"github.com/abwhop/portal_sync/query"
	"time"
)

func (srv *Service) UserSubscribedNewsRubric(bitrixUserId int) ([]*models.RubricDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var rubrics []*models.RubricDB
	var respondModel struct {
		Data struct {
			Users []*struct {
				Rubrics *struct {
					News []*models.RubricDB `json:"news"`
				} `json:"rubrics"`
			} `json:"users"`
		} `json:"data"`
	}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.UserSubscribedRubric, bitrixUserId), &respondModel); err != nil {
		return nil, err
	}
	for _, rb := range respondModel.Data.Users {
		rubrics = append(rubrics, rb.Rubrics.News...)
	}
	return rubrics, nil
}

func (srv *Service) UnSubscribeRubric(itemId int, bitrixUserId int) ([]*models.RubricDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var respondModel struct {
		Data struct {
			Users []*struct {
				Rubrics *struct {
					News []*models.RubricDB `json:"news"`
				} `json:"rubrics"`
			} `json:"users"`
		} `json:"data"`
	}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.UnSubscribeRubric, bitrixUserId, itemId), &respondModel); err != nil {
		return nil, err
	}
	//if respondModel.Data.UnsubscribeRubric.Success {
	return srv.UserSubscribedNewsRubric(bitrixUserId)
	//}
	//return nil, nil
}

func (srv *Service) SubscribeRubric(itemId int, bitrixUserId int) ([]*models.RubricDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var respondModel struct {
		Data struct {
			UnsubscribeRubric struct {
				Success bool   `json:"success"`
				Message string `json:"message"`
			} `json:"unsubscribeRubric"`
		} `json:"data"`
	}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.SubscribeRubric, bitrixUserId, itemId), &respondModel); err != nil {
		return nil, err
	}
	//if respondModel.Data.UnsubscribeRubric.Success {
	return srv.UserSubscribedNewsRubric(bitrixUserId)
	//}
	//return nil, nil
}

func (srv *Service) AddFavorite(itemId int, bitrixUserId int) ([]int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var result []int
	var respondModel struct {
		Data struct {
			AddFavorites struct {
				News []*struct {
					Id int `json:"id"`
				} `json:"news"`
			} `json:"addFavorites"`
		} `json:"data"`
	}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.AddFavoriteQuery, bitrixUserId, itemId), &respondModel); err != nil {
		return nil, err
	}
	for _, item := range respondModel.Data.AddFavorites.News {
		result = append(result, item.Id)
	}
	return result, nil
}

func (srv *Service) RemoveFavorite(itemId int, bitrixUserId int) ([]int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var result []int
	var respondModel struct {
		Data struct {
			AddFavorites struct {
				News []*struct {
					Id int `json:"id"`
				} `json:"news"`
			} `json:"addFavorites"`
		} `json:"data"`
	}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.RemoveFavoriteQuery, itemId, bitrixUserId), &respondModel); err != nil {
		return nil, err
	}
	for _, item := range respondModel.Data.AddFavorites.News {
		result = append(result, item.Id)
	}
	return result, nil
}

func (srv *Service) SetComment(itemId int, commentText string, bitrixUserId int) (*models.ListOfCommentDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var respondModel struct {
		Data struct {
			Comments []*models.CommentAPI `json:"setComment"`
		} `json:"data"`
	}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.SetCommentQuery, itemId, commentText, bitrixUserId), &respondModel); err != nil {
		return nil, err
	}
	_, comments, err := convertComments(respondModel.Data.Comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (srv *Service) SetLike(itemId int, isLiked bool, bitrixUserId int) (*models.LikesDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var respondModel struct {
		Data struct {
			Likes *models.LikesAPI `json:"setLike"`
		} `json:"data"`
	}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.SetLikeQuery, itemId, bitrixUserId, !isLiked), &respondModel); err != nil {
		return nil, err
	}
	return convertLikes(respondModel.Data.Likes)
}

func (srv *Service) SetView(itemId int, bitrixUserId int) (*models.ViewsDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	var respondModel struct {
		Data struct {
			View *models.ViewsAPI `json:"setView"`
		} `json:"data"`
	}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.SetViewQuery, itemId, bitrixUserId), &respondModel); err != nil {
		return nil, err
	}
	return convertViews(respondModel.Data.View)
}

func (srv *Service) SubscribeToBlog(authorId int, subscriberId int, subscribeFlag bool) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	var respondModel interface{}
	if err := gql.NewGql(srv.config.Portal).Query(ctx, fmt.Sprintf(query.SubscribeBlogQuery, authorId, subscriberId, subscribeFlag), &respondModel); err != nil {
		return nil, err
	}
	return respondModel, nil
}

func convertViews(viewsAPI *models.ViewsAPI) (*models.ViewsDB, error) {
	if viewsAPI == nil {
		return nil, nil
	}
	usersDB, err := convertUsers(viewsAPI.Users)
	if err != nil {
		usersDB = nil
	}
	return &models.ViewsDB{
		Count: viewsAPI.Count,
		Users: usersDB,
	}, nil
}

func convertUsers(usersAPI []*models.UserAPI) (*models.ListOfUsersDB, error) {
	var usersDB []*models.UserDB
	var list *models.ListOfUsersDB
	for _, userAPI := range usersAPI {
		userDB, err := convertUser(userAPI)
		if err != nil {
			continue
		}
		usersDB = append(usersDB, userDB)
	}
	marshal, err := json.Marshal(usersDB)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}
	return list, nil
}
func convertUser(userAPI *models.UserAPI) (*models.UserDB, error) {
	gender := "men"
	if userAPI.Gender == "F" {
		gender = "women"
	}
	return &models.UserDB{
		BitrixUserId:    userAPI.Id,
		FirstName:       userAPI.Name,
		LastName:        userAPI.LastName,
		MiddleName:      userAPI.SecondName,
		Email:           userAPI.Email,
		PositionName:    userAPI.Position,
		PersonnelNumber: userAPI.PersonalNumber,
		Active:          userAPI.Active == "Y",
		Gender:          gender,
		Photo:           userAPI.Photo,
		LoginAd:         userAPI.LoginAd,
	}, nil
}
func convertLikes(likesAPI *models.LikesAPI) (*models.LikesDB, error) {
	usersDB, err := convertUsers(likesAPI.Users)
	if err != nil {
		usersDB = nil
	}
	return &models.LikesDB{
		Count: likesAPI.Count,
		Users: usersDB,
	}, nil
}

func convertComments(commentsAPI []*models.CommentAPI) (int, *models.ListOfCommentDB, error) {
	var commentsDb []*models.CommentDB
	var list *models.ListOfCommentDB
	for _, commentAPI := range commentsAPI {
		commentDb, err := convertComment(commentAPI)
		if err != nil {
			continue
		}
		commentsDb = append(commentsDb, commentDb)
	}
	marshal, err := json.Marshal(commentsDb)
	if err != nil {
		return 0, nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return 0, nil, err
	}
	return len(commentsDb), list, nil
}

func convertComment(commentAPI *models.CommentAPI) (*models.CommentDB, error) {
	authorDB, err := convertUser(commentAPI.Author)
	if err != nil {
		authorDB = nil
	}
	return &models.CommentDB{
		Id:          commentAPI.Id,
		Text:        commentAPI.Text,
		DateCreated: commentAPI.DateCreate * 1000,
		Author:      authorDB,
	}, nil
}
