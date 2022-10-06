package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HandlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name")

	if !ok {
		name = "default"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件m1：统计请求处理函数的执行时间
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")

	start := time.Now() //计时
	// go funcXX(c.Copy()) // 在funcXX中只能使用c的拷贝
	c.Next() //调用后续的处理函数
	// c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)

	fmt.Printf("cost : %v\n", cost)
	fmt.Println("m1 out ...")
}

// 定义一个中间件m2
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")

	c.Set("name", "Bruce")

	c.Next() //调用后续的处理函数
	// c.Abort() // 阻止调用后续的处理函数
	fmt.Println("m2 out ...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 链接数据库
	// 等准备工作
	return func(c *gin.Context) {
		// 存放具体的逻辑
		if doCheck {
			// 如：是否登录的判断
			// if 登录
			// c.Next()
			// else
			// c.Abort()
		} else {
			c.Next()
		}

	}
}

func main() {
	r := gin.Default()

	r.Use(m1, m2, authMiddleware(true)) // 全局注册中间件函数m1, m2

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	r.GET("/index", indexHandler)

	r.GET("/user", func(c *gin.Context) {
		fmt.Println("user")
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	// 为路由组注册中间件方法
	// videoGroup := r.Group("/video", m1)
	videoGroup := r.Group("/video")
	videoGroup.Use(m1)
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
