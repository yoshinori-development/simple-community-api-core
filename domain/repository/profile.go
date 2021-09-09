package repository

import "github.com/yoshinori-development/simple-community-api-core/domain/model"

type ProfileRepositoryGetInput struct {
	Sub string
}

type ProfileRepositoryCreateOrUpdateInput struct {
	Profile model.Profile
}

type ProfileRepository interface {
	Get(ProfileRepositoryGetInput) (*model.Profile, error)
	CreateOrUpdate(ProfileRepositoryCreateOrUpdateInput) error
}
