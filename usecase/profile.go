package usecase

import (
	"github.com/yoshinori-development/simple-community-api-core/domain/model"
	"github.com/yoshinori-development/simple-community-api-core/domain/repository"
)

type ProfileUsecase interface {
	Get(ProfileUsecaseGetInput) (*model.Profile, error)
	CreateOrUpdate(ProfileUsecaseCreateOrUpdateInput) error
}

type profileUsecase struct {
	ProfileRepository repository.ProfileRepository
}

type NewProfileUsecaseInput struct {
	ProfileRepository repository.ProfileRepository
}

func NewProfileUsecase(input NewProfileUsecaseInput) *profileUsecase {
	return &profileUsecase{
		ProfileRepository: input.ProfileRepository,
	}
}

type ProfileUsecaseGetInput struct {
	Sub string
}

func (usecase *profileUsecase) Get(input ProfileUsecaseGetInput) (*model.Profile, error) {
	profile, err := usecase.ProfileRepository.Get(repository.ProfileRepositoryGetInput{
		Sub: input.Sub,
	})
	if err != nil {
		return nil, err
	}
	return profile, nil
}

type ProfileUsecaseCreateOrUpdateInput struct {
	Profile model.Profile
}

func (usecase *profileUsecase) CreateOrUpdate(input ProfileUsecaseCreateOrUpdateInput) error {
	err := usecase.ProfileRepository.CreateOrUpdate(repository.ProfileRepositoryCreateOrUpdateInput{
		Profile: model.Profile{
			Sub:       input.Profile.Sub,
			Nickname:  input.Profile.Nickname,
			Age:       input.Profile.Age,
			Birthdate: input.Profile.Birthdate,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
