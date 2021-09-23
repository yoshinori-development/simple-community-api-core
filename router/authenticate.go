package router

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/yoshinori-development/simple-community-api-main/config"
)

type xAmznOidc struct {
	AccessToken string `header:"x-amzn-oidc-accesstoken"`
	Identity    string `header:"x-amzn-oidc-identity"`
	Data        string `header:"x-amzn-oidc-data"`
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

		c.Set("sub", "test-sub-3")
		if h.Data != "" {
			token, err := jwt.Parse(h.Data, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					log.Printf("Unexpected signing method: %v", token.Header["alg"])
					c.Status(http.StatusUnauthorized)
				}
				url := fmt.Sprintf("https://public-keys.auth.elb.%s.amazonaws.com/%s", awsConf.DefaultRegion, token.Header["kid"])
				publicKey, err := ioutil.ReadFile(url)
				if err != nil {
					log.Print(err)
					c.Status(http.StatusUnauthorized)
				}
				return publicKey, nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("sub", claims["sub"])
			} else {
				fmt.Println(err)
			}
		}

		c.Next()
	}
}
