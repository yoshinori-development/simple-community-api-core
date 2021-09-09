package db_core

import (
	"log"
	"time"

	"github.com/yoshinori-development/simple-community-api-core/domain/model"
	"gorm.io/gorm"
)

type Announcement struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AnnouncementRepository struct {
	db *gorm.DB
}

func NewAnnouncementRepository(db *gorm.DB) *AnnouncementRepository {
	return &AnnouncementRepository{
		db: db,
	}
}

func (repository *AnnouncementRepository) List() ([]model.Announcement, error) {
	var announcements []Announcement

	result := repository.db.Find(&announcements)
	if result.Error != nil {
		log.Print(result.Error)
	}

	var modelAnnouncements []model.Announcement
	for _, v := range announcements {
		modelAnnouncements = append(modelAnnouncements, repository.convertToModel(v))
	}
	return modelAnnouncements, nil
}

func (repository *AnnouncementRepository) convertToModel(announcement Announcement) model.Announcement {
	var modelAnnouncement model.Announcement

	modelAnnouncement.ID = announcement.ID
	modelAnnouncement.Title = announcement.Title
	modelAnnouncement.Content = announcement.Content
	modelAnnouncement.CreatedAt = announcement.CreatedAt
	modelAnnouncement.UpdatedAt = announcement.UpdatedAt

	return modelAnnouncement
}
