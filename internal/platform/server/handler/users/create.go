package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	domain "simci-go/internal"
)

type createUserRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @BasePath /api

// Create godoc
// @Summary Create new user
// @Param user body createUserRequest true "User data"
// @Description Create new user
// @Produce json
// @Success 200 {string} success
// @Router /users [POST]
func Create(repository domain.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		var user createUserRequest

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Malformed data",
			})
			return
		}

		newUser, err := domain.NewUser(user.ID, user.Name, user.Email, user.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Malformed payload",
			})
			return
		}

		err = repository.Save(c, newUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal error",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "success",
		})
	}
}
