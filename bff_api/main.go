package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("Hello BFF")
	handler := ApiHandler{
		ApiOneUrl: "http://localhost:5000",
		ApiTwoUrl: "http://localhost:6000",
	}
	GetMainEngine(handler).Run(":8082")
}

func GetMainEngine(handler ApiHandler) * gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/apis", handler.Apis)
		//v1.POST("/apis/all_wait", handler.AllWait)
		//v1.POST("/apis/quick_response", handler.QuickResponse)
	}
	return r
}

type ApiHandler struct {
	ApiOneUrl string
	ApiTwoUrl string
}

func(h *ApiHandler) Apis(c *gin.Context) {
	//var apiResponse []map[string]interface{}

	// port 5000にfetch
	// port 6000にfetch
	// merge
	//http.Geth.ApiOneUrl


	c.JSON(http.StatusOK, gin.H{"hoge": "fuga"})
	return
}

func (h *ApiHandler) GetApiOneInfo() map[string]interface{} {
	m := make(map[string]interface{})

	return m
}


func (h *ApiHandler) GetApiTwoInfo() map[string]interface{} {
	m := make(map[string]interface{})

	return m
}