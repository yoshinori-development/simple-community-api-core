package repositories

import (
	"log"
	"time"

	"github.com/yoshinori-development/simple-community-api-main/models"
)

type Announcement struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AnnouncementRepository struct {
}

func NewAnnouncementRepository() *AnnouncementRepository {
	return &AnnouncementRepository{}
}

func (repositories *AnnouncementRepository) List() ([]models.Announcement, error) {
	var announcements []Announcement

	result := db.Find(&announcements)
	if result.Error != nil {
		log.Print(result.Error)
	}

	var modelAnnouncements []models.Announcement
	for _, v := range announcements {
		modelAnnouncements = append(modelAnnouncements, repositories.convertToModel(v))
	}
	return modelAnnouncements, nil
}

func (repositories *AnnouncementRepository) convertToModel(announcement Announcement) models.Announcement {
	var modelAnnouncement models.Announcement

	modelAnnouncement.ID = announcement.ID
	modelAnnouncement.Title = announcement.Title
	modelAnnouncement.Content = announcement.Content
	modelAnnouncement.CreatedAt = announcement.CreatedAt
	modelAnnouncement.UpdatedAt = announcement.UpdatedAt

	return modelAnnouncement
}

func (repositories *AnnouncementRepository) Create(input models.AnnouncementRepositoryCreateInput) error {
	db.Create(&input.Announcement)

	return nil
}
