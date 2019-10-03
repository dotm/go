package main

import (
	"fmt"
	"net/http"
	"tech101/db"
	"tech101/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	err := db.Init("postgres://gzpkvgqxjddpka:3fa8aff11204dbea316573579e0e089d41c9be0964721f58fe4c7dde95f6391d@ec2-54-83-9-169.compute-1.amazonaws.com:5432/d6p5komkb8t14e")
	if err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()
	// /v1/api/product?id=1
	router.GET("/v1/api/product", handler.GetProductHandler)
	router.GET("/v1/api/product/list", handler.GetProductListHandler)
	router.POST("/v1/api/product/create", handler.CreateProductHandler)
	// /v1/api/user/1
	router.GET("/v1/api/user/:user_id", handler.GetUserHandler)
	router.POST("/v1/api/user/create", handler.CreateUserHandler)

	router.GET("/v1/api/poll/:poll_id", handler.GetPollFromGC)

	//  /v1/api/tome/bulk?ids=1,2,3
	router.GET("/v1/api/tome/bulk", handler.GetBulkProductFromTome)

	var chanExample chan int
	router.GET("v1/test", func(c *gin.Context) {
		chanExample = make(chan int, 0)
		resp := map[string]interface{}{}
		fmt.Println("test chan")
		value := <-chanExample
		fmt.Println("value:", value)
		resp["ok"] = 1
		close(chanExample)
		c.JSON(http.StatusOK, resp)
	})
	router.GET("v1/test2", func(c *gin.Context) {
		fmt.Println("test chan2")
		chanExample <- 10
	})
	router.Run(":8080")

}
