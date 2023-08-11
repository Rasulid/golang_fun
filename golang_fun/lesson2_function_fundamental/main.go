package main

import "fmt"

func main() {

}
func Print() {
	fmt.Println("hello test")
}

func GetSomthing(a string) {
	fmt.Printf("hello %s\n", a)
}

func ReturnName(name string) string {
	fmt.Printf("hello %s\n", name)
	return name
}

func returnManyName(name, l_name string) (string, string) {
	return name, l_name
}

func returnManyNameV2(name, l_name string) (n string, l string) {

	n = name
	l = l_name
	return
}
