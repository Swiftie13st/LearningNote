package main

import "github.com/gin-gonic/gin"
import (
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello World!")
    })

    r.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
    })

    r.Run(":9003") // listen and serve on 0.0.0.0:PORT(default:8080)
}
