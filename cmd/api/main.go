package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yoshinori-development/simple-community-api-core/config"
	"github.com/yoshinori-development/simple-community-api-core/drivers/db_core"
	"github.com/yoshinori-development/simple-community-api-core/router"
)

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

	err = db_core.Migrate(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	r, err := router.NewRouter(db, config)
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Run()
}
