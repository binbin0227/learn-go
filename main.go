package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	secret string // 注意：这是一个小写开头的私有字段
}

func ReadStruct(obj any) {
	t := reflect.TypeOf(obj)  // 拿 X光机
	v := reflect.ValueOf(obj) // 拿 机械臂

	fmt.Println("=== 开始扫描保险箱 ===")
	// t.NumField() 获取一共有几个格子
	for i := 0; i < t.NumField(); i++ {
		// 用 X光机 看第 i 个格子的名字和标签
		fieldName := t.Field(i).Name
		fieldTag := t.Field(i).Tag.Get("json")
		
		// 用 机械臂 掏出第 i 个格子里的真实数据
		fieldValue := v.Field(i).Interface()

		fmt.Printf("字段: %s | 标签: %s | 值: %v\n", fieldName, fieldTag, fieldValue)
	}
}

func main() {
	u := User{Name: "积积", Age: 23, secret: "123456"}
	ReadStruct(u)
}