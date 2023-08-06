package main

import "fmt"

type Procedure func(int) int

func TranformSlice(slic []int, procedure Procedure) []int {
	transFormSlice := make([]int, len(slic))
	for i, v := range slic {
		transFormSlice[i] = procedure(v)
	}

	return transFormSlice
}

func MultyplyTo2(n int) int {
	return n * 2
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	transFormSlice := TranformSlice(s, MultyplyTo2)

	fmt.Println(transFormSlice)
}
