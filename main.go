package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong from private service",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		testUrl := os.Getenv("TEST_URL")
		resp, err := http.Get(testUrl)
		if err != nil {
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("error: %s", err.Error()),
			})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("error: %s", err.Error()),
			})
			return
		}
		c.String(200, string(body))
	})
	r.Run(":8080")
}
