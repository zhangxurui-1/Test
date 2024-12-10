package demo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"plugin"
)

func init() {
	log.Println("demo init")
}

// 定义接口
type MyInterface interface {
	M1()
}

func LoadAndInvokeFromPlugin(pluginPath string) error {
	p, err := plugin.Open(pluginPath) // 在第一次Open时，plugin中所有不属于主程序包的init函数将被调用
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 导出整型变量
	v, err := p.Lookup("V")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("type of v: %T\n", v) // v是指针，是*int类型
	*v.(*int) = 15

	// 导出函数变量
	f, err := p.Lookup("F")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.(func())()

	// 导出自定义类型变量
	f1, err := p.Lookup("Foo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	i, ok := f1.(MyInterface) // 类型断言
	if !ok {
		return errors.New("f1 does not implement MyInterface")
	}
	i.M1()

	return nil
}

func LoadPlugin(pluginPath string) error {
	_, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}
