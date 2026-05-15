package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(ctx *gin.Context) {
		// 从请求中读取文件
		f, err := ctx.FormFile("f1")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			dst := path.Join("./", f.Filename)
			ctx.SaveUploadedFile(f, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":10086")
}
