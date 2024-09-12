package demo

import (
	"fmt"
	"reflect"
)

// 接口是一种实现多态的机制

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

// 非指针接收者方法
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// 指针接收者方法
func (r *Rectangle) Area2() float64 {
	return r.length * r.width
}

// 非指针接收者方法
func (r Rectangle) Expand1() {
	r.length = r.length * 2
	r.width = r.width * 2
}

// 指针接收者方法
func (r *Rectangle) Expand2() {
	r.length = r.length * 2
	r.width = r.width * 2
}

func (r Rectangle) Perimeter() float64 {
	return (r.length + r.width) * 2
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (c *Circle) Expand() {
	c.radius = c.radius * 2
}

func (r Rectangle) PrintInfo() {
	fmt.Println("length: ", r.length, "\twidth: ", r.width)
	fmt.Println("rec.Area(): ", r.Area(), "\trec.Area2(): ", r.Area2())
	fmt.Println()
}

func ParamPassTest() {
	var (
		rec  = Rectangle{length: 3, width: 4}
		recp = &rec
	)
	rec.Expand1()
	rec.PrintInfo() // length:  3      width:  4

	rec.Expand2()   // 使用 非指针类型 调用 指针接收者方法，编译时会自动取引用
	rec.PrintInfo() // length:  6      width:  8

	recp.Expand1()  // 使用 指针类型 调用 非指针接收者方法，编译时会自动解引用
	rec.PrintInfo() // length:  6      width:  8

	recp.Expand2()
	rec.PrintInfo() // length:  12     width:  16

	(*recp).Expand1()
	rec.PrintInfo() // length:  12     width:  16

	(*recp).Expand2() // 自动取引用
	rec.PrintInfo()   // length:  24     width:  32
}

func PrintInterfaceInfo(shape Shape) {
	fmt.Println("Area: ", shape.Area(), "\tPerimeter: ", shape.Perimeter())
	fmt.Println("Type: ", reflect.TypeOf(shape), "\tValue: ", reflect.ValueOf(shape))

}

func InterfaceTest() {
	// ParamPassTest()
	var (
		rec  = Rectangle{length: 3, width: 4}
		cir  = Circle{radius: 5}
		recp = &rec
		cirp = &cir
	)

	// Rectangle 类型使用 非指针接收器 实现接口方法，但是 Rectangle 和 *Rectangle 都是 Shape 接口类型
	// 因为「包装方法」。编译器会为 方法接收器为非指针类型的方法 自动生成 一个方法接收器是指针类型的方法。
	// 为什么编译器要生成「包装方法」? 是为了支持接口类型调用 非指针接收者方法。
	PrintInterfaceInfo(rec)
	PrintInterfaceInfo(recp)

	// Circle 类型使用 指针接收器 实现接口方法，只有 *Circle 才是 Shape 接口类型，Circle 不是 Shape 接口类型
	PrintInterfaceInfo(cirp)
	// PrintInterfaceInfo(cir)
}
