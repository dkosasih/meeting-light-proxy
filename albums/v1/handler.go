package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type albumHandler struct {
}

func NewAlbumHandler() *albumHandler {
	return &albumHandler{}
}

// albumsCollections slice to seed record album data.
var albumsCollections = []Album{
	{ID: "1", Title: "Blue Train v1", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru v1", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown v1", Artist: "Sarah Vaughan", Price: 39.99},
}

// postAlbums adds an album from JSON received in the request body.
func (ah *albumHandler) PostAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albumsCollections = append(albumsCollections, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (ah *albumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albumsCollections {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// getAlbums responds with the list of all albums as JSON.
func (ah *albumHandler) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumsCollections)
}
