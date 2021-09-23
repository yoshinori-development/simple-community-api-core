package router

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-main/models"
	"github.com/yoshinori-development/simple-community-api-main/services"
	"gorm.io/gorm"
)

type ProfileHandler struct {
	ProfileService services.ProfileService
}

type NewProfileHandlerInput struct {
	ProfileService services.ProfileService
}

func NewProfileHandler(input NewProfileHandlerInput) *ProfileHandler {
	return &ProfileHandler{
		ProfileService: input.ProfileService,
	}
}

type ProfileResponse struct {
	Nickname  string    `json:"nickname"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updateAt"`
}

func (controller *ProfileHandler) Get(c *gin.Context) {
	sub, subExists := c.Get("sub")
	if subExists {
		c.Status(http.StatusUnauthorized)
	}
	profile, err := controller.ProfileService.Get(services.ProfileServiceGetInput{
		Sub: sub.(string),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	response := ProfileResponse{
		Nickname:  profile.Nickname,
		Age:       profile.Age,
		UpdatedAt: profile.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

type ProfileHandlerCreateOrUpdateInput struct {
	Nickname string `form:"nickname" binding:"required,min=5"`
	Age      uint   `form:"age" binding:"numeric,max=150"`
}

func (controller *ProfileHandler) CreateOrUpdate(c *gin.Context) {
	var input ProfileHandlerCreateOrUpdateInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": formatValidationErrors(err)})
		return
	}

	sub, subExists := c.Get("sub")
	if !subExists {
		c.Status(http.StatusUnauthorized)
		return
	}
	err := controller.ProfileService.CreateOrUpdate(services.ProfileServiceCreateOrUpdateInput{
		Profile: models.Profile{
			Sub:      sub.(string),
			Nickname: input.Nickname,
			Age:      input.Age,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}
