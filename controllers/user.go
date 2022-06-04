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

func (controller *UserController) GetAll(c *gin.Context) {
	db := controller.db

	var users []models.User

	db.Model(&models.User{}).Find(&users)

	c.JSON(http.StatusOK, users)
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{db: db}
}
