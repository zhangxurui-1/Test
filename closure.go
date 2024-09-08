package main

import (
	"fmt"
)

func foo(x int) []func() {
	var fs []func()
	var values = []int{1, 2, 3, 4, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo value: \t%v\n", val+x)
		})
	}

	return fs
}

func Output() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%d ", i)
		}()
		// time.Sleep(1 * time.Second)
	}
	fmt.Printf("\nOutput() end\n\n")
}

func Output2() {
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Printf("%d ", i)
		}()
	}
	fmt.Printf("\nOutput2() end\n\n")
}

func ClosureTest() {
	var fs = foo(10)
	for _, f := range fs {
		f()
	}

	Output()
	Output2()
}
