package main

import "github.com/gin-gonic/gin"
import "net/http"
import "time"


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post.tmpl", gin.H{
			"title": "My first post",
            "date": time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	router.Run(":8080")
}
