package services

import (
	"fmt"
	"github.com/abwhop/portal_sync/repository"
	"time"
)

func (srv *Service) RefreshFeed(repo *repository.Repository) error {
	startSaveTime := time.Now()
	if err := repo.RefreshFeed(); err != nil {
		return err
	}
	fmt.Printf("Data Feed refreshed: time: %s\n", time.Since(startSaveTime))
	return nil
}
