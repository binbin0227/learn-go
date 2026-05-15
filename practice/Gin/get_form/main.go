package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html","index.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	// 处理login请求
	r.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username") // 也可使用 DefaultPostForm, GetPostForm
		password := ctx.PostForm("password")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":10086")
}
