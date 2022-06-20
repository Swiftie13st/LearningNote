package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("./index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// router.POST("/upload", func(c *gin.Context) {
	// 	// 单个文件
	// 	file, err := c.FormFile("f1")
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	log.Println(file.Filename)
	// 	dst := fmt.Sprintf("F:/tmp/%s", file.Filename)
	// 	// 上传文件到指定的目录
	// 	c.SaveUploadedFile(file, dst)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	// 	})
	// })

	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("F:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	router.Run(":9003") // listen and serve on 0.0.0.0:PORT(default:8080)

}
