package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	m "cuong-go-ws/model"
)

func GetAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, m.Albums)
}

func isDuplicatedId(id string) bool{
	for _, al := range m.Albums {
		if id == al.Id {
			return true
		}
	}
	return false
}

func PostAlbums(c *gin.Context){
	var newAlbum m.Album
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
	m.Albums = append(m.Albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

// /?id=...
func GetAlbumsById(c *gin.Context){
	id := c.Param("id")
	log.Printf("request album id: %v",id)
	for _, al := range m.Albums {
		if al.Id == id {
			c.IndentedJSON(http.StatusOK, al)
			return
		}
	}
	unfoundStr := fmt.Sprintf("Album id [%v] not found", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":unfoundStr})
}