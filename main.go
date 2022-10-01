package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	Id string `json:"id"`
	Title string `json:"title"`
}

var Albums = []album{
	{Id: "1", Title: "Gone with the wind"},
	{Id: "2", Title: "War and Peace"},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, Albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	Albums = append(Albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func main(){
	router := gin.Default()
	router.GET("/", getAlbums)
	router.POST("/", postAlbums)
	router.Run("localhost:8888")
}