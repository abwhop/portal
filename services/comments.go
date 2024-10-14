package services

import (
	"encoding/json"
	"github.com/abwhop/portal_models/models"
)

func ConvertComment(commentAPI *models.CommentAPI) (*models.CommentDB, error) {
	authorDB, err := ConvertUser(commentAPI.Author)
	likesDB, err := ConvertLikes(commentAPI.Likes)
	if err != nil {
		authorDB = nil
	}
	return &models.CommentDB{
		Id:             commentAPI.Id,
		Text:           commentAPI.Text,
		SourceId:       commentAPI.SourceId,
		ParentSourceId: commentAPI.ParentSourceId,
		DateCreated:    commentAPI.DateCreate * 1000,
		Likes:          likesDB,
		Author:         authorDB,
	}, nil
}

func ConvertComments(commentsAPI []*models.CommentAPI) (int, *models.ListOfCommentDB, error) {
	var commentsDb []*models.CommentDB
	var list *models.ListOfCommentDB
	for _, commentAPI := range commentsAPI {
		if commentAPI.Likes == nil {
			commentAPI.Likes = new(models.LikesAPI)
		}
		commentDb, err := ConvertComment(commentAPI)
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
