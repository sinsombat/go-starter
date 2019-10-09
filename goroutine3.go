package main

import (
	"fmt"
	// "runtime"
	"time"
)

func main() {
	// runtime.GOMAXPROCS(8)
	go abcGen()
	fmt.Println("This comes first!")

	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		fmt.Println(string(l))
	}
}
