package model

import (
	"time"
)

type Announcement struct {
	ID        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
