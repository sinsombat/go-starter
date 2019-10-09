package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	go forloop1(&messages, 100000)
	forloop2(10000)
	fmt.Println("hello", messages)
}

func forloop1(messages *chan string, n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("go routine ", i)
	}
	messages <- "ping"
}

func forloop2(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("go routine2 ", i)
	}
}
