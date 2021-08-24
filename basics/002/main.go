package main

import (
	"fmt"
)

func ZArrSlice() {
	var a [3]int
	var b = a[:]
	var c []int

	fmt.Printf("%p %v %p %v \n", a, a, &a, &a)
	fmt.Printf("%p %v %p %v \n", b, b, &b, &b)
	fmt.Printf("%p %v %p %v \n", c, c, &c, &c)
}

func ZTestValZeroVal() {
	var a int
	var b float32
	var c string
	var d uintptr
	var e [2]int

	var f []int
	var g map[int]int
	var h chan int
	var i func(int) int

	var j interface{}

	fmt.Printf("%p %v %p %v \n", &a, &a, a, a)
	fmt.Printf("%p %v %p %v \n", &b, &b, b, b)
	fmt.Printf("%p %v %p %v \n", &c, &c, c, c)
	fmt.Printf("%p %v %p %v \n", &d, &d, d, d)
	fmt.Printf("%p %v %p %v \n", &e, &e, e, e)

	fmt.Printf("%p %v %p %v \n", &f, &f, f, f)
	fmt.Printf("%p %v %p %v \n", &g, &g, g, g)
	fmt.Printf("%p %v %p %v \n", &h, &h, h, h)
	fmt.Printf("%p %v %p %v \n", &i, &i, i, i)
	fmt.Printf("%p %v %p %v \n", &j, &j, j, j)
}

func main() {
	ZTestValZeroVal()
}
