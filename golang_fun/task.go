package main

import "fmt"

/*
Write two functions that calculates Fibonacci numbers.
The first function should use recursive strategy (recursion) to accomplish its job.
The second function should use iterative strategy (iteration) to accomplish its job.

*/

func main() {

	fmt.Println(FibonacciWithRecursion(10))
	fmt.Println(FibonacciWithIteration(10))

}

func FibonacciWithRecursion(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciWithRecursion(n-1) + FibonacciWithRecursion(n-2)
	}
}

func FibonacciWithIteration(n int) int {

	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
