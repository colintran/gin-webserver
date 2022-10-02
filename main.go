package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
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

func isDuplicatedId(id string) bool{
	for _, al := range Albums {
		if id == al.Id {
			return true
		}
	}
	return false
}

func postAlbums(c *gin.Context){
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Bad request"})
		return
	}
	// validation
	if isDuplicatedId(newAlbum.Id) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Bad request"})
		log.Printf("Album id [%v] is duplicated",newAlbum.Id)
		return
	}
	if len(newAlbum.Id) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Bad request"})
		log.Printf("Album id is empty")
		return
	}
	Albums = append(Albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

// /?id=...
func getAlbumsById(c *gin.Context){
	id := c.Param("id")
	log.Printf("request album id: %v",id)
	for _, al := range Albums {
		if al.Id == id {
			c.IndentedJSON(http.StatusOK, al)
			return
		}
	}
	unfoundStr := fmt.Sprintf("Album id [%v] not found", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":unfoundStr})
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8888")
}