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
			token, err := jwt.Parse(h.Data, func(token *jwt.Token) (interface{}, error) {
				fmt.Println(token)
				fmt.Println("aaaaaaaaaaaaaa")
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					fmt.Println("ccccccccccccccccc")
					log.Printf("Unexpected signing method: %v", token.Header["alg"])
					c.Status(http.StatusUnauthorized)
				}
				fmt.Println("bbbbbbbbbbb")
				url := fmt.Sprintf("https://public-keys.auth.elb.%s.amazonaws.com/%s", awsConf.DefaultRegion, token.Header["kid"])
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
