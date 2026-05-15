// 你要为一款游戏编写核心的战斗模块。游戏里有不同的职业（比如战士、法师），他们都能发起攻击，但法师有一个特殊的专属技能“闪现”。
// 【开发需求（请严格按照以下规则编写）】
// 1. 定义通用接口（契约）
// 定义一个名为 Attacker 的接口。
// 里面只有一个方法：Attack()，不需要参数，没有返回值。
// 2. 定义基础零件（组合的思想）
// 定义一个 BaseCharacter（基础角色）结构体。
// 包含两个属性：Name (string，角色名字)，Damage (int，基础攻击力)。
// 3. 定义具体职业（鸭子类型的实现）
// 定义 Warrior（战士）结构体，匿名嵌套 BaseCharacter。
// 定义 Mage（法师）结构体，匿名嵌套 BaseCharacter，且多一个独有属性 Mana (int，魔法值)。
// 4. 绑定方法（指针接收者的死亡陷阱）
// 为 Warrior 实现 Attack() 方法：必须使用指针接收者。打印出："⚔️ 战士 [名字] 挥舞大剑，造成 [Damage] 点物理伤害！"
// 为 Mage 实现 Attack() 方法：必须使用指针接收者。打印出："🔥 法师 [名字] 吟唱火球，消耗 10 点魔法，造成 [Damage] 点魔法伤害！"
// 附加题：单独给 Mage（依然是指针接收者）写一个特有方法 Teleport()，打印出："✨ 法师 [名字] 使用了闪现，瞬间移动了！"
// 5. 核心调度系统（面向接口编程与类型断言）
// 写一个全局函数 StartCombat(a Attacker)。
// 在函数内，调用传入参数的 Attack() 方法。
// 致命考验：在 StartCombat 函数内部，使用类型断言。如果发现这个 Attacker 的真身是一个 Mage（法师），请立刻调用他的 Teleport() 方法。如果不是法师，什么都不做。
// 6. 运行测试（main 函数）
// 实例化一个战士（比如名字叫“亚瑟”，伤害 100）。
// 实例化一个法师（比如名字叫“安琪拉”，伤害 80，魔法值 200）。
// 把他们分别扔进 StartCombat() 函数里执行。
package main

import (
	"fmt"
)

type Attacker interface {
	Attack()
}
type BaseCharacter struct {
	Name   string
	Damage int
}
type Warrior struct {
	BaseCharacter
}
type Mage struct {
	BaseCharacter
	Mana int
}

func (w *Warrior) Attack() {
	fmt.Printf("⚔️ 战士 %v 挥舞大剑，造成 %v 点物理伤害！", w.Name, w.Damage)
}
func (m *Mage) Attack() {
	fmt.Printf("🔥 法师 %v 吟唱火球，消耗 10 点魔法，造成 %v 点魔法伤害！", m.Name, m.Damage)
}
func (m *Mage) Teleport() {
	fmt.Printf("✨ 法师 %v 使用了闪现，瞬间移动了！", m.Name)
}
func StartCombat(a Attacker) {
	a.Attack()
	if m, ok := a.(*Mage); ok {
		m.Teleport()
	}
}
func main() {
	ys := Warrior{
		BaseCharacter: BaseCharacter{Name: "亚瑟", Damage: 100},
	}
	aql := Mage{
		BaseCharacter: BaseCharacter{Name: "安琪拉", Damage: 80},
		Mana:          200,
	}
	StartCombat(&ys)
	StartCombat(&aql)
}
