package services

import (
	"gitlab.com/kirill_ussr/portal_sync/models"
)

func ConvertLikes(likesAPI *models.LikesAPI) (*models.LikesDB, error) {
	usersDB, err := ConvertUsers(likesAPI.Users)
	if err != nil {
		usersDB = nil
	}
	return &models.LikesDB{
		Count: likesAPI.Count,
		Users: usersDB,
	}, nil
}
