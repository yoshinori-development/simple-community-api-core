package router

import (
	"net/http"
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
	root := r.Group("/api")
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
	profile := r.Group("/profile")
	profile.Use(mustAuthenticated())
	{
		profile.GET("", profileHandler.Get)
		profile.PUT("", profileHandler.CreateOrUpdate)
	}

	return r, nil
}
