package main

import (
	"fmt"
	"os"
	"strconv"
)

type CustomError struct {
	code int
	msg  string
}

func (e *CustomError) Error() string {
	return strconv.Itoa(e.code) + " -> " + e.msg
}

func logic1() error {
	var err *CustomError
	if os.Args[0] == "abcdefg1234567" {
		return nil
	} else if os.Args[0] == "abcdefg" {
		err = &CustomError{
			code: 400,
			msg:  "Bad Parameter",
		}
	}
	// ... logic
	return err
}

func logic2() {
	err := logic1()

	// if fmt.Sprint(err) != "<nil>" { // ???
	if err != nil { // interface && typed nil != untyped nil
		fmt.Println("1", err) // 运行到这里
		return
	}
	// ... logic
}

func compareError(err1, err2 error) {
	fmt.Println(err1, err1 == nil) // nil false | typed nil
	fmt.Println(err2, err2 == nil) // nil true | untyped nil
	fmt.Println(err1 == err2)      // false | typed nil != untyped nil
}

func CompareError() {
	var e *CustomError
	compareError(e, nil)
}

func main() {
	logic2()
}
