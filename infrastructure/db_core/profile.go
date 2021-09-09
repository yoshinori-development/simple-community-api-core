package db_core

import (
	"errors"
	"fmt"
	"time"

	"github.com/yoshinori-development/simple-community-api-core/domain/model"
	"github.com/yoshinori-development/simple-community-api-core/domain/repository"
	"gorm.io/gorm"
)

type Profile struct {
	ID        uint `gorm:"primarykey"`
	Sub       string
	Nickname  string
	Age       uint
	Birthdate string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (repository *ProfileRepository) Get(input repository.ProfileRepositoryGetInput) (*model.Profile, error) {
	var profile model.Profile

	result := repository.db.Where("sub = ?", input.Sub).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}

func (repository *ProfileRepository) CreateOrUpdate(input repository.ProfileRepositoryCreateOrUpdateInput) error {
	var profile model.Profile

	result := repository.db.Where("sub = ?", input.Profile.Sub).First(&profile)
	fmt.Print(result.Error)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		repository.db.Create(&Profile{
			Sub:       input.Profile.Sub,
			Nickname:  input.Profile.Nickname,
			Age:       input.Profile.Age,
			Birthdate: input.Profile.Birthdate,
		})
	} else if result.Error == nil {
		repository.db.Where("sub = ?", input.Profile.Sub).Updates(&Profile{
			Sub:       input.Profile.Sub,
			Nickname:  input.Profile.Nickname,
			Age:       input.Profile.Age,
			Birthdate: input.Profile.Birthdate,
		})
	} else {
		return result.Error
	}

	return nil
}
