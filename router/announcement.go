package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-main/services"
)

type AnnouncementHandler struct {
	AnnouncementService services.AnnouncementService
}

type NewAnnouncementHandlerInput struct {
	AnnouncementService services.AnnouncementService
}

func NewAnnouncementHandler(input NewAnnouncementHandlerInput) *AnnouncementHandler {
	return &AnnouncementHandler{
		AnnouncementService: input.AnnouncementService,
	}
}

type AnnouncementResponse struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updateAt"`
}

type AnnouncementListResponse []AnnouncementResponse

func (controller *AnnouncementHandler) List(c *gin.Context) {
	announcements, err := controller.AnnouncementService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var response AnnouncementListResponse
	for _, v := range announcements {
		announceResponse := AnnouncementResponse{
			Title:     v.Title,
			Content:   v.Content,
			UpdatedAt: v.UpdatedAt,
		}
		response = append(response, announceResponse)
	}
	c.JSON(http.StatusOK, response)
}
