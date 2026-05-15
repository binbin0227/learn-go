package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Student struct {
	Name   string `orm:"name"`
	Age    int    `orm:"age"`
	secret string
}

func GenerateInsertSQL(obj any) string {
	t := reflect.TypeOf(obj)
	v :=reflect.ValueOf(obj)
	if t.Kind()!=reflect.Struct{
		return "err"
	}

	//根据结构体名字生成表名 (Student -> students)
	tableName:=strings.ToLower(t.Name())+"s"

	// 准备两个切片，分别装“列名”和“值”
	var columns []string
	var values []string
	for i:=0;i<t.NumField();i++{
		typeField :=t.Field(i)
		valueField :=v.Field(i)

		// ⭐️ 重点防坑：跳过私有字段！
		// 如果首字母是小写（未导出），反射去强行读写会引发致命危险，直接跳过！
		if !typeField.IsExported() {
			continue
		}

		// 第三步：提取列名（优先读 orm 标签，如果没有就用字段名的小写）
		if !typeField.IsExported() {
			continue
		}

		// 第三步：提取列名（优先读 orm 标签，如果没有就用字段名的小写）
		colName:=typeField.Tag.Get("orm")
		if colName==""{
			colName=strings.ToLower(typeField.Name)
		}
		columns=append(columns, colName)

		// 第四步：提取真值，并进行【智能格式化】
		// SQL 语法规定：字符串必须加单引号 '积积'，数字不需要加单引号 23
		if valueField.Kind()==reflect.String{
			values=append(values, fmt.Sprintf("'%s'",valueField))
		}else{
			values=append(values, fmt.Sprintf("%v",valueField))
		}
	}
	// 第五步：把切片用逗号拼起来
	// columns 变成: "name, age"
	colsStr := strings.Join(columns, ", ")
	// values 变成: "'积积', 23"
	valsStr := strings.Join(values, ", ")

	// 终极拼装！
	finalSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, colsStr, valsStr)
	return finalSQL
}
func main() {
	// 测试 1：测学生
	stu := Student{Name: "积积", Age: 23, secret: "你看不到我"}
	sql1 := GenerateInsertSQL(stu)
	fmt.Println("学生插入语句：")
	fmt.Println(sql1)
	fmt.Println("------------------------")

	// 测试 2：测一个完全不同的匿名结构体，看框架能不能顶住！
	type Order struct {
		OrderNo string `orm:"order_no"`
		Amount  int    `orm:"total_amount"`
	}
	ord := Order{OrderNo: "SN123456", Amount: 998}
	sql2 := GenerateInsertSQL(ord)
	fmt.Println("订单插入语句：")
	fmt.Println(sql2)
}