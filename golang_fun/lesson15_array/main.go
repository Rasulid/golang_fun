package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//Arreys()

	Slice()
}

func Arreys() {

	var arr [3]int
	fmt.Printf("Type: %T, value: %#v\n", arr, arr)

	arr[0] = 1
	fmt.Printf("Type: %T, value: %#v\n", arr, arr)

	person := [2]Person{
		{
			Name: "<NAME1>",
			Age:  30,
		},
		{
			Name: "<NAME2>",
			Age:  20,
		},
	}
	fmt.Printf("Type: %T, value: %#v\n", person, person)

	lenArr := [...]string{"name", "age", "surname", "email"}
	fmt.Printf("Type: %T, value: %#v\n", lenArr, lenArr)

	// len and for loop

	fmt.Printf("len %d, cap %d\n", len(lenArr), cap(lenArr))

	for i := 0; i < len(lenArr); i++ {
		fmt.Printf("index :%d , value: %s\n", i, lenArr[i])
	}

	for i, v := range lenArr {
		fmt.Printf("index :%d, value: %s\n", i, v)
	}

	lenArr2 := [...]int{1, 2, 3}
	newArr := ChangeArr(lenArr2)
	fmt.Println(lenArr2)
	fmt.Println(newArr)
}
func ChangeArr(arr [3]int) [3]int {
	arr[2] = 100
	return arr
}

// Slice

func Slice() {
	var sli []int
	fmt.Printf("Type: %T,  value: %v\n", sli, sli)

	sli = append(sli, 1)
	fmt.Printf("Type: %T,  value: %v\n", sli, sli)

	sli2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Type: %T,  value: %v\n", sli2, sli2)
	fmt.Printf("Len: %d,  cap: %d\n", len(sli2), cap(sli2))

	sli3 := make([]int, 4, 6)
	fmt.Printf("Type: %T,  value: %v\n", sli3, sli3)
	fmt.Printf("Len: %d,  cap: %d\n", len(sli3), cap(sli3))

}
