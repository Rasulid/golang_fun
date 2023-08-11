package main

import (
	"fmt"
	"os"
)

//https://www.golinuxcloud.com/golang-interface/

type selaryCalculator interface {
	CalculateSelary() float64
	report()
}

type PermanentEmployee struct {
	id          int
	basicSelary float64
	commission  float64
}

type ContractEmployee struct {
	id          int
	basicSalary float64
}

func (p PermanentEmployee) CalculateSelary() float64 {
	return p.basicSelary + (p.commission/100)*p.basicSelary
}

func (c ContractEmployee) CalculateSelary() float64 {
	return c.basicSalary
}

func (p PermanentEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month \n", p.id, p.CalculateSelary())
}

func (c ContractEmployee) report() {
	fmt.Printf("Employee ID %d earns USD %f per month \n", c.id, c.CalculateSelary())
}

func TotalSelary(s []selaryCalculator) float64 {
	var total float64

	for _, v := range s {
		total += v.CalculateSelary()
	}
	return total
}

func main() {
	//p1 := PermanentEmployee{id: 1, basicSelary: 2300, commission: 13}
	//p2 := PermanentEmployee{id: 2, basicSelary: 1500, commission: 18}
	//p3 := PermanentEmployee{id: 3, basicSelary: 2300, commission: 10}
	//c1 := ContractEmployee{id: 4, basicSalary: 500}
	//c2 := ContractEmployee{id: 5, basicSalary: 1100}
	//c3 := ContractEmployee{id: 6, basicSalary: 700}
	//
	//employees := []selaryCalculator{p1, p2, p3, c1, c2, c3}
	//
	//var totalSelary float64
	//
	//totalSelary = TotalSelary(employees)
	//fmt.Printf("Company total salary is : USD %f", totalSelary)

	_, err := os.ReadFile("employees.txt")
	if err != nil {
		e := ReadFileError{message: err.Error()}
		errorHendler(e)
		return
	}
	fmt.Println("Reading data from file")

}

// read file

type ReadFileError struct {
	message string
}

func (rfe ReadFileError) Error() string {
	return fmt.Sprintf("Custom Read File Error! Error Message : %s", rfe.message)

}

func errorHendler(err error) {
	fmt.Println("Error", err.Error())
}
