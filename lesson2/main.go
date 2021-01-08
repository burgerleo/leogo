package main

import "fmt"

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"
var aa1, aa2, aa3 = 1123, "aab", true

var (
	a1 int    = 1
	a2 string = "a23"
	a3        = 1
)

func main() {
	fmt.Println(x, y, z)

	fmt.Println(c, python, java)
	fmt.Println(aa1, aa2, aa3)

	fmt.Println(a1)
	fmt.Println(a2)
	c2 := "123"

	fmt.Println(c2)

}
