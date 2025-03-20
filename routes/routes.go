package routes

import (
	"github.com/gin-gonic/gin"
	"music-store/handlers"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "music-store/docs" // нужна для работы Swagger
)

func SetupRoutes(handler *handlers.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/songs", handler.GetSongs)
	r.POST("/songs", handler.AddSong)
	r.DELETE("/songs/:id", handler.DeleteSong)
	r.PATCH("/songs/:id", handler.UpdateSong)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
