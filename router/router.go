package router

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
	"github.com/yoshinori-development/simple-community-api-main/config"
	"github.com/yoshinori-development/simple-community-api-main/repositories"
	"github.com/yoshinori-development/simple-community-api-main/services"
)

var trans ut.Translator

func Init() (*gin.Engine, error) {
	config := config.Get()
	routerConf := config.Router
	awsConf := config.Aws

	r := gin.Default()
	setupCors(r, routerConf)
	setupAuthenticate(r, awsConf)

	uni := ut.New(ja.New())
	trans, _ = uni.GetTranslator("ja")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = ja_translations.RegisterDefaultTranslations(v, trans)
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	r.GET("/health", func(c *gin.Context) {
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
	r.GET("/announcements", announcementHandler.List)

	profileHandler := NewProfileHandler(NewProfileHandlerInput{
		ProfileService: profileService,
	})
	r.GET("/profile", profileHandler.Get)
	r.PUT("/profile", profileHandler.CreateOrUpdate)

	return r, nil
}

type validationError map[string]string

func formatValidationErrors(err error) validationError {
	errs := err.(validator.ValidationErrors)
	formattedErrs := make(validationError, len(errs))
	for _, err := range errs {
		formattedErrs[err.Field()] = err.Translate(trans)
	}
	fmt.Println(formattedErrs)
	return formattedErrs
}
