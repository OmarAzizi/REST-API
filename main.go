package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album silce data
var albums = []album{
	{ID: "1", Title: "Ride The Lightning", Artist: "Metallica", Price: 56.99},
	{ID: "2", Title: "Paranoid", Artist: "Black Sabbath", Price: 59.99},
	{ID: "3", Title: "Morningrise", Artist: "Opeth", Price: 54.99},
}

// this function responds with a list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Binding the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add newAlbum to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// locates album whose ID value matches the id
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
