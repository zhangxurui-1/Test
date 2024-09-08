package main

import "fmt"

func foo(x int) []func() {
	var fs []func()
	var values = []int{1, 2, 3, 4, 5}
	for _, v := range values {
		fs = append(fs, func() {
			fmt.Printf("%v\n", v+x)
		})
	}

	return fs
}

func ClosureTest() {
	foo(10)
}
