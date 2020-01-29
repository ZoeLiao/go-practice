package main

import "github.com/gin-gonic/gin"
import "net/http"


func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
    // *gin.Context: get parameters
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

    router.Run(":8001") // default: 8080
}
