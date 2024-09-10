package demo

import (
	"fmt"
)

func deferf() (result int) {
	// defer是在return之前执行的
	defer func() {
		result++
	}()

	return 0
}

func deferf2() int {
	// 函数返回的过程是: (1)给返回值赋值; (2)调用defer表达式; (3)返回到调用函数中
	var t int = 5
	defer func() {
		t = t + 5
	}()

	return t
}

/*
func gof2() int {
	var t int = 5
	fmt.Println("addr of t: \t", &t)
	go func() {
		for {
			fmt.Println("go func()")
			time.Sleep(1 * time.Second)
		}
	}()

	return t
}
*/

func A1() {
	fmt.Println("call A1()")
}

func A2() {
	fmt.Println("call A2()")
}

func A3() {
	fmt.Println("call A3()")
}

func DeferTest() {
	fmt.Println("deferf(): \t", deferf())
	fmt.Println("deferf(): \t", deferf())

	fmt.Println()
	fmt.Println("deferf2(): \t", deferf2())

	fmt.Println()
	defer A1()
	defer A2()
	defer A3()
	fmt.Println("exit DeferTest()")
}
