package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type S struct {
	a int
	b []int
}

func (s *S) M1() {}
func F1(s *S)    {}

func (s S) M2() {}
func F2(s S)    {}

func New1() *S {
	return &S{
		b: []int{0},
	}
}

func New2() S {
	return S{
		b: []int{0},
	}
}

func Check1() {
	n1 := New1() // 堆中存放 S, n1 存地址
	n2 := n1     // n2 存地址指向 n1
	n1 = New1()  // 堆中存放 新S, n1 存地址
	// 现在 n1 n2 指向不同的 S
	fmt.Println(n1, n2)
	n1.a = 10
	n2.a = 20
	fmt.Println(n1, n2)
}

func Check2() {
	n1 := New1()  // 堆中存放 S, n1 存地址
	n2 := n1      // n2 存地址指向 n1
	*n1 = *New1() // 新创建的 S 数据拷贝到 原来的 S
	// 现在 n1 n2 指向相同的 S
	fmt.Println(n1, n2)
	n1.a = 10
	n2.a = 20
	fmt.Println(n1, n2)
}

func Check3() {
	n1 := New2() // 栈中存放 a(int) 、 b([]int) SliceHeader 两个字段
	n2 := n1     // 拷贝两个字段
	// 现在 n1 和 n2 的 a 存在栈中不同位置， b 存的地址相同
	fmt.Println(n1, n2)
	n1.a = 10
	n1.b[0] = 11
	n2.a = 20
	n2.b[0] = 21
	fmt.Println(n1, n2) // a值不同 b值相同
	n1.b = []int{11}    // b 指向不同的堆位置
	fmt.Println(n1, n2) // a值不同 b值相同
}

func main1() {
	var a int8
	var b int8
	var c []int8
	// fmt.Printf("%p %p %p\n", &a, &b, &c)
	// c = []int8{}
	// fmt.Printf("%p %p %p\n", &a, &b, &c)
	// c = append(c, 1)
	fmt.Printf("%p %p %p\n", &a, &b, &c)
	p1 := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	fmt.Printf("0x%x\n", p1.Data)
	fmt.Printf("%p\n", &p1.Len)
	fmt.Printf("%p\n", &p1.Cap)

	fmt.Println("------------------------------------")

	c = append(c, 1)
	fmt.Printf("%p %p %p\n", &a, &b, &c)
	p1 = (*reflect.SliceHeader)(unsafe.Pointer(&c))
	fmt.Printf("0x%x\n", p1.Data)
	fmt.Printf("%p\n", &p1.Len)
	fmt.Printf("%p\n", &p1.Cap)

}

func main() {
	var a int
	a1 := [2]int{1, 2}
	s1 := a1[0:1]
	s2 := a1[:]
	p1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	p2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))

	fmt.Printf("0x%x\n", &a)
	fmt.Printf("0x%x\n", &a1[0])
	fmt.Printf("0x%x\n", &a1[1])
	fmt.Printf("0x%x\n", &p1)
	fmt.Printf("0x%x\n", &p2)

	fmt.Printf("0x%x\n", p1.Data)
	fmt.Printf("0x%x\n", p2.Data)
}
