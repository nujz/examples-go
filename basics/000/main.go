package main

type T interface {
	echo()
}

type One struct {
}

func (o *One) echo() {}

var one = &One{}
var t1 T = &one // err

var t2 T = One{} // err

func main() {

}
