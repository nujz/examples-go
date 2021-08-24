package main

type T interface {
	echo()
}

type One struct {
}

func (o *One) echo() {}

/*

接口能接受的类型取决于实现方法的接受者类型

*/

var one = &One{}
var t1 T = &one // err

var t2 T = One{} // err

func main() {

}
