package handlers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dkosasih/meeting-light-proxy/models"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

const openhabUrl = "https://192.168.0.185:8443"

func UpdateOpenHab(c *gin.Context) {
	var proxyModel models.MeetingStatus
	if err := c.BindJSON(&proxyModel); err != nil {
		c.Error(fmt.Errorf("error parsing meeting status model: %v", err))
	}
	data, _ := json.Marshal(proxyModel)
	fmt.Printf("model: %s\n\r", string(data))

	chann := make(chan error)

	go sendCommand(proxyModel, chann)

	if err := <-chann; err != nil {
		c.Error(fmt.Errorf("error making request to OpenHAB server: %v", err))
	}
}

func sendCommand(proxyModel models.MeetingStatus, c chan error) {
	json_data, marshalerr := json.Marshal(proxyModel)
	if marshalerr != nil {
		log.Fatal(marshalerr)
		c <- marshalerr
		return
	}

	req, err := http.NewRequest("POST", openhabUrl+"/rest/items/MeetingStatus", bytes.NewBuffer((json_data)))
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Authorization", "basic ZGtvc2FzaWg6ZEtob21lMTIzNA==")
	if err != nil {
		c <- err
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
		c <- err
		return
	}

	c <- nil
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
