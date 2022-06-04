package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(router *gin.RouterGroup, db *gorm.DB) {
	SetUserRoutes(router, db)
}
