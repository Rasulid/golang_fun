package main

import "fmt"

type Animal interface {
	makeSound()
}

type greeting interface {
	SayHello(name string) string
}

type russian struct{}
type american struct{}

type Cat struct{}
type Dog struct{}

func (c *Cat) makeSound() {
	fmt.Println("Catttt ")
}

func (d *Dog) makeSound() {
	fmt.Println("Dogggg ")
}

func (r *russian) SayHello(name string) string {
	return fmt.Sprintf("Привет %s", name)
}

func (a *american) SayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func Greet(g greeting, name string) {
	fmt.Println(g.SayHello(name))
}

func main() {
	var cat, dog Animal = &Cat{}, &Dog{}

	dog.makeSound()
	cat.makeSound()

	Greet(&russian{}, "Жавох")
	Greet(&american{}, "Piter")

}
