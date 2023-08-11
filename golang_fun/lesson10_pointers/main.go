package main

import "fmt"

func main() {
	num := 3
	fmt.Printf("%#v\n", num)
	square(&num)
	fmt.Printf("%#v\n", num)

	var wallet *int
	fmt.Println(hasWallet(wallet))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 100
	fmt.Println(hasWallet(&wallet3))

}

func square(num *int) {
	*num = *num * *num
}

func hasWallet(money *int) bool {
	return money != nil
}
