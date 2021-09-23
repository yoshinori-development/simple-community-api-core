package main

import (
	"log"

	"github.com/yoshinori-development/simple-community-api-main/config"
	"github.com/yoshinori-development/simple-community-api-main/repositories"
	"github.com/yoshinori-development/simple-community-api-main/router"
)

// _ "github.com/go-sql-driver/mysql"

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

	r, err := router.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Run()
}
