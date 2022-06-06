package routers

import (
	"simci-go/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRouter := router.Group("/users")
	controllerUser := controllers.NewUserController(db)

	userRouter.GET("", controllerUser.GetAll)
	userRouter.POST("", controllerUser.Create)
}
