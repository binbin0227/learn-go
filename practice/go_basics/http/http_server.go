package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. 注册路由（菜单）
	// 当有人访问 "/hello" 这个路径时，交给后面这个函数处理
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// r (Request): 包含了客户端发来的所有信息（他是谁？他想要啥？）
		fmt.Printf("收到请求：方法=%s, 路径=%s\n", r.Method, r.URL.Path)

		// w (ResponseWriter): 你用来给客户端写回信的笔
		w.Write([]byte("你好！欢迎来到 Go 的 HTTP 世界！"))
	})

	// 2. 启动服务器并监听 8080 端口
	// 第二个参数 nil 代表使用默认的路由器
	fmt.Println("🚀 HTTP 服务器已启动，请在浏览器访问 http://127.0.0.1:8080/hello")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}