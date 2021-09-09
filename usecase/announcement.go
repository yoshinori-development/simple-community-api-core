package usecase

import (
	"github.com/yoshinori-development/simple-community-api-core/domain/model"
	"github.com/yoshinori-development/simple-community-api-core/domain/repository"
)

type AnnouncementUsecase interface {
	List() ([]model.Announcement, error)
}

type announcementUsecase struct {
	AnnouncementRepository repository.AnnouncementRepository
}

type NewAnnouncementUsecaseInput struct {
	AnnouncementRepository repository.AnnouncementRepository
}

func NewAnnouncementUsecase(input NewAnnouncementUsecaseInput) *announcementUsecase {
	return &announcementUsecase{
		AnnouncementRepository: input.AnnouncementRepository,
	}
}

func (usecase *announcementUsecase) List() ([]model.Announcement, error) {
	announcements, err := usecase.AnnouncementRepository.List()
	if err != nil {
		return nil, err
	}
	return announcements, nil
}
