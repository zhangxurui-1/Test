package demo

import (
	"fmt"
	"sort"
)

// 闭包=函数+引用环境
// https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.6.html

func f(i int) func() int {
	return func() int {
		i++ // i不分配在f的栈中
		return i
	}
}

// 内存逃逸（escape analyze）
// 闭包定义时并未执行
func foo1(x int) []func() {
	var fs []func()
	for i := 0; i < 3; i++ {
		fs = append(fs, func() {
			x++
			fmt.Printf("fs[%v]: \t%v\n", i, x)
		})
	}
	// x++
	return fs
}

func foo2(x int) []func() {
	var fs []func()
	var xs []int
	for i := 0; i < 3; i++ {
		xs = append(xs, x)
		fs = append(fs, func() {
			xs[i]++
			fmt.Printf("fs[%v]: \t%v\n", i, xs[i])
		})
	}
	/*
		for i := 0; i < len(xs); i++ {
			xs[i]++
		}
	*/

	return fs
}

func SortTest() {
	var pps = []string{"Cupid", "Augustus", "Balesego", "Bernhard", "Viliulfo", "Temwani", "Cruz", "Desi", "Flossie"}
	sort.Slice(pps, func(i, j int) bool {
		for k := 0; k < min(len(pps[i]), len(pps[j])); k++ {
			if pps[i][k] < pps[j][k] {
				return true
			} else if pps[i][k] > pps[j][k] {
				return false
			} else {

			}
		}
		return true
	})

	fmt.Println(pps)
}

func ClosureTest() {
	var c1 = f(0)
	var c2 = f(0) // c1跟c2引用的是不同的环境

	fmt.Printf("c1(): \t%v\n", c1())
	fmt.Printf("c1(): \t%v\n", c1())
	fmt.Printf("c2(): \t%v\n", c2())

	fmt.Println()
	var fs1 = foo1(0)
	for _, f := range fs1 {
		f()
	}

	fmt.Println()
	var fs2 = foo2(0)
	for _, f := range fs2 {
		f()
	}

	SortTest()
}
