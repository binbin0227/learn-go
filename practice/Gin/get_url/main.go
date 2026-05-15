package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:str1/:str2", func(ctx *gin.Context) {
		year := ctx.Param("str1")
		month := ctx.Param("str2")
		ctx.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	r.Run(":10086")
	// 访问 http://localhost:10086/2006/02
}
