package repositories

import (
	"fmt"
	"time"

	"errors"

	"github.com/yoshinori-development/simple-community-api-main/models"
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
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

func (repositories *ProfileRepository) Get(input models.ProfileRepositoryGetInput) (*models.Profile, error) {
	var profile models.Profile

	result := db.Where("sub = ?", input.Sub).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}

func (repositories *ProfileRepository) CreateOrUpdate(input models.ProfileRepositoryCreateOrUpdateInput) error {
	var profile models.Profile

	result := db.Where("sub = ?", input.Profile.Sub).First(&profile)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		createResult := db.Create(&Profile{
			Sub:      input.Profile.Sub,
			Nickname: input.Profile.Nickname,
			Age:      input.Profile.Age,
		})
		if createResult.Error != nil {
			return fmt.Errorf("failed to create profile: %w", result.Error)
		}
	} else if result.Error == nil {
		updateResult := db.Where("sub = ?", input.Profile.Sub).Updates(&Profile{
			Sub:      input.Profile.Sub,
			Nickname: input.Profile.Nickname,
			Age:      input.Profile.Age,
		})
		if updateResult.Error == nil {
			return fmt.Errorf("failed to update profile: %w", result.Error)
		}
	} else {
		return fmt.Errorf("failed to get profile: %w", result.Error)
	}

	return nil
}
