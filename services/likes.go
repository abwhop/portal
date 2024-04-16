package services

import (
	"github.com/abwhop/portal_models/models"
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
