package main

import "fmt"

type Runner interface {
	Run() string
}

type Duck struct {
	Name, Surname string
}

type Human struct {
	Name string
}

func (d Duck) Run() string {
	return "Duck is run"
}

func (h Human) Run() string {
	return fmt.Sprintf("human is running %s", h.Name)
}

func main() {

	interfaceValues()

}

func interfaceValues() {
	var run Runner
	fmt.Printf("Type : %T, value : %#v\n", run, run)

	if run == nil {
		fmt.Println("Runner is nill")
	}

	var unNumedRunner *Human
	fmt.Printf("Type : %T, value : %#v\n", unNumedRunner, unNumedRunner)

	run = unNumedRunner
	fmt.Printf("Type : %T, value : %#v\n", run, run)

	if run == nil {
		fmt.Println("Runner is nill")
	}

	namedRunner := &Human{
		Name: "John",
	}

	fmt.Printf("Type : %T, value : %#v\n", namedRunner, namedRunner)

	run = *namedRunner
	fmt.Printf("Type : %T, value : %#v\n", run, run)

	fmt.Println(run.Run())

	// empty interface

	var mpInterface interface{}
	fmt.Printf("Type : %T, value : %#v\n", mpInterface, mpInterface)

	mpInterface = int64(5)
	fmt.Printf("Type : %T, value : %#v\n", mpInterface, mpInterface)

	mpInterface = true
	fmt.Printf("Type : %T, value : %#v\n", mpInterface, mpInterface)

}
