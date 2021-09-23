package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-main/config"
)

func setupCors(r *gin.Engine, routerConf config.Router) {
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
