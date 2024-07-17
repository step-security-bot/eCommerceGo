package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // Define a GET endpoint
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  // Start the server
  r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
