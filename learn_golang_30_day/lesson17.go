package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

// 區塊
type square struct {
	width, height float64
}

// 圈
type circle struct {
	radius float64
}

type triangle struct {
	ground, height float64
}

// 面積
func (s square) area() float64 {
	return s.width * s.height
}

// 周長
func (s square) perimeter() float64 {
	return 2*s.width + 2*s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (t triangle) area() float64 {

	return (t.ground * t.height) / 2
}

func (t triangle) perimeter() float64 {
	// 另一邊長
	c := math.Sqrt(math.Pow(t.ground, 2) + math.Pow(t.height, 2))

	return t.ground + t.height + c
}

// 測量
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("面積", g.area())
	fmt.Println("周長", g.perimeter())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}
	t := triangle{ground: 3, height: 4}

	measure(s)
	measure(c)
	measure(t)
}
