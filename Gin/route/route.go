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

	r.Any("/any", func(c *gin.Context) {

		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case "POST":
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"method": "Others",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"method": "Any",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "NotFound",
		})
	})

	// 路由组，把共用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/about", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/video/about",
			})
		})
		videoGroup.GET("/home", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "/video/home",
			})
		})
	}

	r.Run(":9003") // listen and serve on 0.0.0.0:PORT(default:8080)

}
