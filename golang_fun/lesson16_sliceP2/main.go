package main

import (
	"fmt"
)

func main() {
	variadicFunctions()
	convertToArray()
}

func variadicFunctions() {
	showAllElems(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	print("\n")
	showAllElems(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)

	firstSlice := []int{1, 2, 3, 4, 5, 6, 7}
	secondSlice := []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	showAllElems(firstSlice...)
	print("\n")
	showAllElems(secondSlice...)

	thirdSlice := append(firstSlice, secondSlice...)
	fmt.Println(thirdSlice)

}

func showAllElems(val ...int) {
	for _, v := range val {
		fmt.Println("Value :", v)
	}
}

func convertToArray() {
	slice := []int{1, 2}
	fmt.Printf("Type : %T, value: %v\n", slice, slice)
	fmt.Printf("Len : %d, Cap: %d\n", len(slice), cap(slice))

	sliceInArray := (*[2]int)(slice)
	fmt.Printf("Type : %T, value: %v\n", sliceInArray, sliceInArray)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(sliceInArray), cap(sliceInArray))

	// change slice

	changeslice := []int{1, 2}
	fmt.Printf("Type : %T, value: %v\n", changeslice, changeslice)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(changeslice), cap(changeslice))

	changeToValue(changeslice)

	fmt.Printf("Type : %T, value: %v\n", changeslice, changeslice)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(changeslice), cap(changeslice))

	// append

	newSlice := append(changeslice, 5)
	fmt.Printf("Type : %T, value: %v\n", newSlice, newSlice)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(newSlice), cap(newSlice))

	newSlice = appendValue(newSlice)
	fmt.Printf("Type : %T, value: %v\n", newSlice, newSlice)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(newSlice), cap(newSlice))

}

func changeToValue(val []int) {
	val[0] = 100
}

func appendValue(val []int) []int {
	val = append(val, 5, 6)
	//fmt.Printf("Type : %T, value: %v\n", val, val)
	//fmt.Printf("Len : %d, Cap: %d\n\n", len(val), cap(val))

	return val
}
