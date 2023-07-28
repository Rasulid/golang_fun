package main

import (
	"fmt"
	"unsafe"
)

//	variables , Constants , Types , Scope, Retype

const name = "variables"

func main() {
	//var hello bool
	hello := false
	fmt.Println(false)
	fmt.Println(name)
	fmt.Printf("Type : %T  Value : %v\n", hello, hello)

	s := Sprint("Rasul")
	println(s)

	fmt.Println(Retype())

	var num1 int8 = 12
	var num2 int64 = 1

	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num1))
}

func Sprint(name string) string {
	return fmt.Sprintf("hello %s", name)
}

func Retype() int64 {
	var i int64 = 23
	var n int = 21

	return int64(n) + i
}
