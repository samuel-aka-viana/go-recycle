package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go-project/docs"
	"go-project/handler"
)

func InitializeRoutes(router *gin.Engine) {
	// initialize handler

	handler.InitializeHandler()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.BulkUpdateOpeningHandler)
		v1.PATCH("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	// initialize swaggo
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
