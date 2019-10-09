package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(1000000))
}

func factorial(num int32) int32 {
	if num == 0 {
		return 1
	}
	return num + factorial(num-1)
}
