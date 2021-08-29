package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-core/config"
	"github.com/yoshinori-development/simple-community-api-core/controllers"
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

	announcementController := controllers.NewAnnouncementController(controllers.NewAnnouncementControllerInput{
		DB: db,
	})
	r.GET("/announcements", announcementController.List)

	profileController := controllers.NewProfileController(controllers.NewProfileControllerInput{
		DB: db,
	})
	r.GET("/profile", profileController.Get)
	r.POST("/profile", profileController.Create)

	return r, nil
}

func setupMiddlewareCors(r *gin.Engine, routerConf config.Router) {
	var allowOrigins []string
	for _, origin := range routerConf.AllowOrigins {
		allowOrigins = append(allowOrigins, origin)
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
}
