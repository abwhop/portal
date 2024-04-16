package services

import "github.com/abwhop/portal_sync"

type Service struct {
	config *portal_sync.Config
}

func NewService(config *portal_sync.Config) *Service {
	return &Service{
		config: config,
	}
}
