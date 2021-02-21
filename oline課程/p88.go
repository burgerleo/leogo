package main

import "fmt"

func main() {
	p1 := worker{
		name: "leo",
		age:  27,
	}
	fmt.Println(p1)
	p1.work()
	p1.getStatus()
	p1.rest()
	p1.getStatus()

}

// worker 工人
type worker struct {
	name   string
	age    int
	status string
}

/*
下列這些是方法 method
一個方法包含接受者的函數，接受者可以是命名的類型或是 struct 的一個值，又或是一個記憶體位置
*/
func (p *worker) work() {
	p.status = "working"
	fmt.Println(p.name, "is working...")
}

func (p *worker) rest() {
	p.status = "resting"

	fmt.Println(p.name, "is resting...")
}

func (p worker) getStatus() {
	fmt.Println(p.name, "status is ...", p.status)
}
