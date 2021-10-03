package router

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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
			tokenString := strings.Replace(h.Data, "=", "", -1)
			token, err := jwt.Parse(tokenString, func(tk *jwt.Token) (interface{}, error) {
				fmt.Println(tk)
				fmt.Println("aaaaaaaaaaaaaa")
				if _, ok := tk.Method.(*jwt.SigningMethodECDSA); !ok {
					fmt.Println("ccccccccccccccccc")
					fmt.Printf("Unexpected signing method: %v", tk.Header["alg"])
					c.Status(http.StatusUnauthorized)
				}
				fmt.Println("bbbbbbbbbbb")
				url := fmt.Sprintf("https://public-keys.auth.elb.%s.amazonaws.com/%s", awsConf.DefaultRegion, tk.Header["kid"])
				fmt.Println(url)
				resp, err := http.Get(url)
				if err != nil {
					log.Print(err)
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("ddddddddddddddddd")
					fmt.Print(err)
					c.Status(http.StatusUnauthorized)
				}
				fmt.Println(body)
				fmt.Printf("%s", body)
				publicKey, err := jwt.ParseECPublicKeyFromPEM(body)
				if err != nil {
					fmt.Println("eeeeeeeeeeeeeee")
					fmt.Print(err)
				}
				fmt.Println("public key")
				fmt.Println(publicKey)
				return publicKey, nil
			})

			if err != nil {
				fmt.Println("hhhhhhhhhhhhh")
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
