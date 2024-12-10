package main

import (
	"demo/demo"
	"fmt"
	"os"
)

type info struct {
	height int
	weight int
	age    int
}

func f1() {
	fmt.Println("call f1()")
}

func f2() {
	fmt.Println("call f2()")
}

// 修改fmt.Println()
func Println(a interface{}) (n int, err error) {
	println("--------------------")
	return fmt.Fprintln(os.Stdout, a)
}

// plugin测试
func PluginTest() {
	err := demo.LoadAndInvokeFromPlugin("./plugins/plugin1.so")
	if err != nil {
		fmt.Println("LoadAndInvokeFromPlugin error: ", err)
		os.Exit(1)
	}
	fmt.Println("LoadAndInvokeFromPlugin OK")
}

// plugin init()函数调用测试
func LoadPluginTest() {
	// 第一次加载
	// 只有第一次加载会执行init()
	fmt.Println("LoadPlugin...")
	err := demo.LoadPlugin("./plugins/plugin1.so")
	if err != nil {
		fmt.Println("LoadPlugin error: ", err)
		os.Exit(1)
	}
	fmt.Println("LoadPlugin OK")

	// 第二次加载
	fmt.Println("Re-LoadPlugin...")
	err = demo.LoadPlugin("./plugins/plugin1.so")
	if err != nil {
		fmt.Println("LoadPlugin error: ", err)
		os.Exit(1)
	}
	fmt.Println("Re-LoadPlugin OK")
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
	fmt.Printf("type of c: \t%T\n", c)

	fmt.Println()
	var a = [3]int{1, 2, 3}
	fmt.Printf("a[:]: \t%v\n", a[:])
	fmt.Printf("type of a: \t%T\n", a)
	fmt.Printf("type of a[:]: \t%T\n", a[:])
	var pa *[3]int = &a
	fmt.Printf("pa: \t%v(%p)\n\n", pa, pa)

	var b = []int{1, 2, 3}
	var pb1 *[]int = &b
	var pb2 *[3]int = (*[3]int)(b)
	fmt.Printf("pb1: \t%v(%p)\n", pb1, pb1)
	fmt.Printf("pb2: \t%v(%p)\n\n", pb2, pb2)

	for i, c := range s3 {
		fmt.Printf("%v\t%v\n", i, c)
	}

	fmt.Println()
	for k, v := range m {
		fmt.Printf("%v\t%v\n", k, v)
	}

	var t = make(chan int, 100) // channel
	fmt.Printf("t: \t%v\n", t)

	var n int = 1
	switch n {
	case 1:
		fallthrough
	case 2, 3, 4:
		fmt.Printf("case 1")
	case 5:
		fmt.Printf("case 2")
	case 6, 7:
		fmt.Printf("case 3")
	}

	fmt.Println()
	var (
		n1 int = 1
		n2 int = 2
		n3 int = 3
	)
	switch {
	case n1 < n2:
		f1()
	case n2 < n3:
		f2()
	default:
		fmt.Println("default")
	}

	fmt.Println()

	// demo.ClosureTest()

	// demo.DeferTest()

	// demo.MethodTest()

	// demo.InterfaceTest()

	// demo.GenericTest()

	// demo.SortTest_1()

	// demo.GoRoutineTest()

	// PluginTest()

	LoadPluginTest()
}
