package main

import (
	"fmt"
	"time"
)

type Deliverer interface {
	Deliver(id string) bool
}
type Robot struct {
	Name  string
	Speed int
}

func (r *Robot) Deliver(id string) bool {
	fmt.Printf("机器人 %s 开始派送快递 %s ...\n", r.Name, id)
	time.Sleep(time.Duration(r.Speed) * time.Second)
	return true
}
func HandleTask(id string, d Deliverer, timeout time.Duration) {
	resultCh := make(chan bool)
	go func() {
		resultCh <- d.Deliver(id)
	}()
	select {
	case ok := <-resultCh:
		if ok {
			fmt.Printf("任务成功：快递 %s 已送达\n", id)
		}
	case <-time.After(timeout):
		fmt.Printf("❌ 任务失败：快递 %s 派送超时！\n", id)
	}
}
func main() {
	xl := Robot{Name: "xiaolan", Speed: 5}
	HandleTask("SF123", &xl, 3*time.Second)
	HandleTask("SF456", &xl, 6*time.Second)
}
