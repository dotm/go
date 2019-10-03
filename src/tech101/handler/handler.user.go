package handler

import (
	"net/http"
	"strconv"

	"tech101/model"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {

	var u model.User
	resp := map[string]interface{}{}

	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {

		resp["error"] = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if userID == 1 {
		u = model.User{
			Name:  "eric",
			Email: "eric.suwardi@tokopedia.com",
			Age:   20,
		}
	} else if userID == 2 {
		u = model.User{
			Name:  "riki",
			Email: "riki@tokopedia.com",
			Age:   24,
		}
	} else {
		resp["error"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp["status"] = "ok"
	resp["data"] = u
	c.JSON(http.StatusOK, resp)
	return
}

func CreateUserHandler(c *gin.Context) {

	var userInput model.User
	resp := map[string]interface{}{}

	err := c.BindJSON(&userInput)
	if err != nil {
		resp["error"] = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}
	resp["status"] = 1
	resp["data"] = c
	c.JSON(http.StatusOK, resp)
	return
}
