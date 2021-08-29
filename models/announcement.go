package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type AnnouncementModel struct {
	db *gorm.DB
}

type NewAnnouncementModelInput struct {
	DB *gorm.DB
}

func NewAnnouncementModel(input NewAnnouncementModelInput) *AnnouncementModel {
	return &AnnouncementModel{
		db: input.DB,
	}
}

type Announcement struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (model *AnnouncementModel) List() (*[]Announcement, error) {
	var announcements []Announcement

	result := model.db.Find(&announcements)
	if result.Error != nil {
		log.Print(result.Error)
	}
	return &announcements, nil
}
