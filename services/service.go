package services

import (
	"github.com/abwhop/portal_models/models"
)

type Service struct {
	config *models.Config
}

func NewService(config *models.Config) *Service {
	return &Service{
		config: config,
	}
}
