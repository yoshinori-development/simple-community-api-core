package main

import (
	"log"

	"github.com/yoshinori-development/simple-community-api-main/config"
	"github.com/yoshinori-development/simple-community-api-main/i18n"
	"github.com/yoshinori-development/simple-community-api-main/repositories"
	"github.com/yoshinori-development/simple-community-api-main/router"
)

func main() {
	var err error

	err = config.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = repositories.InitDbMain()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer repositories.Close()

	err = i18n.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	r, err := router.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	r.Run()
}
