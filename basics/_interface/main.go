package main

import "fmt"

type EI interface{} // eface

type I interface { // iface
	echo()
}

func main() {
	var ei EI
	ei = map[int]int{}
	ei.(map[int]int)[0] = 1
	fmt.Println(ei)

	ei = []int{1, 2}
	fmt.Println(ei)
}
