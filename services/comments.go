package services

import (
	"encoding/json"
	"github.com/abwhop/portal_sync/models"
)

func ConvertComment(commentAPI *models.CommentAPI) (*models.CommentDB, error) {
	authorDB, err := ConvertUser(commentAPI.Author)
	if err != nil {
		authorDB = nil
	}
	return &models.CommentDB{
		Id:          commentAPI.Id,
		Text:        commentAPI.Text,
		DateCreated: commentAPI.DateCreated * 1000,
		Author:      authorDB,
	}, nil
}

func ConvertComments(commentsAPI []*models.CommentAPI) (int, *models.ListOfCommentDB, error) {
	var commentsDb []*models.CommentDB
	var list *models.ListOfCommentDB
	for _, commentAPI := range commentsAPI {
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
