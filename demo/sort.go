package demo

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	age    int
	height float64
	weight float64
}

type People struct {
	people []Person
	by     func(p *Person, q *Person) bool
}

func (pp *People) Len() int {
	return len(pp.people)
}

func (pp *People) Swap(i int, j int) {
	pp.people[i], pp.people[j] = pp.people[j], pp.people[i]
}

func (pp *People) Less(i int, j int) bool {
	return pp.by(&pp.people[i], &pp.people[j])
}

func SortTest_1() {
	var (
		numList     = []int{11, 2, 4, 5, 20, 5, 77, 1}
		float64List = []float64{4.2, 5.9, 12.3, 10.0, 50.4, 99.9, 31.4, 27.81828, 3.14}
		strList     = []string{"Cupid", "Augustus", "Balesego", "Bernhard", "Viliulfo", "Temwani", "Cruz", "Desi", "Flossie"}
	)

	// 升序排序
	sort.Ints(numList)
	sort.Float64s(float64List)
	sort.Strings(strList)

	fmt.Println(numList)
	fmt.Println(float64List)
	fmt.Println(strList)
	fmt.Println()

	// 降序排序
	// sort.IntSlice/sort.Float64Slice/sort.StringSlice 实现了Len(), Less(i, j), Swap(i, j)
	// sort.Reverse(slice) 调换 sort.Interface.Less，即比较函数
	sort.Sort(sort.Reverse(sort.IntSlice(numList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float64List)))
	sort.Sort(sort.Reverse(sort.StringSlice(strList)))

	fmt.Println(numList)
	fmt.Println(float64List)
	fmt.Println(strList)
	fmt.Println()

	var persons = []Person{
		{"Cupid", 18, 175, 65},
		{"Augustus", 20, 180, 66.6},
		{"Bernhard", 22, 178, 75.9},
		{"Desi", 19, 166, 55},
		{"Balesego", 21, 168, 60.2},
	}

	fmt.Println(persons)
	// 按name升序排序
	sort.Sort(&People{persons, func(p, q *Person) bool {
		return p.name < q.name
	}})
	fmt.Println("Sorted by name: ", persons)

	// 按age升序排序
	sort.Sort(&People{persons, func(p, q *Person) bool {
		return p.age < q.age
	}})
	fmt.Println("Sorted by age: ", persons)

	// 按height降序排序
	sort.Sort(&People{persons, func(p, q *Person) bool {
		return p.height > q.height
	}})
	fmt.Println("Sorted by height (reversed): ", persons)
}
