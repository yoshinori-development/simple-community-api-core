package services

import (
	"github.com/yoshinori-development/simple-community-api-main/models"
)

type ProfileService interface {
	Get(ProfileServiceGetInput) (*models.Profile, error)
	CreateOrUpdate(ProfileServiceCreateOrUpdateInput) error
}

type profileService struct {
	ProfileRepository models.ProfileRepository
}

type NewProfileServiceInput struct {
	ProfileRepository models.ProfileRepository
}

func NewProfileService(input NewProfileServiceInput) *profileService {
	return &profileService{
		ProfileRepository: input.ProfileRepository,
	}
}

type ProfileServiceGetInput struct {
	Sub string
}

func (services *profileService) Get(input ProfileServiceGetInput) (*models.Profile, error) {
	profile, err := services.ProfileRepository.Get(models.ProfileRepositoryGetInput{
		Sub: input.Sub,
	})
	if err != nil {
		return nil, err
	}
	return profile, nil
}

type ProfileServiceCreateOrUpdateInput struct {
	Profile models.Profile
}

func (services *profileService) CreateOrUpdate(input ProfileServiceCreateOrUpdateInput) error {
	err := services.ProfileRepository.CreateOrUpdate(models.ProfileRepositoryCreateOrUpdateInput{
		Profile: models.Profile{
			Sub:      input.Profile.Sub,
			Nickname: input.Profile.Nickname,
			Age:      input.Profile.Age,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
