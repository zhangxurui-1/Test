package demo

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// 方法是带有receiver的函数
func (p Point) length() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func MethodTest() {
	var p = Point{3, 4}
	fmt.Println(p)
	fmt.Println("p.length(): \t\t", p.length())
	fmt.Println("Point.length(p): \t", Point.length(p))
}
