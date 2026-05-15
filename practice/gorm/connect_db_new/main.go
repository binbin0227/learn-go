package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product 定义模型（包含 gorm.Model 的神奇软删除机制）
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 1. 换成你的 Docker MySQL 连接密码和数据库名
	dsn := "root:123456@tcp(127.0.0.1:3306)/learngo?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ 连接数据库失败")
	}
	fmt.Println("✅ 数据库连接成功！")

	// 2. 准备上下文 (大厂规范，用于控制超时等)
	ctx := context.Background()

	// 3. 自动建表
	db.AutoMigrate(&Product{})
	fmt.Println("✅ 自动建表完成！")

	// ==================== 开始最新版的 CRUD 操作 ====================

	// 【增】Create - 创建一条记录
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
	if err == nil {
		fmt.Println("✅ 成功插入产品 D42")
	}

	// 【查】Read - 查询主键为 1 的产品
	// 爽点：不需要提前 var product Product，查出来直接赋值！
	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)
	if err == nil {
		fmt.Printf("🔍 查找到产品：编号=%s, 价格=%d\n", product.Code, product.Price)
	}

	// 【改】Update - 将产品价格更新为 200
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
	if err == nil {
		fmt.Println("✅ 价格已更新为 200")
	}

	// 【删】Delete - 软删除该产品
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
	if err == nil {
		fmt.Println("✅ 产品已删除（软删除）")
	}
}
