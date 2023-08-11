package main

import (
	"fmt"
	"net/http"
)

type Rectangle struct {
	with   float64
	height float64
}

func (r *Rectangle) RectangleArea() float64 {
	return r.with * r.height
}

type Counter struct {
	value int
}

func (c *Counter) Increment() int {
	return c.value + 1
}

func (c *Counter) Decrement() int {
	return c.value - 1
}

func main() {
	value := Rectangle{with: 5, height: 6}
	fmt.Println(value.RectangleArea())

	val := Counter{
		value: 12,
	}
	fmt.Println(val.Increment())

	fmt.Println(val.value)

	fmt.Println(val.Decrement())

	fmt.Println(val.value)
	fmt.Println(http.Response{})

}
