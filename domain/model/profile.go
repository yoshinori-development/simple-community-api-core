package model

import (
	"time"
)

type Profile struct {
	ID        uint
	Sub       string
	Nickname  string
	Age       uint
	Birthdate string
	CreatedAt time.Time
	UpdatedAt time.Time
}
