package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-core/usecase"
)

type AnnouncementHandler struct {
	AnnouncementUsecase usecase.AnnouncementUsecase
}

type NewAnnouncementHandlerInput struct {
	AnnouncementUsecase usecase.AnnouncementUsecase
}

func NewAnnouncementHandler(input NewAnnouncementHandlerInput) *AnnouncementHandler {
	return &AnnouncementHandler{
		AnnouncementUsecase: input.AnnouncementUsecase,
	}
}

type AnnouncementResponse struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updateAt"`
}

type AnnouncementListResponse []AnnouncementResponse

func (handler *AnnouncementHandler) List(c *gin.Context) {
	announcements, err := handler.AnnouncementUsecase.List()
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
