package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模版
	t, err := template.ParseFiles("./Hello.tmpl")
	if err != nil {
		fmt.Printf("Parse Files failed, err: %v", err)
		return
	}
	// 3. 渲染模版
	err = t.Execute(w, "Bin")
	if err != nil {
		fmt.Printf("Render Files failed, err: %v", err)
		return
	}
}
func main() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":10086", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err: %v", err)
	}
}
