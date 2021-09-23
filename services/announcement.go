package services

import (
	"github.com/yoshinori-development/simple-community-api-main/models"
)

type AnnouncementService interface {
	List() ([]models.Announcement, error)
}

type announcementService struct {
	AnnouncementRepository models.AnnouncementRepository
}

type NewAnnouncementServiceInput struct {
	AnnouncementRepository models.AnnouncementRepository
}

func NewAnnouncementService(input NewAnnouncementServiceInput) *announcementService {
	return &announcementService{
		AnnouncementRepository: input.AnnouncementRepository,
	}
}

func (services *announcementService) List() ([]models.Announcement, error) {
	announcements, err := services.AnnouncementRepository.List()
	if err != nil {
		return nil, err
	}
	return announcements, nil
}
