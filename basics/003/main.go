package main

import (
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

func prints(a string) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&a))
	spew.Dump(sh)
}

func main() {
	prints("abcd")
}
