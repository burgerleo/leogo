package main

import "fmt"

// 型態
func main() {
	fmt.Println("1 + 1 =", 1+1)

	var a bool
	a = true
	fmt.Println("a =", a)

	b := false
	fmt.Println("b =", b)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!!true)

	i := 1
	fmt.Println("Hello World", i)

	add(2, 3)
}

func add(a int, b int) {

	fmt.Println(a, "+", b, "=", a+b)
}
