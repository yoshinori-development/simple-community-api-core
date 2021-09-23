package main

import (
	"fmt"
	"log"

	"github.com/yoshinori-development/simple-community-api-main/config"
	"github.com/yoshinori-development/simple-community-api-main/models"
	"github.com/yoshinori-development/simple-community-api-main/repositories"
)

func seeder() error {
	announcementRepository := repositories.NewAnnouncementRepository()
	for i := 0; i < 10; i++ {
		announcementInput := models.AnnouncementRepositoryCreateInput{
			Announcement: models.Announcement{
				Title:   fmt.Sprintf("title%d", i),
				Content: fmt.Sprintf("content%d", i),
			},
		}
		announcementRepository.Create(announcementInput)
	}
	return nil
}

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = repositories.InitDbCore()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer repositories.Close()

	seeder()
}
