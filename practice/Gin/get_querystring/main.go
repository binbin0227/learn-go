package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {

		// name := c.Query("query")
		// name := c.DefaultQuery("query","somebody") // 指定默认值
		name, ok := c.GetQuery("query") // 返回ok
		if !ok {
			name = "somebody"
		}

		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":10086")
}
