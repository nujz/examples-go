package main

type EI interface{}
type I interface{ Echo() }
type ES struct{}
type S struct {
	a int
	b []int
}

func (s *S) Echo() {}

func IfaceWithNil() {
	// var b *I
	var s *S

	var aa I = s
	println(s, aa == nil)
}

func EfaceWithNil() {
	var a *int
	var b chan int
	var c func()
	var d interface{} //
	var e map[int]int
	var f []int
	var g EI //
	var h I  //
	var i *I

	var aa interface{} = a
	var ab = (interface{})(a)
	println("a", aa == nil, ab == nil, aa == ab)

	var ba interface{} = b
	var bb = (interface{})(b)
	println("b", ba == nil, bb == nil, ba == bb)

	var ca interface{} = c
	var cb = (interface{})(c)
	println("c", ca == nil, cb == nil)

	var da interface{} = d
	var db = (interface{})(d)
	println("d", da == nil, db == nil, da == db)

	var ea interface{} = e
	var eb = (interface{})(e)
	println("e", ea == nil, eb == nil)

	var fa interface{} = f
	var fb = (interface{})(f)
	println("f", fa == nil, fb == nil)

	var ga interface{} = g
	var gb = (interface{})(g)
	println("g", ga == nil, gb == nil, ga == gb)

	var ha interface{} = h
	var hb = (interface{})(h)
	println("h", ha == nil, hb == nil, ha == hb)

	var ia interface{} = i
	var ib = (interface{})(i)
	println("i", ia == nil, ib == nil, ia == ib)
}

func ZeroValueNil() {
	var a *int
	var b chan int
	var c func()
	var d interface{}
	var e map[int]int
	var f []int
	var g EI
	var h I
	var i *I

	println("a", a, a == nil)
	println("b", b, b == nil)
	println("c", c, c == nil)
	println("d", d, d == nil)
	println("e", e, e == nil)
	println("f", f, f == nil)
	println("g", g, g == nil)
	println("h", h, h == nil)
	println("i", i, i == nil)
}

func ZeroValueNotNil() {
	var a int     // 0
	var b float32 // 0.0
	var c bool    // false
	var d byte    // 0
	var e rune    // 0
	var f [2]int  // [0 0]
	var g uintptr // 0
	// var h ES
	// var i S

	// var l = []int{}       // =make([]int{},0)
	// var m = map[int]int{} // =make(map[int]int)

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f[0], f[1])
	println(g)

	// println(h)
	// println(i)
}

func main() {
	IfaceWithNil()
	// EfaceWithNil()
	// ZeroValueNil()
	// ZeroValueNotNil()
}
