package main

import (
	"fmt"
)

type animalAndPerson interface {
	print()
}

type person struct {
	name string
	age  int
}

type animal struct {
	name string
	age  int
}

func personPrint(person animalAndPerson) {
	person.print()
}

func (p *person) print() {
	fmt.Println("Name: ", p.name)
	fmt.Println("age: ", p.age)

}

func (a *animal) print() {
	fmt.Println("Name: ", a.name)
	fmt.Println("age: ", a.age)

}

func main() {
	var tom = person{"tom", 12}
	var a = animal{name: "LIon", age: 12}
	
	var obraz animalAndPerson = &tom
	obraz.print()

	var intanimal animalAndPerson = &a
	intanimal.print()

	structure := person{name: "John", age: 12}

	personPrint(&structure)
}
