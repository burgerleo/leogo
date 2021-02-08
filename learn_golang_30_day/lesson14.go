package main

import "fmt"

func main() {

	// 延遲執行
	defer func() {
		fmt.Println("first")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end")
	}()

	defer fmt.Println("LEEO")

	f()
}

func f() {
	fmt.Println("test")

	// 觸發錯誤
	panic(1)

	fmt.Println("test2")
}
