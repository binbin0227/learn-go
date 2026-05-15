package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserInfo struct{
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("login",func(ctx *gin.Context) {
		var u UserInfo
		err := ctx.ShouldBind(&u)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			ctx.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})
	r.Run(":10086")
	// http://127.0.0.1:10086/login?username=Bin&password=123456
}