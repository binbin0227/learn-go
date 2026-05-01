package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp","127.0.0.1:10086")
	if err!=nil{
		fmt.Println("监听失败:", err)
		return
	}
	defer listener.Close()
	fmt.Println("🚀 服务器启动，等待连接...")
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("接收连接失败:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("✅ 新客户端连入:", conn.RemoteAddr().String())
	buf:=make([]byte,1024)
	for{
		n,err:=conn.Read(buf)
		if err!=nil{
			// 判断是不是正常的断开 (对方挂电话了)
    if err == io.EOF {
        fmt.Println("👋 客户端正常退出，断开连接。")
    } else {
        // 其他的才是真正的网络异常，比如网线被拔了
        fmt.Println("❌ 真实的读取异常:", err)
    }
    return // 退出当前协程
		}
		received:=string(buf[:n])
		fmt.Println("收到消息: ", received)
		conn.Write([]byte("服务器已收到: " + received))
	}
}