package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default() // router

	// setup routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/foo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bar",
		})
	})
	r.GET("/OPTIONS", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"online": true,
		})
	})

	// run on port 80
	r.Run(":80")
}
