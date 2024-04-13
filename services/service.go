package services

import "portal_sync/models"

type Service struct {
	config *models.Config
}

func NewService(config *models.Config) *Service {
	return &Service{
		config: config,
	}
}
