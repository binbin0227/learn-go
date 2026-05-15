package main

import "fmt"

// ==========================================
// 1. 【制定规则】定义接口 (甲方爸爸)
// 潜台词：我不管你是什么软件，只要你能扣钱 (有 Pay 方法)，你就是个支付工具！
// ==========================================
type Payer interface {
	Pay(amount float64)
}

// ==========================================
// 2. 【具体实现】定义结构体和它们的方法 (乙方干活)
// 注意：它们脑门上绝没有 implements Payer 的字样！
// ==========================================

// 微信支付
type WeChatPay struct {
	Account string
}
func (w WeChatPay) Pay(amount float64) {
	fmt.Printf("🚀 正在调用微信支付... 账号[%s] 扣款: %.2f 元\n", w.Account, amount)
}

// 支付宝
type AliPay struct {
	Account string
}
func (a AliPay) Pay(amount float64) {
	fmt.Printf("🛡️ 正在调用支付宝... 账号[%s] 扣款: %.2f 元\n", a.Account, amount)
}


// ==========================================
// 3. 【核心价值】面向接口编程 (神奇的结算台)
// ==========================================
// 注意看参数：这里接收的是 Payer 接口！而不是具体的微信或支付宝。
func Checkout(p Payer, amount float64) {
	fmt.Println(">>> 欢迎来到收银台，开始处理订单 <<<")
	// 结算台根本不知道此时的 p 到底是微信还是支付宝
	// 它只知道：p 一定会 Pay()！直接调用就行了！
	p.Pay(amount) 
	fmt.Println(">>> 订单处理完成 <<<\n")
}

// ==========================================
// 4. 【系统运行】
// ==========================================
func main() {
	// 实例化具体的支付工具
	wx := WeChatPay{Account: "枫枫_WX"}
	ali := AliPay{Account: "枫枫_ZFB"}

	// 场景 A：用户在前端点击了“微信支付”
	// 编译器瞬间检查：wx 有 Pay 方法吗？有！放行！
	Checkout(wx, 99.8)

	// 场景 B：用户在前端点击了“支付宝”
	// 编译器瞬间检查：ali 有 Pay 方法吗？有！放行！
	Checkout(ali, 299.0)
}