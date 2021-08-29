package models

import (
	"time"

	"gorm.io/gorm"
)

type ProfileModel struct {
	db *gorm.DB
}

type NewProfileModelInput struct {
	DB *gorm.DB
}

func NewProfileModel(input NewProfileModelInput) *ProfileModel {
	return &ProfileModel{
		db: input.DB,
	}
}

type Profile struct {
	ID        uint `gorm:"primarykey"`
	Sub       string
	Nickname  string
	Age       uint
	Birthdate string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProfileModelGetInput struct {
	Sub string
}

func (model *ProfileModel) Get(input ProfileModelGetInput) (*Profile, error) {
	var profile Profile

	result := model.db.Where("sub = ?", input.Sub).First(&profile)
	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}

type ProfileModelCreateInput struct {
	Sub       string
	Nickname  string
	Age       uint
	Birthdate string
}

func (model *ProfileModel) Create(input ProfileModelCreateInput) error {
	model.db.Create(&Profile{
		Sub:       input.Sub,
		Nickname:  input.Nickname,
		Age:       input.Age,
		Birthdate: input.Birthdate,
	})
	return nil
}
