package services

import (
	"gitlab.com/kirill_ussr/portal_sync/models"
)

func ConvertViews(viewsAPI *models.ViewsAPI) (*models.ViewsDB, error) {
	if viewsAPI == nil {
		return nil, nil
	}
	usersDB, err := ConvertUsers(viewsAPI.Users)
	if err != nil {
		usersDB = nil
	}
	return &models.ViewsDB{
		Count: viewsAPI.Count,
		Users: usersDB,
	}, nil
}
