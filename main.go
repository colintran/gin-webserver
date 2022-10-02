package main

import (
	"github.com/gin-gonic/gin"
	"cuong-go-ws/handlers"
)

func main(){
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumsById)
	router.POST("/albums", handlers.PostAlbums)
	router.Run("localhost:8888")
}