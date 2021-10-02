package router

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
	"github.com/yoshinori-development/simple-community-api-main/config"
	"github.com/yoshinori-development/simple-community-api-main/i18n"
	"github.com/yoshinori-development/simple-community-api-main/repositories"
	"github.com/yoshinori-development/simple-community-api-main/services"
)

func Init() (*gin.Engine, error) {
	config := config.Get()
	routerConf := config.Router
	awsConf := config.Aws

	r := gin.Default()
	setupCors(r, routerConf)
	setupAuthenticate(r, awsConf)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = ja_translations.RegisterDefaultTranslations(v, i18n.Translator)
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	root := r.Group("/api")
	root.Use(RequestLogger())
	root.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	announcementRepository := repositories.NewAnnouncementRepository()
	profileRepository := repositories.NewProfileRepository()

	announcementService := services.NewAnnouncementService(services.NewAnnouncementServiceInput{
		AnnouncementRepository: announcementRepository,
	})
	profileService := services.NewProfileService(services.NewProfileServiceInput{
		ProfileRepository: profileRepository,
	})

	announcementHandler := NewAnnouncementHandler(NewAnnouncementHandlerInput{
		AnnouncementService: announcementService,
	})
	root.GET("/announcements", announcementHandler.List)

	profileHandler := NewProfileHandler(NewProfileHandlerInput{
		ProfileService: profileService,
	})
	profile := root.Group("/profile")
	profile.Use(mustAuthenticated())
	{
		profile.GET("", profileHandler.Get)
		profile.PUT("", profileHandler.CreateOrUpdate)
	}

	return r, nil
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Host, c.Request.RemoteAddr, c.Request.RequestURI)

		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))

		c.Next()
	}
}
