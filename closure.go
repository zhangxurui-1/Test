package main

import (
	"fmt"
)

// 闭包=函数+引用环境
// https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.6.html

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func ClosureTest() {
	var c1 = f(0)
	var c2 = f(0)

	fmt.Println(c1())
	fmt.Println(c2())
}
