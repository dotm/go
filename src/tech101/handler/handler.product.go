package handler

import (
	"net/http"
	"strconv"

	"tech101/model"
	"tech101/resource/product"

	"github.com/gin-gonic/gin"
)

func GetProductHandler(c *gin.Context) {

	var (
		productID int64
		p         model.Product
		err       error
	)
	ProductIDStr := c.DefaultQuery("id", "")

	productID, err = strconv.ParseInt(ProductIDStr, 10, 64)
	if err != nil {
		respBad := map[string]interface{}{
			"error": "parsing product id error",
		}

		c.JSON(http.StatusBadRequest, respBad)
		return
	}

	p, err = product.Get(productID)
	if err != nil {
		respBad := map[string]interface{}{
			"error": "parsing product error",
		}

		c.JSON(http.StatusBadRequest, respBad)
		return
	}

	resp := map[string]interface{}{
		"data": p,
	}

	c.JSON(http.StatusOK, resp)
	return
}

func GetProductListHandler(c *gin.Context) {
	var (
		err    error
		limit  int64
		offset int64
	)

	resp := map[string]interface{}{}
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err = strconv.ParseInt(limitStr, 10, 64)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		return
	}

	products, err := product.GetAll(limit, offset)
	if err != nil {
		return
	}

	resp["data"] = products
	resp["total_data"] = len(products)
	c.JSON(http.StatusOK, resp)
}

func CreateProductHandler(c *gin.Context) {

	resp := map[string]interface{}{}
	name := c.PostForm("name")
	description := c.PostForm("description")

	p, err := product.Create(name, description)
	if err != nil {
		resp["error"] = err.Error()
		c.JSON(http.StatusBadRequest, resp)
	}

	resp["status"] = 1
	resp["data"] = p
	c.JSON(http.StatusOK, resp)
	return
}
