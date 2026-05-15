package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// http://127.0.0.1:10086/index
	// 跳转B站
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "https://www.bilibili.com")
	})

	// http://127.0.0.1:10086/a
	// 跳转到 /b 对应的路由处理函数
	r.GET("/a", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/b" // 修改请求的uri
		r.HandleContext(ctx) //继续后续的处理
	})
	r.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	r.Run(":10086")
}
