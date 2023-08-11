package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//WaitGroup()
	//WriteWithoutMutex()
	//WriteWithOneGoRuntimes()
	//WriteWithMutex()

	ReadWriteMutexWithout()
	ReadWriteMutexWithRWMutex()
}

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {

		go func(i int) {
			defer wg.Done()
			println(i)
		}(i)

	}

	wg.Wait()
	fmt.Println("Done")

}

func WriteWithoutMutex() {
	start := time.Now()
	var wg sync.WaitGroup
	var counter int

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()

	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func WriteWithOneGoRuntimes() {
	start := time.Now()
	var count int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		count++
	}
	fmt.Println(count)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func WriteWithMutex() {
	start := time.Now()
	var wg sync.WaitGroup
	var counter int
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()
			counter++
			mu.Unlock()
		}()

	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func ReadWriteMutexWithout() {
	var (
		start = time.Now()
		count int
		mu    sync.Mutex
		wg    sync.WaitGroup
	)

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()

			mu.Lock()
			_ = count
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			count++
			mu.Unlock()
		}()

	}
	wg.Wait()

	fmt.Println(count)
	fmt.Println(time.Now().Sub(start).Seconds())

}

func ReadWriteMutexWithRWMutex() {
	var (
		start = time.Now()
		count int
		mu    sync.RWMutex
		wg    sync.WaitGroup
	)
	for i := 0; i < 50; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()

			mu.RLock()
			_ = count
			mu.RUnlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(count)
	fmt.Println(time.Now().Sub(start).Seconds())
}
