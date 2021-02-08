package main

import "fmt"

func zero(x int) {
	x = 0
}

func yToZero(abc *int) {
	*abc = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5

	y := 5
	yToZero(&y)
	fmt.Println(y)
}
