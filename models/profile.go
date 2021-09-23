package models

import (
	"time"
)

type ProfileRepositoryGetInput struct {
	Sub string
}

type ProfileRepositoryCreateOrUpdateInput struct {
	Profile Profile
}

type ProfileRepository interface {
	Get(ProfileRepositoryGetInput) (*Profile, error)
	CreateOrUpdate(ProfileRepositoryCreateOrUpdateInput) error
}

type Profile struct {
	ID        uint
	Sub       string
	Nickname  string
	Age       uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
