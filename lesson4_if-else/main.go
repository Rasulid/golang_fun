package main

import "fmt"

func main() {
	age := 18

	if age >= 7 && /* || */ age < 19 {
		fmt.Println("You are in school")
	}
}
