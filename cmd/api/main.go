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

	// tokenStringRaw := "eyJ0eXAiOiJKV1QiLCJraWQiOiJkZjhlMjRjOC0xODk0LTRmODgtYjBhNS0yMDc4ODIzMDU4YzciLCJhbGciOiJFUzI1NiIsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAuYXAtbm9ydGhlYXN0LTEuYW1hem9uYXdzLmNvbS9hcC1ub3J0aGVhc3QtMV9lS0NkeVRFZ1YiLCJjbGllbnQiOiI0aHZqdmpvYWQ1aXQyMzVndGZuMnFwNXFjZSIsInNpZ25lciI6ImFybjphd3M6ZWxhc3RpY2xvYWRiYWxhbmNpbmc6YXAtbm9ydGhlYXN0LTE6ODU4ODg0MTk4MDQ0OmxvYWRiYWxhbmNlci9hcHAvc2ltcGxlLWNvbW11bml0eS1kZXZlbG9wLWNvbW1vbi83MWMzZTljOTNjMjcyMjFkIiwiZXhwIjoxNjMzMjI3OTE0fQ=.eyJzdWIiOiIzZDZmYTkzZS04YzYxLTQ2ZDAtYjlhZC1hNDgxNzY4YzYxMzEiLCJlbWFpbF92ZXJpZmllZCI6InRydWUiLCJlbWFpbCI6Inlvc2hpbm9yaS5zYXRvaC50b2t5b0BnbWFpbC5jb20iLCJ1c2VybmFtZSI6Inlvc2hpbm9yaSIsImV4cCI6MTYzMzIyNzkxNCwiaXNzIjoiaHR0cHM6Ly9jb2duaXRvLWlkcC5hcC1ub3J0aGVhc3QtMS5hbWF6b25hd3MuY29tL2FwLW5vcnRoZWFzdC0xX2VLQ2R5VEVnViJ9.rc3wbMVRI-B6FdWy65Yf552uN7-aRaMzf7Q94EnsvmqS9NdDXC9NC3PGu5UPj-qgaGWVKf4oxO05eSCIzJdXjA=="
	// tokenString := strings.Replace(tokenStringRaw, "=", "", -1)
	// parts := strings.Split(tokenString, ".")
	// if len(parts) != 3 {
	// 	fmt.Println("fffffffffffffffff")
	// }
	// fmt.Println(parts)

	// dec, err := base64.RawURLEncoding.DecodeString(parts[0])
	// fmt.Println(parts[0])
	// if err != nil {
	// 	fmt.Println("ggggggggggggggg")
	// 	log.Println(err)
	// }
	// fmt.Println(dec)

	// token, err := jwt.Parse(tokenString, func(tk *jwt.Token) (interface{}, error) {
	// 	fmt.Println("bbbbbbbbbbb")
	// 	url := fmt.Sprintf("https://public-keys.auth.elb.%s.amazonaws.com/%s", "ap-northeast-1", "df8e24c8-1894-4f88-b0a5-2078823058c7")
	// 	fmt.Println(url)
	// 	resp, err := http.Get(url)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	defer resp.Body.Close()
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	fmt.Println(body)
	// 	if err != nil {
	// 		fmt.Println("ddddddddddddddddd")
	// 		log.Print(err)
	// 	}
	// 	return jwt.ParseECPublicKeyFromPEM(body)
	// })

	// fmt.Println("token")
	// fmt.Println(token)
	// if err != nil {
	// 	log.Println(err)
	// }

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
