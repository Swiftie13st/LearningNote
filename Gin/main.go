package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	// r.GET("/test", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	// })

	r.GET("/test", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	r.Run(":9003") // listen and serve on 0.0.0.0:PORT(default:8080)

}
