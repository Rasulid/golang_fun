package main

import "fmt"

func main() {
	var x *int
	fmt.Printf("type :%T , value : %#v\n", x, x)

	var a int64 = 7
	fmt.Printf("type :%T , value : %#v\n", a, a)

	var pointerToA *int64 = &a
	fmt.Printf("type :%T , addres : %#v , value : %#v\n", pointerToA, pointerToA, *pointerToA)

	var newPointer = new(float32)
	fmt.Printf("type: %T, address: %v, value: %v\n", newPointer, newPointer, *newPointer)

	*newPointer = 3
	fmt.Printf("type: %T, address: %v, value: %v\n", newPointer, newPointer, *newPointer)

}
