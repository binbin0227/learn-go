package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 方法一：map
	r.GET("/json", func(c *gin.Context) {
		data := gin.H{"name": "Bin", "message": "你好", "age": 20}
		c.JSON(http.StatusOK, data)
	})

	// 方法二：结构体
	type msg struct {
		Name    string
		Message string
		Age     int
	}
	r.GET("/json1", func(c *gin.Context) {
		data1 := msg{
			"哈基萍",
			"你好",
			21,
		}
		c.JSON(http.StatusOK, data1)
	})

	r.Run(":10086")
}
