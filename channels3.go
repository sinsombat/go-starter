package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)
	doneCh := make(chan bool)
	go abcGen(ch)
	go printer(ch, doneCh)
	fmt.Println("This comes first!")

	<-doneCh
}

func abcGen(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}

	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		println(l)
	}

	doneCh <- true
}
