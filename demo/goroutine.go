package demo

import (
	"fmt"
)

func GoRoutineTest() {
	var c = make(chan int)
	go func(a int, b int) {
		c <- a + b
	}(3, 4)
	fmt.Println(<-c)
}
