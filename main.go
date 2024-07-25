package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
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
		out := fmt.Sprintf("calling: %s", testUrl)
		c.String(200, out+" : "+string(body))
	})
	r.GET("/mysql", func(c *gin.Context) {
		mysqlUrl := os.Getenv("MYSQL_URL")
		db, err := sql.Open("mysql", mysqlUrl)

		if err != nil {
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("error: %s", err.Error()),
			})
			return
		}

		// defer the close till after the main function has finished
		// executing
		defer db.Close()

		err = db.Ping()

		if err != nil {
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("error: %s", err.Error()),
			})
			return
		}

		c.String(200, "mysql connect success")
	})
	r.Run(":8080")
}
