package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var oper string

	fmt.Println("Enter first number := ")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number := ")
	fmt.Scanln(&num2)

	fmt.Println("Enter operation := ")
	fmt.Scanln(&oper)

	if oper == "+" {
		fmt.Println(plus(num1, num2))
	} else if oper == "-" {
		fmt.Println(minus(num1, num2))
	} else if oper == "/" {
		fmt.Println(divide(num1, num2))
	} else if oper == "*" {
		fmt.Println(multiply(num1, num2))
	} else {
		fmt.Println("Invalid oper")
	}
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func divide(a, b int) int {
	return a / b
}

func multiply(a, b int) int {
	return a * b
}
