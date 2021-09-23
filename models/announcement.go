package models

import (
	"time"
)

type AnnouncementRepository interface {
	List() ([]Announcement, error)
}

type AnnouncementRepositoryCreateInput struct {
	Announcement Announcement
}

type Announcement struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
