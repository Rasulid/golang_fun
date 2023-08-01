package main

import "fmt"

type constant struct {
	email       string
	phonenumber string
}

type person struct {
	name         string
	age          int
	constantInfo constant
}

func main() {
	tom := person{name: "Tom", age: 12,
		constantInfo: constant{
			email:       "rasulabduvaitov@gmail.com",
			phonenumber: "+998914774712",
		}}

	var TomPointer *person = &tom

	TomPointer.age = 29
	fmt.Println(TomPointer.age)
	(*TomPointer).age = 32
	fmt.Println(tom.age)

	(*TomPointer).name = "Dior"

	fmt.Println(tom.constantInfo.phonenumber)

}
