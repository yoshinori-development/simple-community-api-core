package repository

import "github.com/yoshinori-development/simple-community-api-core/domain/model"

type AnnouncementRepository interface {
	List() ([]model.Announcement, error)
}
