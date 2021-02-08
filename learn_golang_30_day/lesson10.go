package lesson10

import "fmt"

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("Hello", "Leo")
	fmt.Println(a, b)

	x, y := split(12)
	fmt.Println(x, y)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (b, a string) {
	a = x + "1"
	b = y + "2"
	return
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
