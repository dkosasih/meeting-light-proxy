package openhab

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const openhabUrl = "https://192.168.0.185:8443"

func (h *openhabHandler) UpdateOpenHab(c *gin.Context) {
	var proxyModel MeetingStatus
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

func sendCommand(proxyModel MeetingStatus, c chan error) {
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
