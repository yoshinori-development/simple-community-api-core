package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-core/domain/model"
	"github.com/yoshinori-development/simple-community-api-core/usecase"
	"gorm.io/gorm"
)

type ProfileHandler struct {
	ProfileUsecase usecase.ProfileUsecase
}

type NewProfileHandlerInput struct {
	ProfileUsecase usecase.ProfileUsecase
}

func NewProfileHandler(input NewProfileHandlerInput) *ProfileHandler {
	return &ProfileHandler{
		ProfileUsecase: input.ProfileUsecase,
	}
}

type ProfileResponse struct {
	Nickname  string    `json:"nickname"`
	Age       uint      `json:"age"`
	Birthdate string    `json:"birthdate"`
	UpdatedAt time.Time `json:"updateAt"`
}

func (handler *ProfileHandler) Get(c *gin.Context) {
	sub, subExists := c.Get("sub")
	if subExists {
		c.Status(http.StatusUnauthorized)
	}
	profile, err := handler.ProfileUsecase.Get(usecase.ProfileUsecaseGetInput{
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
		Birthdate: profile.Birthdate,
		UpdatedAt: profile.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

type ProfileHandlerCreateInput struct {
	Nickname  string `form:"nickname" binding:"required,min=5"`
	Age       uint   `form:"age" validate:"numeric"`
	Birthdate string `form:"birthdate" validate:"string"`
}

func (handler *ProfileHandler) CreateOrUpdate(c *gin.Context) {
	var input ProfileHandlerCreateInput
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, subExists := c.Get("sub")
	if !subExists {
		c.Status(http.StatusUnauthorized)
		return
	}
	err := handler.ProfileUsecase.CreateOrUpdate(usecase.ProfileUsecaseCreateOrUpdateInput{
		Profile: model.Profile{
			Sub:       sub.(string),
			Nickname:  input.Nickname,
			Age:       input.Age,
			Birthdate: input.Birthdate,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}
