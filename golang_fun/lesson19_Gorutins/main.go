package main

import (
	"fmt"
)

func main() {
	// deffered function

	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//fmt.Println(Summary(1, 2))

	//deferValue()

	// PANIC

	MakaPanic()

	//runtime.GOMAXPROCS(1) //можно настроить сколько go рутин будет работать паралельно , в нашем случае (1)

	//fmt.Println(runtime.NumCPU())

	//go ShowNumber(100)

	//runtime.Gosched() //планировщик когда go рутин выполняется Это функция даёт времени чтобы выполнялся другие go рутины

	//time.Sleep(time.Second) // sleep for second

	print("Exiting...\n")

}

func ShowNumber(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func Summary(x, y int) (sum int) {

	defer func() { // defer Функция работае после return
		sum = sum * 2
	}()
	sum = x + y
	return

}

func deferValue() {
	for i := 0; i < 10; i++ {
		defer func(k int) {
			fmt.Println(k)
		}(i)
	}
}

func MakaPanic() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("Som Panic")

	//fmt.Println("Error")
}
