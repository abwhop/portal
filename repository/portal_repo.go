package repository

import (
	"github.com/abwhop/portal_models/models"
	"github.com/abwhop/portal_sync/services"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(config *services.DatabaseConfig) *Repository {
	var db *gorm.DB
	sqlDebugFlag := os.Getenv("SQL_DEBUG")
	if sqlDebugFlag == "true" {
		db = getDB(config.Server, config.Database, config.User, config.Password).Debug()
	} else {
		db = getDB(config.Server, config.Database, config.User, config.Password)
	}
	return &Repository{
		db: db,
	}
}

func (b *Repository) SetBlogs(resModel []*models.BlogDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}
func (b *Repository) SetBlogPosts(resModel []*models.PostDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}
func (b *Repository) SetNews(resModel []*models.NewsBreafe) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}

func (b *Repository) SetAppointment(resModel []*models.AppointmentsDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}
func (b *Repository) SetCommunity(resModel []*models.CommunityDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}

func (b *Repository) SetCommunityTypes(resModel []*models.CommunityTypeDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}

func (b *Repository) SetCommunitySubjects(resModel []*models.CommunitySubjectDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}

func (b *Repository) SetUsers(resModel []*models.UserFullDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}

func (b *Repository) SetForms(resModel []*models.FormDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}

func (b *Repository) SetVotes(resModel []*models.VoteDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}
func (b *Repository) SetVoteResults(resModel []*models.VoteResultDB) error {
	return b.db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&resModel).Error
}
