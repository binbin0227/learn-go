package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp","127.0.0.1:10086")
	if err!=nil{
		fmt.Println("连接服务器失败:", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("你好，我是 Go 客户端！\n"))
	buf:=make([]byte,1024)
	n,err:=conn.Read(buf)
	if err!=nil{
		fmt.Println("读取响应失败:", err)
		return
	}
	fmt.Println("收到服务器回复:", string(buf[:n]))
}