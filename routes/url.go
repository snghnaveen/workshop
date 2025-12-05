package routes

import (
	"github.com/snghnaveen/workshop/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.Engine) {
	r.POST("/api/shorten", controllers.LongURL)
	r.GET("/r/:shortCode", controllers.ShortURL)
}
