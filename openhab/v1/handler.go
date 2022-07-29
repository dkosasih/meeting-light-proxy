package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type openhabHandler struct {
	HttpClient *http.Client
}

const OPENHAB_URL_ENV_KEY = "OPENHAB_URL"
const OPENHAB_PASS_ENV_KEY = "OPENHAB_PASS"

func NewOpenHabHandler(httpClient *http.Client) *openhabHandler {
	return &openhabHandler{httpClient}
}

func (h *openhabHandler) UpdateOpenHab(c *gin.Context) {
	var proxyModel *MeetingStatus
	if err := c.BindJSON(&proxyModel); err != nil {
		c.Error(fmt.Errorf("error parsing meeting status model: %v", err))
		return
	}
	data, _ := json.Marshal(proxyModel)
	fmt.Printf("model: %s\n\r", string(data))

	chann := make(chan error)

	go func() {
		defer close(chann)
		sendCommand(proxyModel, h.HttpClient, chann)
	}()

	if err := <-chann; err != nil {
		c.Error(fmt.Errorf("error making request to OpenHAB server: %v", err))
		return
	}
}

func sendCommand(proxyModel *MeetingStatus, client *http.Client, c chan<- error) {
	if proxyModel == nil {
		c <- fmt.Errorf("Nil proxy model is not allowed")
		return
	}

	json_data, marshalerr := json.Marshal(proxyModel)
	if marshalerr != nil {
		log.Println(marshalerr)
		c <- marshalerr
		return
	}

	req, err := http.NewRequest("POST", os.Getenv(OPENHAB_URL_ENV_KEY)+"/rest/items/MeetingStatus", bytes.NewBuffer((json_data)))
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Authorization", "basic "+os.Getenv(OPENHAB_PASS_ENV_KEY))
	if err != nil {
		c <- err
		return
	}

	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
		c <- err
		return
	}

	c <- nil
}
