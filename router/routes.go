package router

import (
	"github.com/gin-gonic/gin"
	"github.com/italopatrick/api-opportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize Handler
	handler.InitializedHandler()

	v1 := router.Group("api/v1")
	// Show opening
	{
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)

	}
}
