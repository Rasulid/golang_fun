package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type WoodBuilder struct {
	Person
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

func main() {
	//wb := WoodBuilder{Person{Name: "Rasul", Age: 12}}
	//fmt.Println(wb)

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
