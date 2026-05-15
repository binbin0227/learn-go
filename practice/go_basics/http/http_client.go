package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 1. 发起 GET 请求 (就像在浏览器地址栏敲回车)
	resp, err := http.Get("http://127.0.0.1:8080/hello")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	
	// ⚠️ 大坑警告：只要请求成功了，就必须手动关闭响应体 (Body)，否则会内存泄漏！
	defer resp.Body.Close()

	// 2. 查看服务器的状态码 (比如 200 代表成功，404 代表找不到)
	fmt.Println("服务器状态码:", resp.StatusCode)

	// 3. 读取服务器返回的具体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取内容失败:", err)
		return
	}

	fmt.Println("服务器返回的内容是:", string(body))
}