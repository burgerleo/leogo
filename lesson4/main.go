package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		fmt.Println(sum, "+", i, "=", sum+i)
		sum += i
	}

	fmt.Println(sum)

	sum2 := 1
	for sum2 <= 2000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	fmt.Println(pow(3, 2, 10))

	fmt.Println(pow(3, 3, 20))

	fmt.Println(sqrt(2), sqrt(9), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	// pow 開平方 x^n

	fmt.Println(math.Pow(x, n))

	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// 平方根
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
