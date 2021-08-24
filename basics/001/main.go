package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

func test(a []int) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&a))

	fmt.Printf("%p %v %p %v \n", a, a, &a, &a)
	spew.Dump(sh)
}

/*
var a []int 栈中存一个地址 地址指向 SliceHeader 零值为 0x0

*/
func main() {
	var a []int

	fmt.Printf("%p %v %p %v \n", a, a, &a, &a) // 第一个是栈中存储的地址->指向堆中SliceHeader 第三个是本栈的地址
	spew.Dump(a)
	test(a)
	fmt.Println("-------------------------------")

	a = append(a, 1)

	fmt.Printf("%p %v %p %v \n", a, a, &a, &a)
	spew.Dump(a)
	test(a)
	fmt.Println("-------------------------------")

	a = append(a, 2)
	fmt.Printf("%p %v %p %v \n", a, a, &a, &a)
	spew.Dump(a)
	test(a)
	fmt.Println("-------------------------------")
}
