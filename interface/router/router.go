package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-core/config"
	"github.com/yoshinori-development/simple-community-api-core/infrastructure/db_core"
	"github.com/yoshinori-development/simple-community-api-core/interface/handlers"
	"github.com/yoshinori-development/simple-community-api-core/usecase"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, config config.Config) (*gin.Engine, error) {
	routerConf := config.Router
	awsConf := config.Aws

	r := gin.Default()
	setupMiddlewareCors(r, routerConf)
	setupMiddlewareAuthenticate(r, awsConf)

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	announcementRepository := db_core.NewAnnouncementRepository(db)
	profileRepository := db_core.NewProfileRepository(db)

	announcementUsecase := usecase.NewAnnouncementUsecase(usecase.NewAnnouncementUsecaseInput{
		AnnouncementRepository: announcementRepository,
	})
	profileUsecase := usecase.NewProfileUsecase(usecase.NewProfileUsecaseInput{
		ProfileRepository: profileRepository,
	})

	announcementHandler := handlers.NewAnnouncementHandler(handlers.NewAnnouncementHandlerInput{
		AnnouncementUsecase: announcementUsecase,
	})
	r.GET("/announcements", announcementHandler.List)

	profileHandler := handlers.NewProfileHandler(handlers.NewProfileHandlerInput{
		ProfileUsecase: profileUsecase,
	})
	r.GET("/profile", profileHandler.Get)
	r.PUT("/profile", profileHandler.CreateOrUpdate)

	return r, nil
}
