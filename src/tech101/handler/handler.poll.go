package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"tech101/model"
)

func GetPollFromGC(c *gin.Context) {

	var poll model.PollResponse
	data := map[string]interface{}{}

	req, err := http.NewRequest("GET", "https://chat.tokopedia.com/gmf/api/v1/poll/7415", nil)
	if err != nil {
		data["error"] = err.Error()
		c.JSON(http.StatusInternalServerError, data)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		data["error"] = err.Error()
		c.JSON(http.StatusInternalServerError, data)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&poll); err != nil {
		log.Println(err)
	}

	data["status"] = 1
	data["data"] = poll
	c.JSON(http.StatusOK, data)

	return
}
