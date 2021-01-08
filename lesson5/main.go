package main

import "fmt"

// 陣列

func main() {
	var x [5]float64
	x[0] = 211
	x[1] = 25
	x[2] = 99
	x[3] = 100
	x[4] = 33

	fmt.Println(x)
	var total float64 = 0

	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	var total2 float64 = 0

	for _, value := range x {
		total2 += value
	}

	fmt.Println(total2)

	y := [3]float64{
		23,
		45,
		33,
		// 21,
	}

	fmt.Println(y)

}
