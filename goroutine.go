package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	var message1 string
	var message2 string

	go func() {
		message1 = forloop1(3, &wg)
	}()

	go func() {
		message2 = forloop2(5, &wg)
	}()

	wg.Wait()

	fmt.Println("hello")
	fmt.Println(message1)
	fmt.Println(message2)
}

func forloop1(n int, wg *sync.WaitGroup) string {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		fmt.Println("go routine1 ", i)
	}
	return "loop1"
}

func forloop2(n int, wg *sync.WaitGroup) string {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		fmt.Println("go routine2 ", i)
	}
	return "loop2"
}
