package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-core/models"
	"gorm.io/gorm"
)

type AnnouncementController struct {
	announcementModel *models.AnnouncementModel
}

type NewAnnouncementControllerInput struct {
	DB *gorm.DB
}

func NewAnnouncementController(input NewAnnouncementControllerInput) *AnnouncementController {
	return &AnnouncementController{
		announcementModel: models.NewAnnouncementModel(models.NewAnnouncementModelInput{
			DB: input.DB,
		}),
	}
}

func (controller *AnnouncementController) List(c *gin.Context) {
	announcements, err := controller.announcementModel.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, announcements)
}
