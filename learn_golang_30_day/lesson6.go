package main

import "fmt"

func main() {
	// x := make([]int, 5, 10)

	arr := [5]float64{1, 2, 3, 4, 5}

	x := arr[0:5]

	fmt.Println(x)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	slice1 = append(slice1, 10, 11)
	fmt.Println(slice1, slice2)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999, 330}
	fmt.Println("numArr04", numArr04)

	strArr := [...]string{1: "tom", 0: "jack", 2: "leo"}
	fmt.Println("strArr", strArr)

	citys := [...]string{"北京", "上海", "深圳"} //索引從0-2
	// 1 根據索引遍歷
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	// 2 for range 遍歷
	// i 是索引，v是值
	for i, v := range citys {
		fmt.Println(i, v)
	}

}
