package main

import "fmt"

type info struct {
	height int
	weight int
	age    int
}

func main() {

	fmt.Println("Hello world!")

	var m = map[string]int{}
	m["zxr"] = 22
	m["zzz"] = 20

	fmt.Printf("len(m): \t%v\n", len(m))

	var m1 = make(map[string]info, 100)

	m1["zxr"] = info{181, 70, 22}

	fmt.Printf("len(m1): \t%v\n", len(m1))

	fmt.Printf("m1: \t%v\n", m1)
	fmt.Printf("m1[\"zxr\"]: \t%v\n", m1["zxr"])

	var ru = []rune{0x767d, 0x9d6c, 0x7fd4}
	fmt.Printf("\nru: \t%v\n", ru)
	var r = string(ru) // []rune切片到string的转换
	fmt.Printf("r: \t%v\n", r)

	var s1 = []byte{'h', 1, 2, 'l'}
	fmt.Printf("\ns1: \t%v\n", s1)
	var s2 = string(s1)
	fmt.Printf("s2: \t%v\n", s2)

	var s3 string = "hello"
	var c = s3[1]
	fmt.Printf("s3[1]: \t%v\nc: \t%v\n", s3[1], c)

	for i, c := range s3 {
		fmt.Printf("%v\t%v\n", i, c)
	}

	fmt.Println()
	for k, v := range m {
		fmt.Printf("%v\t%v\n", k, v)
	}

	var t = make(chan int, 100) // channel
	fmt.Printf("t: \t%v\n", t)

	var n int = 5
	switch n {
	case 2, 3, 4:
		fmt.Printf("case 1")
	case 5:
		fmt.Printf("case 2")
	case 6, 7:
		fmt.Printf("case 3")
	}
}
