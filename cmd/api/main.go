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

	// fmt.Println("bbbbbbbbbbb")
	// url := fmt.Sprintf("https://public-keys.auth.elb.%s.amazonaws.com/%s", "ap-northeast-1", "df8e24c8-1894-4f88-b0a5-2078823058c7")
	// fmt.Println(url)
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Print(err)
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(body)
	// if err != nil {
	// 	fmt.Println("ddddddddddddddddd")
	// 	log.Print(err)
	// }
	// publicKey, err := jwt.ParseECPublicKeyFromPEM(body)
	// if err != nil {
	// 	fmt.Println("eeeeeeeeeeeeeee")
	// 	log.Print(err)
	// }
	// fmt.Println("public key")
	// fmt.Println(publicKey)

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
