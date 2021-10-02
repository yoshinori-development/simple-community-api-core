package router

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/yoshinori-development/simple-community-api-main/config"
)

type xAmznOidc struct {
	AccessToken string `header:"X-Amzn-Oidc-Accesstoken"`
	Identity    string `header:"X-Amzn-Oidc-Identity"`
	Data        string `header:"X-Amzn-Oidc-Data"`
}

type xAmznOidcDataHeader struct {
	Alg    string `header:"alg"`
	Kid    string `header:"kid"`
	Signer string `header:"signer"`
	Iss    string `header:"iss"`
	Client string `header:"client"`
	Exp    string `header:"exp"`
}

type xAmznOidcDataPayload struct {
	Sub   string `header:"sub"`
	Email string `header:"email"`
}

func setupAuthenticate(r *gin.Engine, awsConf config.Aws) {
	r.Use(authenticate(awsConf))
}

func authenticate(awsConf config.Aws) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := xAmznOidc{}
		if err := c.ShouldBindHeader(&h); err != nil {
			log.Print(err)
		}

		fmt.Println("111111111111")

		log.Print(h)
		if h.Data != "" {
			fmt.Println(h.Data)
			token, err := jwt.Parse("eyJ0eXAiOiJKV1QiLCJraWQiOiJkZjhlMjRjOC0xODk0LTRmODgtYjBhNS0yMDc4ODIzMDU4YzciLCJhbGciOiJFUzI1NiIsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAuYXAtbm9ydGhlYXN0LTEuYW1hem9uYXdzLmNvbS9hcC1ub3J0aGVhc3QtMV9lS0NkeVRFZ1YiLCJjbGllbnQiOiI0aHZqdmpvYWQ1aXQyMzVndGZuMnFwNXFjZSIsInNpZ25lciI6ImFybjphd3M6ZWxhc3RpY2xvYWRiYWxhbmNpbmc6YXAtbm9ydGhlYXN0LTE6ODU4ODg0MTk4MDQ0OmxvYWRiYWxhbmNlci9hcHAvc2ltcGxlLWNvbW11bml0eS1kZXZlbG9wLWNvbW1vbi83MWMzZTljOTNjMjcyMjFkIiwiZXhwIjoxNjMzMTkyOTk5fQ==.eyJzdWIiOiIzZDZmYTkzZS04YzYxLTQ2ZDAtYjlhZC1hNDgxNzY4YzYxMzEiLCJlbWFpbF92ZXJpZmllZCI6InRydWUiLCJlbWFpbCI6Inlvc2hpbm9yaS5zYXRvaC50b2t5b0BnbWFpbC5jb20iLCJ1c2VybmFtZSI6Inlvc2hpbm9yaSIsImV4cCI6MTYzMzE5Mjk5OSwiaXNzIjoiaHR0cHM6Ly9jb2duaXRvLWlkcC5hcC1ub3J0aGVhc3QtMS5hbWF6b25hd3MuY29tL2FwLW5vcnRoZWFzdC0xX2VLQ2R5VEVnViJ9.eS86Jmql7pfoJxOYSkzT7OSkC677vi6-PptGtUeiQZzXq9Hmv7RA2n53EPIR7uRto1N97JvZl_6vuJpJeNbE4A==", func(tk *jwt.Token) (interface{}, error) {
				fmt.Println(tk)
				fmt.Println("aaaaaaaaaaaaaa")
				if _, ok := tk.Method.(*jwt.SigningMethodRSA); !ok {
					fmt.Println("ccccccccccccccccc")
					log.Printf("Unexpected signing method: %v", tk.Header["alg"])
					c.Status(http.StatusUnauthorized)
				}
				fmt.Println("bbbbbbbbbbb")
				url := fmt.Sprintf("https://public-keys.auth.elb.%s.amazonaws.com/%s", awsConf.DefaultRegion, tk.Header["kid"])
				fmt.Println(url)
				publicKey, err := ioutil.ReadFile(url)
				fmt.Println(publicKey)
				if err != nil {
					fmt.Println("ddddddddddddddddd")
					log.Print(err)
					c.Status(http.StatusUnauthorized)
				}
				return publicKey, nil
			})

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("222222222222")
			fmt.Println(token)
			fmt.Println(token.Claims)
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				fmt.Println("AAAAAAAAAAAa")
				c.Set("sub", claims["sub"])
			} else {
				fmt.Println("BBBBBBBBB")
				fmt.Println(err)
			}
		}

		c.Next()
	}
}

func mustAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		if sub, exists := c.Get("sub"); !exists {
			fmt.Println(sub)
			c.JSON(http.StatusUnauthorized, RenderMessageError(errors.New("not authenticated"), "ログインが必要です"))
			c.Abort()
		}
		c.Next()
	}
}
