package demo

import (
	"fmt"
)

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~string
}

// comparable 是一个接口，指代所有可以用!=或者==来进行比较的元素
// The comparable interface may only be used as a type parameter constraint
type UnorderedSet[T comparable] map[T]struct{}

func (set UnorderedSet[T]) Insert(value T) {
	if _, exist := set[value]; !exist {
		set[value] = struct{}{}
	}
}

func (set UnorderedSet[T]) Remove(value T) {
	delete(set, value)
}

func (set UnorderedSet[T]) PrintInfo() {
	fmt.Printf("%T: {", set)
	var tmp []T
	for k := range set {
		tmp = append(tmp, k)
	}
	for i := 0; i < len(tmp)-1; i++ {
		fmt.Print(tmp[i], ", ")
	}
	fmt.Print(tmp[len(tmp)-1], "}\n")
}

func max[T ordered](nums []T) T {
	// [T ordered] 是Go泛型的类型参数列表（type parameters list）
	// ordered 是类型参数的类型约束（type constraint）
	var m T
	for _, v := range nums {
		if m < v {
			m = v
		}
	}
	return m
}

type Node[T any] struct {
	next  *Node[T]
	value T
}

type List[T any] struct {
	head   *Node[T]
	length int
}

// 不能写成 func (list List[T]) Reset() {}
func (list *List[T]) Reset() {
	list.head = nil
	list.length = 0
}

// 不能写成 func (list List[T]) Insert(value T)
func (list *List[T]) Insert(value T) {
	var (
		node = new(Node[T])
	)
	node.next = list.head
	node.value = value
	list.head = node
	list.length++
}

func (list *List[T]) PrintInfo() {
	fmt.Printf("%T: {", *list)
	var p *Node[T] = nil
	for p = list.head; p != nil && p.next != nil; p = p.next {
		fmt.Print(p.value, ", ")
	}
	if p != nil {
		fmt.Print(p.value)
	}
	fmt.Print("}\n")
}

func ListTest() {
	var (
		intlistp    = new(List[int])
		stringlistp = new(List[string])
		nums        = []int{0, 1, 2, 3, 4, 5}
		strs        = []string{"why", "does", "it", "have", "so", "much", "?"}
	)
	intlistp.Reset()
	stringlistp.Reset()
	for _, n := range nums {
		intlistp.Insert(n)
	}

	for _, s := range strs {
		stringlistp.Insert(s)
	}

	intlistp.PrintInfo()
	stringlistp.PrintInfo()

}

func GenericTest() {
	var (
		v1 = []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
		v2 = []byte{'a', 'r', 'd', 'z', 'p'}
		v3 = []string{"github", "zhangxurui", "zhangxuruii"}
	)

	fmt.Println(max(v1))
	fmt.Println(max[int](v1))
	fmt.Println(max[uint8](v2))
	fmt.Println(max[string](v3))
	fmt.Println()

	var set = make(UnorderedSet[string])
	set["test"] = struct{}{}
	v, ok := set["test"]
	fmt.Println(v, " ", ok)

	// 这里 指针接收者方法 和 非指针接收者方法 效果相同
	var (
		uoSet  = make(UnorderedSet[string])
		uoSetp = &uoSet
	)
	uoSet.Insert("")
	uoSet.Insert("zhang")
	uoSetp.Insert("xu")
	uoSetp.Insert("xu")
	uoSetp.Insert("rui")
	uoSet.PrintInfo()

	uoSetp.Remove("xu")
	uoSetp.PrintInfo()

	// 这里，自定义结构体类型要区分 指针接收者方法 和 非指针接收者方法
	ListTest()
}
