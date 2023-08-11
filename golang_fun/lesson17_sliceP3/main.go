package main

import "fmt"

func main() {
	//getSlice()
	//copySlice()

	SLICE := []int{1, 2, 3, 4, 5}

	fmt.Printf("Value : %v\n", deleteElements(SLICE, 2))
}

func getSlice() {
	intArr := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Type : %T, value: %v\n", intArr, intArr)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(intArr), cap(intArr))

	intSlice := intArr[1:3]
	fmt.Printf("Type : %T, value: %v\n", intSlice, intSlice)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(intSlice), cap(intSlice))

	fullSlice := intArr[:]
	fmt.Printf("Type : %T, value: %v\n", fullSlice, fullSlice)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(fullSlice), cap(fullSlice))

	intArr[2] = 100

	fmt.Printf("Type : %T, value: %v\n", fullSlice, fullSlice)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(fullSlice), cap(fullSlice))

}

func copySlice() {
	destination := make([]string, 2, 3)
	source := []string{"a", "b", "c"}

	fmt.Println("Copied :", copy(destination, source))
	fmt.Printf("Type : %T, value: %v\n", destination, destination)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(destination), cap(destination))

	destination = make([]string, len(source))
	fmt.Println("Copied :", copy(destination, source))
	fmt.Printf("Type : %T, value: %v\n", destination, destination)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(destination), cap(destination))

	// Right copy

	rightCopy := append(make([]string, 0, len(source)), source...)
	fmt.Printf("Type : %T, value: %v\n", rightCopy, rightCopy)
	fmt.Printf("Len : %d, Cap: %d\n\n", len(rightCopy), cap(rightCopy))
}

func deleteElements(s []int, index int) []int {

	return append(s[:index], s[index+1:]...)

}
