package controllers

import (
	"net/http"
	"simci-go/core/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

// @BasePath /api

// GetAll godoc
// @Summary Get all users
// @Schemes
// @Description Get all users
// @Produce json
// @Success 200 {string} Helloworld
// @Router /users [get]
func (controller *UserController) GetAll(c *gin.Context) {
	db := controller.db

	var users []models.User

	db.Model(&models.User{}).Find(&users)

	c.JSON(http.StatusOK, users)
}

// @BasePath /api

// Create godoc
// @Summary Create new user
// @Param	user body models.User true "User data"
// @Description Create new user
// @Produce json
// @Success 200 {string} success
// @Router /users [POST]
func (controller *UserController) Create(c *gin.Context) {
	db := controller.db

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Malformed data",
		})
		return
	}

	err := user.HashPassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal error",
		})
		return
	}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{db: db}
}
