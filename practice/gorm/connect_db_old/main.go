package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product 定义一个产品模型 (对应数据库里的一张表)
type Product struct {
	gorm.Model // 这是一个神奇的嵌套，会自动帮你加上 ID、创建时间、更新时间、删除时间
	Code  string
	Price uint
}

func main() {
	// 1. 配置 Docker 里 MySQL 的连接信息 (DSN)
	// 格式: 用户名:密码@tcp(地址:端口)/数据库名?参数
	dsn := "root:123456@tcp(127.0.0.1:3306)/learngo?charset=utf8mb4&parseTime=True&loc=Local"

	// 2. 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ 连接数据库失败，请检查 Docker 里的 MySQL 是否启动了！")
	}
	fmt.Println("✅ 成功连上 Docker 里的 MySQL！")

	// 3. 自动迁移 (魔法建表)
	// 只要这一句，GORM 就会自动在数据库里建一张叫 products 的表
	db.AutoMigrate(&Product{})
	fmt.Println("✅ 自动建表完成！")

	// ==================== 开始 CRUD (增删改查) ====================

	// 【增】Create
	db.Create(&Product{Code: "D42", Price: 100})
	fmt.Println("✅ 成功插入一条数据")

	// 【查】Read
	var product Product
	// 找到第一条记录放到 product 变量里
	db.First(&product, 1) 
	fmt.Printf("🔍 查到的产品: 编号=%s, 价格=%d\n", product.Code, product.Price)

	// 【改】Update
	// 把刚才查出来的这个产品的价格更新为 200
	db.Model(&product).Update("Price", 200)
	fmt.Println("✅ 价格更新成功")

	// 【删】Delete
	// 删除刚才那个产品
	db.Delete(&product, 1)
	fmt.Println("✅ 数据删除完毕")
}