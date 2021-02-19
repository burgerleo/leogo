package main

import "fmt"

func main() {
	leo := Person{
		name: "Leo",
		age:  27,
	}
	fmt.Println(leo)

	leoS := Student{leo, "abc"}
	fmt.Println(leoS)

	p2Sch := Student{
		Person: Person{
			name: "p2",
			age:  10,
		},
		school: "aaabc",
	}

	fmt.Println(p2Sch)

}

// Person 人類
type Person struct {
	name string
	age  int
}

// Student 學校
type Student struct {
	Person        // 模擬繼承
	school string // 子類擁有的屬性
}
