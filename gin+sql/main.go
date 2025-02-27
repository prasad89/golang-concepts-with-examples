package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func pongID(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"message": "pong with " + id})
}

func main() {
	r := gin.Default()

	r.GET("/ping", pong)
	r.POST("/ping/:id", pongID)

	r.Run()
}
