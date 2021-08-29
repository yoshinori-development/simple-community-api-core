package main

import (
	"fmt"
	"log"

	"github.com/yoshinori-development/simple-community-api-core/config"
	"github.com/yoshinori-development/simple-community-api-core/drivers/db_core"
	"github.com/yoshinori-development/simple-community-api-core/models"
	"gorm.io/gorm"
)

func seeder(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		announcement := models.Announcement{
			Title:   fmt.Sprintf("title%d", i),
			Content: fmt.Sprintf("content%d", i),
		}
		if err := db.Create(&announcement).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func main() {
	config, err := config.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := db_core.Open(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db_core.Close(db)

	seeder(db)
}
