package main

import (
	"fmt"
	"time"
)

//https://www.golinuxcloud.com/go-channels/

//func main() {
//s := []string{"localhost", "localhost1", "localhost3", "localhost4"}

//LinkDataToChannel(s)
//linkChannel := make(chan string)
//go ping(linkChannel)
//
//for {
//	_, ok := <-linkChannel
//
//	if ok {
//		fmt.Println("Connect")
//	} else {
//		fmt.Println("Disconnect")
//		break
//	}
//}

//	ch := make(chan string)
//	go sendData(ch)
//	go gettingData(ch)
//
//	fmt.Scanln()
//}
//
//func ping(channel chan string) {
//	links := []string{"localhost", "localhost1", "localhost3", "localhost4"}
//	for _, link := range links {
//		channel <- link
//	}
//	close(channel)
//}
//
//func LinkDataToChannel(links []string) {
//
//	linkChannels := make(chan string)
//
//	go func() {
//		for _, link := range links {
//			linkChannels <- link
//		}
//		close(linkChannels)
//	}()
//
//	for link := range linkChannels {
//		fmt.Println(link)
//	}
//
//}
//
//func sendData(ch chan string) {
//	fmt.Println("Sending a String into a channel")
//	ch <- "Hello world"
//	fmt.Println("String received from a channel...")
//	close(ch)
//}
//
//func gettingData(ch chan string) {
//	time.Sleep(2 * time.Second)
//	fmt.Println("Stating received data into a channel", <-ch)
//}
//
//func goroutine1(ch1 chan string) {
//	time.Sleep(time.Second)
//	for i := 0; i < 3; i++ {
//		ch1 <- fmt.Sprintf("%d ==> Channel 1 message", i)
//	}
//}
//
//func goroutine2(ch2 chan string) {
//	for i := 0; i < 3; i++ {
//		ch2 <- fmt.Sprintf("%d ==> Channel 2 message", i)
//	}
//
//}

//func main() {
//ch1 := make(chan string)
//ch2 := make(chan string)
//
//go goroutine1(ch1)
//go goroutine2(ch2)
//
//for {
//	time.Sleep(time.Second * 3)
//	select {
//
//	case value1 := <-ch1:
//		fmt.Println(value1)
//	case value2 := <-ch2:
//		fmt.Println(value2)
//	default:
//		fmt.Println("All channels are blocking")
//	}
//}
//
//	channel1 := make(chan int)
//	channel2 := make(chan int)
//
//	go func() {
//		channel1 <- 42
//	}()
//
//	go func() {
//		channel2 <- 100
//	}()
//
//	select {
//	case data := <-channel1:
//		fmt.Println("Данные из канала 1:", data)
//	case data := <-channel2:
//		fmt.Println("Данные из канала 2:", data)
//	}
//}

/*
	Send a Structure in a channel
*/

type HTTPResponse struct {
	statusCode int
	body       string
}

func ping(c chan HTTPResponse, resonse []HTTPResponse) {
	for i := 0; i < len(resonse); i++ {
		c <- resonse[i]
	}
	close(c)
}

func main() {
	//response := []HTTPResponse{
	//	{
	//		statusCode: 200,
	//		body:       "Resource OK",
	//	},
	//	{
	//		statusCode: 201,
	//		body:       "Resource Created",
	//	},
	//	{
	//		statusCode: 404,
	//		body:       "Resource Not Found",
	//	},
	//}

	//chanel := make(chan HTTPResponse, 5)
	//go ping(chanel, response)
	//
	//for v := range chanel {
	//	fmt.Println(v)
	//}

	ch := make(chan int)

	go Fib(10, ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func Fib(num int, ch chan int) {
	a, b := 1, 1

	for i := 0; i < num; i++ {
		ch <- a
		a, b = b, a+b
		time.Sleep(time.Second)
	}
	close(ch)
}
