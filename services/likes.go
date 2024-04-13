package services

import (
	"git.nlmk.com/mcs/micro/portal/portal_sync/models"
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
