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
	if err != nil {
		fmt.Println("1", err) // 运行到这里
		return
	}
	// ... logic
}

func main() {
	logic2()
}
