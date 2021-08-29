package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshinori-development/simple-community-api-core/models"
	"gorm.io/gorm"
)

type ProfileController struct {
	profileModel *models.ProfileModel
}

type NewProfileControllerInput struct {
	DB *gorm.DB
}

func NewProfileController(input NewProfileControllerInput) *ProfileController {
	return &ProfileController{
		profileModel: models.NewProfileModel(models.NewProfileModelInput{
			DB: input.DB,
		}),
	}
}

func (controller *ProfileController) Get(c *gin.Context) {
	sub, subExists := c.Get("sub")
	if subExists {
		c.Status(http.StatusUnauthorized)
	}
	profile, err := controller.profileModel.Get(models.ProfileModelGetInput{
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
	c.JSON(http.StatusOK, profile)
}

type ProfileControllerCreateInput struct {
	Nickname  string `form:"nickname"`
	Age       uint   `form:"age"`
	Birthdate string `form:"birthdate"`
}

func (controller *ProfileController) Create(c *gin.Context) {
	var input ProfileControllerCreateInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sub, subExists := c.Get("sub")
	if !subExists {
		c.Status(http.StatusUnauthorized)
		return
	}
	err := controller.profileModel.Create(models.ProfileModelCreateInput{
		Sub:       sub.(string),
		Nickname:  input.Nickname,
		Age:       input.Age,
		Birthdate: input.Birthdate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}
