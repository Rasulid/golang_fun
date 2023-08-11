package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type WoodBuilder struct {
	Person
	Friends []Person
	Name    string
}

type email struct {
	email    string
	password string
}

type StudentInfo struct {
	name   string
	l_name string
	email  email
}

func (p Person) PrintName() {
	fmt.Println("Person: ", p.Name)
}

func (w WoodBuilder) PrintName() {
	fmt.Println("WoodBuilder:", w.Name)
}

func main() {
	wb := WoodBuilder{Person: Person{Name: "Rasul", Age: 12},
		Friends: []Person{
			Person{Name: "John", Age: 12},
			Person{Name: "Mary", Age: 12}},
		Name: "Jovo"}
	wb.PrintName()
	//fmt.Println(wb.Name)
	fmt.Println("Print  :", wb)

	info := &StudentInfo{
		name:   "Rasul",
		l_name: "Abdi",
		email: email{
			email:    "<EMAIL>",
			password: "<PASSWORD>",
		},
	}

	fmt.Println(info)
	fmt.Printf("Type: %T , value: %#v\n", info, info)
}
