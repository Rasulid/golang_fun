package main

import "fmt"

func main() {
	num1, num2 := 192, 200

	sumFunc := func(x, y int) int { return x + y }

	fmt.Println(oper(num1, num2, sumFunc))

	divide := CreateDivide(10)
	fmt.Println("Divide Func", divide((100)))

	// closing function

	dollar := 30

	getDollarValue := func() int { return dollar }

	fmt.Println(getDollarValue())
	dollar = 70

	fmt.Println(getDollarValue())

}

// take function as argument

func oper(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

// return function as object

func CreateDivide(div int) func(x int) int {
	divideFunc := func(x int) int {
		return x / div
	}
	return divideFunc
}
